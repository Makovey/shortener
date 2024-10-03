package logger

type Logger interface {
	Info(msg string, args ...string)
	Debug(msg string, args ...string)
	Warning(msg string, args ...string)
	Error(msg string, args ...string)
}
