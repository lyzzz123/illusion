package converter

import (
	"reflect"
	"strconv"
)

type Uint16Converter struct {
}

func (uint16Converter *Uint16Converter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseUint(param, 10, 16); err != nil {
		return nil, err
	} else {
		return uint16(value), nil
	}
}

func (uint16Converter *Uint16Converter) Support() reflect.Type {
	return reflect.TypeOf(*new(uint16))
}
