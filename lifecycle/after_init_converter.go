package lifecycle

import (
	"github.com/lyzzz123/illusion/converter"
	"reflect"
)

type AfterInitConverter interface {
	AfterInitConverterAction(typeConverterMap map[reflect.Type]converter.Converter) error
}
