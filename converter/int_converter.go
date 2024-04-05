package converter

import (
	"reflect"
	"strconv"
)

type IntConverter struct {
}

func (intConverter *IntConverter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, strconv.IntSize); err != nil {
		return nil, err
	} else {
		return int(value), nil
	}
}

func (intConverter *IntConverter) Support() reflect.Type {
	return reflect.TypeOf(*new(int))
}
