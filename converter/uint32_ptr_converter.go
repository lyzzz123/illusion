package converter

import (
	"reflect"
	"strconv"
)

type Uint32PtrConverter struct {
}

func (uint32PtrConverter *Uint32PtrConverter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseUint(param, 10, 32); err != nil {
		return nil, err
	} else {
		ui32 := uint32(value)
		return &ui32, nil
	}
}

func (uint32PtrConverter *Uint32PtrConverter) Support() reflect.Type {
	return reflect.TypeOf(new(uint32))
}
