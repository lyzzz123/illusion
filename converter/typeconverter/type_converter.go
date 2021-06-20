package typeconverter

import (
	"reflect"
	"strconv"
)

var TypeConverterMap = make(map[reflect.Type]func(string) (interface{}, error))

func RegisterTypeConverter(typ reflect.Type, converterFunc func(string) (interface{}, error)) {
	if typ == nil {
		panic("RegisterTypeConverter typ must not be nil")
	}
	if converterFunc == nil {
		panic("RegisterTypeConverter converterFunc must not be nil")
	}
	TypeConverterMap[typ] = converterFunc
}

func IntConvert(param string) (interface{}, error) {

	if value, err := strconv.Atoi(param); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func Int8Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, 8); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func Int16Convert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, 16); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func Int32Convert(param string) (interface{}, error) {

	if value, err := strconv.ParseInt(param, 10, 32); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func Int64Convert(param string) (interface{}, error) {

	if value, err := strconv.ParseInt(param, 10, 64); err != nil {
		return nil, err
	} else {
		return value, nil
	}

}

func UintConvert(param string) (interface{}, error) {

	if value, err := strconv.ParseUint(param, 10, strconv.IntSize); err != nil {
		return nil, err
	} else {
		return uint(value), nil
	}

}

func Uint8Convert(param string) (interface{}, error) {

	if value, err := strconv.ParseUint(param, 10, 8); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func Uint16Convert(param string) (interface{}, error) {

	if value, err := strconv.ParseUint(param, 10, 16); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func Uint32Convert(param string) (interface{}, error) {

	if value, err := strconv.ParseUint(param, 10, 32); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func Uint64Convert(param string) (interface{}, error) {

	if value, err := strconv.ParseUint(param, 10, 64); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func BoolConvert(param string) (interface{}, error) {

	if value, err := strconv.ParseBool(param); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func Float64Convert(param string) (interface{}, error) {

	if value, err := strconv.ParseFloat(param, 64); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func Float32Convert(param string) (interface{}, error) {

	if value, err := strconv.ParseFloat(param, 32); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func StringConvert(param string) (interface{}, error) {
	return param, nil
}

func IntPtrConvert(param string) (interface{}, error) {

	if value, err := strconv.Atoi(param); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func Int8PtrConvert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, 8); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func Int16PtrConvert(param string) (interface{}, error) {
	if value, err := strconv.ParseInt(param, 10, 16); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func Int32PtrConvert(param string) (interface{}, error) {

	if value, err := strconv.ParseInt(param, 10, 32); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func Int64PtrConvert(param string) (interface{}, error) {

	if value, err := strconv.ParseInt(param, 10, 64); err != nil {
		return nil, err
	} else {
		return &value, nil
	}

}

func UintPtrConvert(param string) (interface{}, error) {

	if value, err := strconv.ParseUint(param, 10, strconv.IntSize); err != nil {
		return nil, err
	} else {
		v := uint(value)
		return &v, nil
	}

}

func Uint8PtrConvert(param string) (interface{}, error) {

	if value, err := strconv.ParseUint(param, 10, 8); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func Uint16PtrConvert(param string) (interface{}, error) {

	if value, err := strconv.ParseUint(param, 10, 16); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func Uint32PtrConvert(param string) (interface{}, error) {

	if value, err := strconv.ParseUint(param, 10, 32); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func Uint64PtrConvert(param string) (interface{}, error) {

	if value, err := strconv.ParseUint(param, 10, 64); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func BoolPtrConvert(param string) (interface{}, error) {

	if value, err := strconv.ParseBool(param); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func Float64PtrConvert(param string) (interface{}, error) {

	if value, err := strconv.ParseFloat(param, 64); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func Float32PtrConvert(param string) (interface{}, error) {

	if value, err := strconv.ParseFloat(param, 32); err != nil {
		return nil, err
	} else {
		return &value, nil
	}
}

func StringPtrConvert(param string) (interface{}, error) {
	return &param, nil
}
