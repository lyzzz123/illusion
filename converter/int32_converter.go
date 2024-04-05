package converter

import (
	"reflect"
	"strconv"
)

type Int32Converter struct {
}

func (int32Converter *Int32Converter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, 32); err != nil {
		return nil, err
	} else {
		return int32(value), nil
	}
}

func (int32Converter *Int32Converter) Support() reflect.Type {
	return reflect.TypeOf(*new(int32))
}
