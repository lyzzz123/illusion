package converter

import (
	"reflect"
	"strconv"
)

type Int16Converter struct {
}

func (int16Converter *Int16Converter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, 16); err != nil {
		return nil, err
	} else {
		return int16(value), nil
	}
}

func (int16Converter *Int16Converter) Support() reflect.Type {
	return reflect.TypeOf(*new(int16))
}
