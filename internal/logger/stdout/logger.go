package stdout

import (
	"log/slog"
	"os"

	def "github.com/Makovey/shortener/internal/logger"
)

type loggerStdout struct {
	log *slog.Logger
}

func (l *loggerStdout) Info(msg string, args ...interface{}) {
	l.log.Info(msg, args)
}

func (l *loggerStdout) Error(msg string, args ...interface{}) {
	l.log.Error(msg, args)
}

func (l *loggerStdout) Debug(msg string, args ...interface{}) {
	l.log.Debug(msg, args)
}

func (l *loggerStdout) Warning(msg string, args ...interface{}) {
	l.log.Warn(msg, args)
}

func NewLoggerStdout() def.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &loggerStdout{log: logger}
}
