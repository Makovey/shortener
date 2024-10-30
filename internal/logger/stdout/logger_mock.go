package stdout

import (
	"github.com/Makovey/shortener/internal/logger"
)

type loggerStdoutMock struct {
}

func (l loggerStdoutMock) Info(msg string, args ...string) {
}

func (l loggerStdoutMock) Error(msg string, args ...string) {
}

func (l loggerStdoutMock) Debug(msg string, args ...string) {
}

func (l loggerStdoutMock) Warning(msg string, args ...string) {
}

func NewLoggerStdoutMock() logger.Logger {
	return &loggerStdoutMock{}
}
