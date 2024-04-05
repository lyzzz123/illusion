package converter

import (
	"reflect"
	"strconv"
)

type Int16PtrConverter struct {
}

func (int16PtrConverter *Int16PtrConverter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, 16); err != nil {
		return nil, err
	} else {
		i16 := int16(value)
		return &i16, nil
	}
}

func (int16PtrConverter *Int16PtrConverter) Support() reflect.Type {
	return reflect.TypeOf(new(int16))
}
