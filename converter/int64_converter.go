package converter

import (
	"reflect"
	"strconv"
)

type Int64Converter struct {
}

func (int64Converter *Int64Converter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, 64); err != nil {
		return nil, err
	} else {
		return int64(value), nil
	}
}

func (int64Converter *Int64Converter) Support() reflect.Type {
	return reflect.TypeOf(*new(int64))
}
