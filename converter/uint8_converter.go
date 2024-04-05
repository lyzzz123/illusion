package converter

import (
	"reflect"
	"strconv"
)

type Uint8Converter struct {
}

func (uint8Converter *Uint8Converter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseUint(param, 10, 8); err != nil {
		return nil, err
	} else {
		return uint8(value), nil
	}
}

func (uint8Converter *Uint8Converter) Support() reflect.Type {
	return reflect.TypeOf(*new(uint8))
}
