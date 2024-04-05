package converter

import (
	"reflect"
	"strconv"
)

type IntPtrConverter struct {
}

func (intPtrConverter *IntPtrConverter) Convert(param string) (interface{}, error) {
	if value, err := strconv.Atoi(param); err != nil {
		return nil, err
	} else {
		i := int(value)
		return &i, nil
	}
}

func (intPtrConverter *IntPtrConverter) Support() reflect.Type {
	return reflect.TypeOf(new(int))
}
