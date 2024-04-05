package converter

import (
	"reflect"
	"strconv"
)

type Int32PtrConverter struct {
}

func (int32PtrConverter *Int32PtrConverter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, 32); err != nil {
		return nil, err
	} else {
		i32 := int32(value)
		return &i32, nil
	}
}

func (int32PtrConverter *Int32PtrConverter) Support() reflect.Type {
	return reflect.TypeOf(new(int32))
}
