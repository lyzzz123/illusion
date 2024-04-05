package converter

import "reflect"

type StringPtrConvert struct {
}

func (stringPtrConvert *StringPtrConvert) Convert(param string) (interface{}, error) {
	return &param, nil
}

func (stringPtrConvert *StringPtrConvert) Support() reflect.Type {
	return reflect.TypeOf(new(string))
}
