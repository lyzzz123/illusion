package converter

import (
	"reflect"
	"strconv"
)

type UintPtrConverter struct {
}

func (uintPtrConverter *UintPtrConverter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseUint(param, 10, strconv.IntSize); err != nil {
		return nil, err
	} else {
		v := uint(value)
		return &v, nil
	}
}

func (uintPtrConverter *UintPtrConverter) Support() reflect.Type {
	return reflect.TypeOf(new(uint))
}
