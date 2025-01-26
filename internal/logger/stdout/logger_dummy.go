package stdout

import (
	"github.com/Makovey/shortener/internal/logger"
)

type loggerDummy struct {
}

// NewLoggerDummy конструктор dummy логгера
func NewLoggerDummy() logger.Logger {
	return &loggerDummy{}
}

// Info - пустой метод для мокирования
func (l loggerDummy) Info(msg string, args ...string) {
}

// Error - пустой метод для мокирования
func (l loggerDummy) Error(msg string, args ...string) {
}

// Debug - пустой метод для мокирования
func (l loggerDummy) Debug(msg string, args ...string) {
}

// Warning - пустой метод для мокирования
func (l loggerDummy) Warning(msg string, args ...string) {
}
