package logger

import (
	std_log "log"
	"strings"
)

var _ = std_log.Print

type Level int

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var levelColor = map[Level]string{
	TRACE: "36",
	DEBUG: "32",
	INFO:  "33",
	WARN:  "31",
	ERROR: "31",
	FATAL: "35",
}

func NewLevel(s string) Level {
	s = strings.ToLower(s)
	switch s {
	case "trace":
		return TRACE
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		std_log.Printf("invalid log level '%s'", s)
		return TRACE
	}
}

func NewLevelText(level Level) string {
	switch level {
	case TRACE:
		return "trace"
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	case FATAL:
		return "fatal"
	default:
		std_log.Printf("invalid log level '%d'", level)
		return "trace"
	}
}

func (l Level) Color() string {
	return levelColor[l]
}
