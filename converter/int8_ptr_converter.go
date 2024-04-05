package converter

import (
	"reflect"
	"strconv"
)

type Int8PtrConverter struct {
}

func (int8PtrConverter *Int8PtrConverter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, 8); err != nil {
		return nil, err
	} else {
		i8 := int8(value)
		return &i8, nil
	}
}

func (int8PtrConverter *Int8PtrConverter) Support() reflect.Type {
	return reflect.TypeOf(new(int8))
}
