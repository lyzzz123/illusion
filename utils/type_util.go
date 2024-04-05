package utils

import (
	"reflect"
)

func Implement(object interface{}, interfaceType reflect.Type) bool {
	objectType := reflect.TypeOf(object)
	return objectType.Implements(interfaceType)
}
