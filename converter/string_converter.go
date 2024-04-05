package converter

import "reflect"

type StringConvert struct {
}

func (stringConvert *StringConvert) Convert(param string) (interface{}, error) {
	return param, nil
}

func (stringConvert *StringConvert) Support() reflect.Type {
	return reflect.TypeOf(*new(string))
}
