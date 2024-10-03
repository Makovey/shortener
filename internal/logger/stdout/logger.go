package stdout

import (
	"log/slog"
	"os"

	def "github.com/Makovey/shortener/internal/logger"
)

type loggerStdout struct {
	log *slog.Logger
}

func (l *loggerStdout) Info(msg string, args ...string) {
	l.log.Info(msg, slog.Any("args", args))
}

func (l *loggerStdout) Error(msg string, args ...string) {
	l.log.Error(msg, slog.Any("args", args))
}

func (l *loggerStdout) Debug(msg string, args ...string) {
	l.log.Debug(msg, slog.Any("args", args))
}

func (l *loggerStdout) Warning(msg string, args ...string) {
	l.log.Warn(msg, slog.Any("args", args))
}

func NewLoggerStdout() def.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &loggerStdout{log: logger}
}
