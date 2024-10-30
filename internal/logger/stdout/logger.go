package stdout

import (
	"os"

	"github.com/sirupsen/logrus"

	def "github.com/Makovey/shortener/internal/logger"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

type loggerStdout struct {
	log *logrus.Logger
}

func (l loggerStdout) Info(msg string, args ...string) {
	fields := makeFieldsFromArgs(args...)
	l.log.WithFields(fields).Infoln(msg)
}

func (l loggerStdout) Error(msg string, args ...string) {
	fields := makeFieldsFromArgs(args...)
	l.log.WithFields(fields).Errorln(msg)
}

func (l loggerStdout) Debug(msg string, args ...string) {
	fields := makeFieldsFromArgs(args...)
	l.log.WithFields(fields).Debugln(msg)
}

func (l loggerStdout) Warning(msg string, args ...string) {
	fields := makeFieldsFromArgs(args...)
	l.log.WithFields(fields).Warningln(msg)
}

func NewLoggerStdout(env string) def.Logger {
	var log *logrus.Logger

	switch env {
	case envDev:
		log = &logrus.Logger{
			Out:       os.Stdout,
			Formatter: new(logrus.JSONFormatter),
			Level:     logrus.DebugLevel,
		}
	case envProd:
		log = &logrus.Logger{
			Out:       os.Stdout,
			Formatter: new(logrus.JSONFormatter),
			Level:     logrus.InfoLevel,
		}
	default:
		log = &logrus.Logger{
			Out: os.Stdout,
			Formatter: &logrus.TextFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
				FullTimestamp:   true,
			},
			Level: logrus.DebugLevel,
		}
	}

	return &loggerStdout{log: log}
}

func makeFieldsFromArgs(args ...string) logrus.Fields {
	var fields = logrus.Fields{}

	if len(args) > 0 {
		for i := 0; i < len(args); i += 2 {
			if i+1 < len(args) {
				fields[args[i]] = args[i+1]
			} else {
				fields[args[i]] = ""
			}
		}
	}

	return fields
}
