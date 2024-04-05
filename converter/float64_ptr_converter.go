package converter

import (
	"reflect"
	"strconv"
)

type Float64PtrConvert struct {
}

func (float64PtrConvert *Float64PtrConvert) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseFloat(param, 64); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func (float64PtrConvert *Float64PtrConvert) Support() reflect.Type {
	return reflect.TypeOf(new(float64))
}
