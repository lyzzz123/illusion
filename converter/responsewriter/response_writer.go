package responsewriter

import (
	"net/http"
	"reflect"
)

var ResponseWriterMap = make(map[reflect.Type]ResponseWriter)

type ResponseWriter interface {
	Write(writer http.ResponseWriter, returnValue interface{}) error

	GetSupportResponseType() reflect.Type
}

func RegisterResponseWriter(responseWriter ResponseWriter) {
	if responseWriter.GetSupportResponseType() == nil {
		panic("ResponseWriter SupportResponseType must not be nil")
	}
	ResponseWriterMap[responseWriter.GetSupportResponseType()] = responseWriter

}

func GetResponseWriter(returnValue interface{}) ResponseWriter {
	if len(ResponseWriterMap) == 0 {
		return nil
	}
	rw := ResponseWriterMap[reflect.TypeOf(returnValue)]
	return rw
}
