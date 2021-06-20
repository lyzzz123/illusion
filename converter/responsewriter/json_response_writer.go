package responsewriter

import (
	"encoding/json"
	"illusion/handler/response"
	"io"
	"net/http"
	"reflect"
)

var jsonResponseType = reflect.TypeOf(new(response.JSONResponse))

type JSONResponseWriter struct {
}

func (jsonResponseWriter JSONResponseWriter) Write(writer http.ResponseWriter, returnValue interface{}) error {
	writer.Header().Set("Content-Type", "application/json")

	if returnValue == nil {
		return nil
	}

	jsonStringResponse := returnValue.(*response.JSONResponse)

	if bytes, err := json.Marshal(jsonStringResponse.Data); err != nil {
		return err
	} else {
		io.WriteString(writer, string(bytes))
	}

	return nil
}

func (jsonResponseWriter JSONResponseWriter) Support(returnValue interface{}) bool {

	if reflect.ValueOf(returnValue).Type() == jsonResponseType {
		return true
	}
	return false
}

func (jsonResponseWriter JSONResponseWriter) Name() string {
	return "JSONResponseWriter"
}
