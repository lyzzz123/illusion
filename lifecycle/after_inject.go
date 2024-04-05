package lifecycle

import (
	"reflect"
)

type AfterInitInject interface {
	AfterInitInjectAction(objectContainer map[reflect.Type]interface{}) error
}
