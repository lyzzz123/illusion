package test

import (
	"github.com/lyzzz123/illusion"
	"github.com/lyzzz123/illusion/log"
	"reflect"
	"testing"
)

type LogConfigure struct {
	LogInstance log.Log `require:"true"`
}

func (logConfigure *LogConfigure) AfterInitInjectAction(objectContainer map[reflect.Type]interface{}) error {

	log.RegisterLog(logConfigure.LogInstance)
	return nil
}

func TestLog(t *testing.T) {

	illusion.Register(&log.DefaultLog{})
	illusion.Register(&LogConfigure{})
	illusion.Start()
	log.Info("asdfasdf")
}
