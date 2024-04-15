package lifecycle

import (
	"reflect"
)

type AfterContainerInject interface {
	AfterContainerInjectAction(objectContainer map[reflect.Type]interface{}) error
	GetPriority() int
}
