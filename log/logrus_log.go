package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

type LogrusLog struct {
}

func (logrusLog *LogrusLog) Config() error {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	return nil
}

func (logrusLog *LogrusLog) Debug(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (logrusLog *LogrusLog) Info(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func (logrusLog *LogrusLog) Warn(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func (logrusLog *LogrusLog) Error(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}
