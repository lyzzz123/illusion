package log

var LogInstance Log

func RegisterLog(l Log) {
	if err := l.Config(); err != nil {
		panic(err)
	}
	LogInstance = l
}

func Debug(format string, args ...interface{}) {
	LogInstance.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	LogInstance.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	LogInstance.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	LogInstance.Error(format, args...)
}

type Log interface {
	Config() error

	Debug(format string, args ...interface{})

	Info(format string, args ...interface{})

	Warn(format string, args ...interface{})

	Error(format string, args ...interface{})
}
