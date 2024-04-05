package converter

import (
	"reflect"
	"strconv"
)

type Uint8PtrConverter struct {
}

func (uint8PtrConverter *Uint8PtrConverter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseUint(param, 10, 8); err != nil {
		return nil, err
	} else {
		ui8 := uint8(value)
		return &ui8, nil
	}
}

func (uint8PtrConverter *Uint8PtrConverter) Support() reflect.Type {
	return reflect.TypeOf(new(uint8))
}
