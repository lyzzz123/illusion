package converter

import (
	"reflect"
	"strconv"
)

type BoolPtrConvert struct {
}

func (boolPtrConvert *BoolPtrConvert) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseBool(param); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func (boolPtrConvert *BoolPtrConvert) Support() reflect.Type {
	return reflect.TypeOf(new(bool))
}
