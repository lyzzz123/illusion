package lifecycle

import "reflect"

type AfterContainerDestroy interface {
	AfterContainerDestroyAction(objectContainer map[reflect.Type]interface{}) error
	GetPriority() int
}
