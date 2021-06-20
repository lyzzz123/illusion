package handler

import (
	"errors"
	"illusion/constant/httpmethod"
	"illusion/handler/handlerwrapper"
	"illusion/utils"
	"reflect"
	"regexp"
)

var pathHandlerMapping = make(map[string]*[]*handlerwrapper.HandlerWrapper)

var handlerMapping = make(map[string]*map[string]*handlerwrapper.HandlerWrapper)

func init() {
	getPathHandlerMapping := make([]*handlerwrapper.HandlerWrapper, 0)
	pathHandlerMapping[httpmethod.GET] = &getPathHandlerMapping

	postPathHandlerMapping := make([]*handlerwrapper.HandlerWrapper, 0)
	pathHandlerMapping[httpmethod.POST] = &postPathHandlerMapping

	putPathHandlerMapping := make([]*handlerwrapper.HandlerWrapper, 0)
	pathHandlerMapping[httpmethod.PUT] = &putPathHandlerMapping

	deletePathHandlerMapping := make([]*handlerwrapper.HandlerWrapper, 0)
	pathHandlerMapping[httpmethod.DELETE] = &deletePathHandlerMapping

	getHandlerMapping := make(map[string]*handlerwrapper.HandlerWrapper)
	handlerMapping[httpmethod.GET] = &getHandlerMapping

	postHandlerMapping := make(map[string]*handlerwrapper.HandlerWrapper)
	handlerMapping[httpmethod.POST] = &postHandlerMapping

	putHandlerMapping := make(map[string]*handlerwrapper.HandlerWrapper)
	handlerMapping[httpmethod.PUT] = &putHandlerMapping

	deleteHandlerMapping := make(map[string]*handlerwrapper.HandlerWrapper)
	handlerMapping[httpmethod.DELETE] = &deleteHandlerMapping
}

func getHandler(method string, path string) *handlerwrapper.HandlerWrapper {

	hm, ok := handlerMapping[method]
	if !ok {
		return nil
	}

	hw := (*hm)[path]
	if hw != nil {
		return hw
	}
	phm, ok := pathHandlerMapping[method]
	if !ok {
		return nil
	}

	for i := 0; i < len(*phm); i++ {
		hw = (*phm)[i]
		if hw.PathRegex.MatchString(path) {
			return hw
		}
	}
	return nil
}

func parsePathValueMap(pathRegexForPath *regexp.Regexp, path string, paramType reflect.Type) map[int]int {

	pathParamMap := make(map[string]int)

	pathValueParamNames := pathRegexForPath.FindStringSubmatch(path)
	if len(pathValueParamNames) > 1 {
		for i := 1; i < len(pathValueParamNames); i++ {
			pathParamMap[pathValueParamNames[i]] = i
		}
	}

	pathParamPositionMap := make(map[int]int)

	for i := 0; i < paramType.NumField(); i++ {
		pathValue := paramType.Field(i).Tag.Get("pathValue")
		if pathValue != "" {
			if index, ok := pathParamMap[pathValue]; ok {
				pathParamPositionMap[index] = i
			}
		}
	}

	return pathParamPositionMap
}

func getInputType(handler interface{}) (reflect.Type, error) {
	methodType := reflect.ValueOf(handler).Type()
	if methodType.NumIn() != 1 {
		return nil, errors.New(methodType.String() + " must has one input param")
	}
	argType := methodType.In(0)
	if argType.Kind() != reflect.Ptr {
		return nil, errors.New(methodType.String() + " must has one ptr input param")
	} else {
		if argType.Elem().Kind() != reflect.Struct {
			return nil, errors.New(methodType.String() + " must has one struct ptr input param")
		}
	}
	return argType.Elem(), nil
}

func getOutputType(handler interface{}) (reflect.Type, error) {
	methodType := reflect.ValueOf(handler).Type()
	if methodType.NumOut() != 2 {
		return nil, errors.New(methodType.String() + " must has two output param")
	}
	firstOutputType := methodType.Out(0)
	if firstOutputType.Kind() != reflect.Ptr {
		return nil, errors.New(methodType.String() + " first output param must be ptr")
	} else {
		if firstOutputType.Elem().Kind() != reflect.Struct {
			return nil, errors.New(methodType.String() + " first output param must be struct ptr")
		}
	}

	secondOutputType := methodType.Out(1)

	if secondOutputType.String() != "error" {
		return nil, errors.New(methodType.String() + " second output param must be error")
	}

	return firstOutputType.Elem(), nil
}

func parseParamValuePositionMap(paramType reflect.Type) map[string]int {
	paramMap := make(map[string]int)
	for i := 0; i < paramType.NumField(); i++ {
		pathValue := paramType.Field(i).Tag.Get("paramValue")
		if pathValue != "" {
			paramMap[pathValue] = i
		}
	}
	return paramMap
}

func RegisterHandler(path string, httpMethod []string, handlerMethod interface{}) {

	if path == "" {
		panic("RegisterHandler path must not be blank")
	}

	if httpMethod == nil {
		panic("RegisterHandler httpMethod must not be nil")
	}

	if len(httpMethod) == 0 {
		panic("RegisterHandler httpMethod length not be zero")
	}

	if handlerMethod == nil {
		panic("RegisterHandler handlerMethod not be nil")
	}

	hw := &handlerwrapper.HandlerWrapper{}
	hw.Path = path

	hw.PathRegex = utils.ParseHandlerUrlPattern(path)
	hw.PathRegexForPath = utils.ParseHandlerUrlPatternForPath(path)
	hw.PathRegexForPathValue = utils.ParseHandlerUrlPatternForPathValue(path)
	hw.InnerHandler = handlerMethod
	if pt, err := getInputType(handlerMethod); err != nil {
		panic(err)
	} else {
		hw.InputType = pt
	}
	if ot, err := getOutputType(handlerMethod); err != nil {
		panic(err)
	} else {
		hw.OutputType = ot
	}

	hw.PathValuePositionMap = parsePathValueMap(hw.PathRegexForPath, path, hw.InputType)
	hw.ParamValuePositionMap = parseParamValuePositionMap(hw.InputType)
	hw.HttpMethod = httpMethod

	for _, v := range httpMethod {
		hm, ok := handlerMapping[v]
		if !ok {
			panic("don't support http method " + v)
		}
		phm, _ := pathHandlerMapping[v]
		if len(hw.PathValuePositionMap) == 0 {
			(*hm)[path] = hw
		} else {
			(*phm) = append((*phm), hw)
		}

	}
}
