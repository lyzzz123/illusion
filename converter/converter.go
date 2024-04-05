package converter

import (
	"reflect"
)

type Converter interface {
	Convert(param string) (interface{}, error)
	Support() reflect.Type
}
