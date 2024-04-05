package converter

import (
	"reflect"
	"strconv"
)

type Int8Converter struct {
}

func (int8Converter *Int8Converter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, 8); err != nil {
		return nil, err
	} else {
		return int8(value), nil
	}
}

func (int8Converter *Int8Converter) Support() reflect.Type {
	return reflect.TypeOf(*new(int8))
}
