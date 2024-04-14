package test

import (
	"github.com/lyzzz123/illusion"
	"github.com/lyzzz123/illusion/log"
	"reflect"
	"testing"
)

type LogConfigure struct {
	LogInstance log.Log `require:"true"`
	TestLogName string  `property:"log.name, false"`
}

func (logConfigure *LogConfigure) AfterContainerInjectAction(objectContainer map[reflect.Type]interface{}) error {
	log.RegisterLog(logConfigure.LogInstance)
	return nil
}

type LogProxy struct {
	Target interface{}
}

func (logProxy *LogProxy) SupportInterface() reflect.Type {
	return reflect.TypeOf(new(log.Log)).Elem()
}

func (logProxy *LogProxy) SetTarget(target interface{}) {
	logProxy.Target = target
}

func (logProxy *LogProxy) Debug(format string, args ...interface{}) {
	logger := logProxy.Target.(log.Log)
	logger.Debug("proxy")
	logger.Debug(format, args...)
}

func (logProxy *LogProxy) Info(format string, args ...interface{}) {
	logger := logProxy.Target.(log.Log)
	logger.Info("proxy")
	logger.Info(format, args...)
}

func (logProxy *LogProxy) Warn(format string, args ...interface{}) {
	logger := logProxy.Target.(log.Log)
	logger.Warn("proxy")
	logger.Warn(format, args...)
}

func (logProxy *LogProxy) Error(format string, args ...interface{}) {
	logger := logProxy.Target.(log.Log)
	logger.Error("proxy")
	logger.Error(format, args...)
}

func TestLog(t *testing.T) {

	illusion.Register(&log.DefaultLog{})
	illusion.Register(&LogProxy{})
	logConfigure := &LogConfigure{}
	illusion.Register(logConfigure)
	illusion.TestStart()
	log.Info("asdfasdf:{}, ff:{}", "ffff", 2)
}
