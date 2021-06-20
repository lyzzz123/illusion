package requestconverter

import (
	"illusion/converter/typeconverter"
	"illusion/handler/handlerwrapper"
	"net/http"
	"reflect"
)

var RequestConverterMap = make(map[string]RequestConverter)

type RequestConverter interface {
	Convert(writer http.ResponseWriter, request *http.Request, param interface{}, hw *handlerwrapper.HandlerWrapper) error

	Support(request *http.Request) bool

	Name() string
}

func RegisterRequestConverter(requestConverter RequestConverter) {
	if requestConverter == nil {
		panic("RegisterRequestConverter requestConverter must not be nil")
	}
	if requestConverter.Name() == "" {
		panic(" requestConverter name must not be blank")
	}

	RequestConverterMap[requestConverter.Name()] = requestConverter
}

func GetRequestConverter(request *http.Request) RequestConverter {

	if len(RequestConverterMap) == 0 {
		return nil
	}

	for _, v := range RequestConverterMap {
		if v.Support(request) {
			return v
		}
	}

	return nil
}

func FillInPathValue(request *http.Request, param interface{}, hw *handlerwrapper.HandlerWrapper) error {

	if len(hw.PathValuePositionMap) == 0 {
		return nil
	}

	paramValues := hw.PathRegexForPathValue.FindStringSubmatch(request.URL.Path)
	paramReflect := reflect.ValueOf(param).Elem()
	for i := 1; i < len(paramValues); i++ {

		if index, ok := hw.PathValuePositionMap[i]; ok {
			field := paramReflect.Field(index)

			if method, ok := typeconverter.TypeConverterMap[field.Type()]; ok {
				if parsedValue, err := method(paramValues[i]); err != nil {
					return err
				} else {
					field.Set(reflect.ValueOf(parsedValue))
				}
			}
		}
	}
	return nil

}

func FillInParamValue(paramMap map[string][]string, param interface{}, hw *handlerwrapper.HandlerWrapper) error {
	pe := reflect.ValueOf(param).Elem()
	for k, v := range paramMap {
		index, ok := hw.ParamValuePositionMap[k]
		if !ok {
			continue
		}
		field := pe.Field(index)
		if field.Kind() == reflect.Slice {
			va := reflect.New(field.Type()).Elem()
			if method, ok := typeconverter.TypeConverterMap[field.Type().Elem()]; ok {
				for i := 0; i < len(v); i++ {
					if revertedValue, err := method(v[i]); err != nil {
						return err
					} else {
						va = reflect.Append(va, reflect.ValueOf(revertedValue))

					}
				}
			}
			field.Set(va)
		} else {
			if method, ok := typeconverter.TypeConverterMap[field.Type()]; ok {
				if revertedValue, err := method(v[0]); err != nil {
					return err
				} else {
					field.Set(reflect.ValueOf(revertedValue))
				}
			}
		}
	}
	return nil
}
