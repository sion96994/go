package logger

import (
	"fmt"
	"io"
	std_log "log"
	"os"
	"path"
	"runtime"
	"strings"
)

type Logger struct {
	Level     Level
	Colorful  bool
	ShowLine  bool
	Prefix    string
	StackSkip int

	l    *std_log.Logger
	init bool
}

func NewLogger(w io.Writer) *Logger {
	return &Logger{
		Level:     TRACE,
		Colorful:  true,
		ShowLine:  true,
		StackSkip: 4,
		l:         std_log.New(w, "", std_log.Ldate|std_log.Ltime|std_log.Lmicroseconds),
	}
}

func NewLoggerBySkip(w io.Writer, skip int) *Logger {
	return &Logger{
		Level:     TRACE,
		Colorful:  true,
		ShowLine:  true,
		StackSkip: skip,
		l:         std_log.New(w, "", std_log.Ldate|std_log.Ltime|std_log.Lmicroseconds),
	}
}

func NewNamedLogger(name string, w io.Writer) *Logger {
	section := "log"
	name = strings.TrimSpace(name)
	if len(name) > 0 {
		section += "_" + name
	}
	l := NewLogger(w)
	loggers[name] = l
	regLogConfig(section, l)
	return l
}

func GetLogger(name string) *Logger {
	l, ok := loggers[name]
	if !ok {
		Fatalf("can't find logger '%s'", name)
		return loggers[""]
	}
	return l
}

var loggers = make(map[string]*Logger)

func init() {
	NewNamedLogger("", os.Stdout)
}

func color(col, s string) string {
	if col == "" {
		return s
	}
	return "\x1b[0;" + col + "m" + s + "\x1b[0m"
}

func (l *Logger) SetPrefix(prefix string) *Logger {
	l.Prefix = prefix
	return l
}

func (l *Logger) Skip(skip int) *Logger {
	nl := *l
	nl.StackSkip = skip
	return &nl
}

func (l *Logger) getPrefix(level Level) string {
	s := level.String()
	if l.Colorful {
		s = color(level.Color(), s)
	}
	s = "[" + s + "]"
	if l.Prefix != "" {
		s = s + " " + l.Prefix
	}
	return s
}

func (l *Logger) getPosix() string {
	if !l.ShowLine {
		return ""
	}
	_, file, line, ok := runtime.Caller(l.StackSkip)

	if !ok {
		return ""
	}
	file = path.Base(file)
	return fmt.Sprintf("[%s:%d]", file, line)
}

func (l *Logger) Output(level Level, format string, a ...interface{}) {
	if level < l.Level {
		return
	}
	s := ""
	if len(a) != 0 {
		s = fmt.Sprintf(format, a...)
	} else {
		s = fmt.Sprint(format)
	}
	content := l.getPrefix(level) + " " + s + " " + l.getPosix()
	l.l.Println(content)
}

func (l *Logger) Printf(format string, a ...interface{}) {
	l.Output(TRACE, format, a...)
}

func (l *Logger) Tracef(format string, a ...interface{}) {
	l.Output(TRACE, format, a...)
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	l.Output(DEBUG, format, a...)
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.Output(INFO, format, a...)
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	l.Output(WARN, format, a...)
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.Output(ERROR, format, a...)
}

func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.Output(FATAL, format, a...)
	os.Exit(1)
}

func Printf(format string, a ...interface{}) {
	loggers[""].Printf(format, a...)
}

func Tracef(format string, a ...interface{}) {
	loggers[""].Tracef(format, a...)
}

func Debugf(format string, a ...interface{}) {
	loggers[""].Debugf(format, a...)
}

func Infof(format string, a ...interface{}) {
	loggers[""].Infof(format, a...)
}

func Warnf(format string, a ...interface{}) {
	loggers[""].Warnf(format, a...)
}

func Errorf(format string, a ...interface{}) {
	loggers[""].Errorf(format, a...)
}

func Fatalf(format string, a ...interface{}) {
	loggers[""].Fatalf(format, a...)
}

func DefaultLogger() *Logger {
	return loggers[""]
}

func ErrorfWithoutReport(err string) {
	std_log.Println(fmt.Sprintf("[%s] %s", color(levelColor[ERROR], "ERROR"), err))
}

func TracefWithoutReport(err string) {
	std_log.Println(fmt.Sprintf("[%s] %s", color(levelColor[TRACE], "TRACE"), err))
}

func InfofWithoutReport(err string) {
	std_log.Println(fmt.Sprintf("[%s] %s", color(levelColor[INFO], "INFO"), err))
}
