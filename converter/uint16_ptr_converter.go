package converter

import (
	"reflect"
	"strconv"
)

type Uint16PtrConverter struct {
}

func (uint16PtrConverter *Uint16PtrConverter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseUint(param, 10, 16); err != nil {
		return nil, err
	} else {
		ui16 := uint16(value)
		return &ui16, nil
	}
}

func (uint16PtrConverter *Uint16PtrConverter) Support() reflect.Type {
	return reflect.TypeOf(new(uint16))
}
