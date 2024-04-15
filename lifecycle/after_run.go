package lifecycle

import "reflect"

type AfterRun interface {
	AfterRunAction(objectContainer map[reflect.Type]interface{}) error
	GetPriority() int
}
