package converter

import (
	"reflect"
	"strconv"
)

type BoolConvert struct {
}

func (boolConvert *BoolConvert) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseBool(param); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func (boolConvert *BoolConvert) Support() reflect.Type {
	return reflect.TypeOf(*new(bool))
}
