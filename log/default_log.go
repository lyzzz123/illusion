package log

import (
	"bytes"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type Formatter struct {
}

func (t Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	//字节缓冲区
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:06")
	//自定义文件路径
	pc, file, line, _ := runtime.Caller(10)
	//自定义输出格式
	fmt.Fprintf(b, "[%s][%s][%s][%s:%d] - %s\n", timestamp, entry.Level, runtime.FuncForPC(pc).Name(), path.Base(file), line, entry.Message)
	return b.Bytes(), nil
}

type DefaultLog struct {
	Path       string `property:"log.path"`
	RotateName string `property:"log.rotate.name"`
	RotateTime int    `property:"log.rotate.time"`
	RotateSize int64  `property:"log.rotate.size"`
	Output     string `property:"log.output"`
	Level      string `property:"log.level"`
}

func (defaultLog *DefaultLog) Init() {
	logrus.SetLevel(logrus.DebugLevel)
	if defaultLog.Level == "warn" {
		logrus.SetLevel(logrus.WarnLevel)
	} else if defaultLog.Level == "info" {
		logrus.SetLevel(logrus.InfoLevel)
	} else if defaultLog.Level == "error" {
		logrus.SetLevel(logrus.ErrorLevel)
	}
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&Formatter{})
	if defaultLog.Output == "file" {
		writer, _ := rotatelogs.New(
			defaultLog.RotateName,
			rotatelogs.WithLinkName(defaultLog.Path),
			rotatelogs.WithMaxAge(time.Duration(defaultLog.RotateTime)*time.Hour),
			rotatelogs.WithRotationSize(1024*1024*defaultLog.RotateSize),
		)
		logrus.SetOutput(writer)
	}
}

func (defaultLog *DefaultLog) Debug(format string, args ...interface{}) {
	format = strings.ReplaceAll(format, "{}", "%v")
	logrus.Debugf(format, args...)
}

func (defaultLog *DefaultLog) Info(format string, args ...interface{}) {
	format = strings.ReplaceAll(format, "{}", "%v")
	logrus.Infof(format, args...)
}

func (defaultLog *DefaultLog) Warn(format string, args ...interface{}) {
	format = strings.ReplaceAll(format, "{}", "%v")
	logrus.Warnf(format, args...)
}

func (defaultLog *DefaultLog) Error(format string, args ...interface{}) {
	format = strings.ReplaceAll(format, "{}", "%v")
	logrus.Errorf(format, args...)
}
