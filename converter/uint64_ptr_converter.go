package converter

import (
	"reflect"
	"strconv"
)

type Uint64PtrConverter struct {
}

func (uint64PtrConverter *Uint64PtrConverter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseUint(param, 10, 64); err != nil {
		return nil, err
	} else {
		ui64 := uint64(value)
		return &ui64, nil
	}
}

func (uint64PtrConverter *Uint64PtrConverter) Support() reflect.Type {
	return reflect.TypeOf(new(uint64))
}
