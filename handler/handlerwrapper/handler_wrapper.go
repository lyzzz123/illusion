package handlerwrapper

import (
	"reflect"
	"regexp"
)

type HandlerWrapper struct {
	Path                  string
	PathRegex             *regexp.Regexp
	PathRegexForPath      *regexp.Regexp
	PathRegexForPathValue *regexp.Regexp
	InnerHandler          interface{}
	PathValuePositionMap  map[int]int
	ParamValuePositionMap map[string]int
	HttpMethod            []string
	InputType             reflect.Type
	OutputType            reflect.Type
}
