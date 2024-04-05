package converter

import (
	"reflect"
	"strconv"
)

type Float64Convert struct {
}

func (float64Convert *Float64Convert) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseFloat(param, 64); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func (float64Convert *Float64Convert) Support() reflect.Type {
	return reflect.TypeOf(*new(float64))
}
