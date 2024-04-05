package converter

import (
	"reflect"
	"strconv"
)

type Int64PtrConverter struct {
}

func (int64PtrConverter *Int64PtrConverter) Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, 64); err != nil {
		return nil, err
	} else {
		i64 := int64(value)
		return &i64, nil
	}
}

func (int64PtrConverter *Int64PtrConverter) Support() reflect.Type {
	return reflect.TypeOf(new(int64))
}
