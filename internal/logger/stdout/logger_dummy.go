package stdout

import (
	"github.com/Makovey/shortener/internal/logger"
)

type loggerDummy struct {
}

func NewLoggerDummy() logger.Logger {
	return &loggerDummy{}
}

func (l loggerDummy) Info(msg string, args ...string) {
}

func (l loggerDummy) Error(msg string, args ...string) {
}

func (l loggerDummy) Debug(msg string, args ...string) {
}

func (l loggerDummy) Warning(msg string, args ...string) {
}
