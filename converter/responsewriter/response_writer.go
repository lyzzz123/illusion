package responsewriter

import (
	"net/http"
)

var ResponseWriterMap = make(map[string]ResponseWriter)

type ResponseWriter interface {
	Write(writer http.ResponseWriter, returnValue interface{}) error

	Support(returnValue interface{}) bool

	Name() string
}

func RegisterResponseWriter(responseWriter ResponseWriter) {
	if responseWriter.Name() == "" {
		panic("ResponseWriter name must not be blank")
	}
	ResponseWriterMap[responseWriter.Name()] = responseWriter

}

func GetResponseWriter(returnValue interface{}) ResponseWriter {
	if len(ResponseWriterMap) == 0 {
		return nil
	}

	for _, v := range ResponseWriterMap {
		if v.Support(returnValue) {
			return v
		}
	}
	return nil
}
