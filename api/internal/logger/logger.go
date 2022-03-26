package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

// Logger log.Trace(conf.LogLevel)
//log.Debug(conf.LogLevel)
//log.Info(conf.LogLevel)
//log.Warn(conf.LogLevel)
//log.Error(conf.LogLevel)
// Calls os.Exit(1) after logging.Fatal("Bye.")
// Calls panic() after logging.Panic(conf.LogLevel)
type Logger struct {
	*logrus.Logger
}

func logLevelFormat(level string) logrus.Level {
	lowerLvl := strings.ToLower(level)
	switch lowerLvl {
	case "trace":
		return 6
	case "debug":
		return 5
	case "info":
		return 4
	case "warn":
		return 3
	case "error":
		return 2
	case "fatal":
		return 1
	case "panic":
		return 0
	default:
		return 1
	}
}

func Init(level string) *Logger {
	lvl := logLevelFormat(level)
	log := logrus.New()
	log.SetLevel(lvl)
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "01/02 15:04:05",
	})
	return &Logger{log}
}

// Struct Logger print struct property
func (l *Logger) Struct(s interface{}) {
	fmt.Printf("Struct: %#v\n", s)
}
