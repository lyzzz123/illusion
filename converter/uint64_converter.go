package converter

import (
	"reflect"
	"strconv"
)

type Uint64Converter struct {
}

func (uint64Converter *Uint64Converter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseUint(param, 10, 64); err != nil {
		return nil, err
	} else {
		return uint64(value), nil
	}
}

func (uint64Converter *Uint64Converter) Support() reflect.Type {
	return reflect.TypeOf(*new(uint64))
}
