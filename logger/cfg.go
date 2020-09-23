package logger

import (
	"strconv"

	"github.com/sion96994/go/utils/cfgo"
)

type LogConfig struct {
	Level     string
	Colorful  string `yaml:",omitempty"`
	ShowLine  string `yaml:",omitempty"`
	Prefix    string `yaml:",omitempty"`
	StackSkip string `yaml:",omitempty"`
	l         *Logger
	init      bool
}

// 不支持reload
func (l *LogConfig) Reload(bind cfgo.BindFunc) error {
	if !l.init {
		if err := bind(); err != nil {
			return err
		}
		l.init = true
		l.l.Level = NewLevel(l.Level)
		if len(l.Colorful) > 0 {
			l.l.Colorful, _ = strconv.ParseBool(l.Colorful)
		}
		if len(l.ShowLine) > 0 {
			l.l.ShowLine, _ = strconv.ParseBool(l.ShowLine)
		}
		if len(l.Prefix) > 0 {
			l.l.Prefix = l.Prefix
		}
		if len(l.StackSkip) > 0 {
			l.l.StackSkip, _ = strconv.Atoi(l.StackSkip)
		}
	}
	return nil
}

func regLogConfig(section string, l *Logger) {
	cfgo.MustReg(section, &LogConfig{
		Level: NewLevelText(l.Level),
		l:     l,
	})
}