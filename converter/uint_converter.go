package converter

import (
	"reflect"
	"strconv"
)

type UintConverter struct {
}

func (uintConverter *UintConverter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseUint(param, 10, strconv.IntSize); err != nil {
		return nil, err
	} else {
		return uint(value), nil
	}
}

func (uintConverter *UintConverter) Support() reflect.Type {
	return reflect.TypeOf(*new(uint))
}
