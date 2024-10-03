package logger

type Logger interface {
	Info(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Warning(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}
