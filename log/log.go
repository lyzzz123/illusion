package log

var instance Log

func RegisterLog(l Log) {
	instance = l
	instance.Init()
}

func Debug(format string, args ...interface{}) {
	instance.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	instance.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	instance.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	instance.Error(format, args...)
}

type Log interface {
	Init()

	Debug(format string, args ...interface{})

	Info(format string, args ...interface{})

	Warn(format string, args ...interface{})

	Error(format string, args ...interface{})
}
