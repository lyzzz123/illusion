package converter

import (
	"reflect"
	"strconv"
)

type Float32PtrConvert struct {
}

func (float32PtrConvert *Float32PtrConvert) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseFloat(param, 32); err != nil {
		return nil, err
	} else {
		f32 := float32(value)
		return &f32, nil
	}
}

func (float32PtrConvert *Float32PtrConvert) Support() reflect.Type {
	return reflect.TypeOf(new(float32))
}
