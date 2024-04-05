package converter

import (
	"reflect"
	"strconv"
)

type Uint32Converter struct {
}

func (uint32Converter *Uint32Converter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseUint(param, 10, 32); err != nil {
		return nil, err
	} else {
		return uint32(value), nil
	}
}

func (uint32Converter *Uint32Converter) Support() reflect.Type {
	return reflect.TypeOf(*new(uint32))
}
