package converter

import (
	"reflect"
	"strconv"
)

type Float32Convert struct {
}

func (float32Convert *Float32Convert) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseFloat(param, 32); err != nil {
		return nil, err
	} else {
		return float32(value), nil
	}
}

func (float32Convert *Float32Convert) Support() reflect.Type {
	return reflect.TypeOf(*new(float32))
}
