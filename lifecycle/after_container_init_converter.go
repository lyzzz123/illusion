package lifecycle

import (
	"github.com/lyzzz123/illusion/converter"
	"reflect"
)

type AfterContainerInitConverter interface {
	AfterContainerInitConverterAction(typeConverterMap map[reflect.Type]converter.Converter) error
}
