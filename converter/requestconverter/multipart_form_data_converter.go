package requestconverter

import (
	"encoding/json"
	"github.com/lyzzz123/illusion/handler/handlerwrapper"
	"net/http"
	"reflect"
	"strings"
)

type MultipartFormDataConverter struct {
}

func (multipartFormDataConverter *MultipartFormDataConverter) Convert(writer http.ResponseWriter, request *http.Request, param interface{}, hw *handlerwrapper.HandlerWrapper) error {

	if err := request.ParseMultipartForm(1 << 20); err != nil {
		return err
	}

	valueMap := request.MultipartForm.Value
	tempJsonMap := make(map[string]interface{})
	for k, v := range valueMap {
		if len(v) == 1 {
			tempJsonMap[k] = v[0]
		} else {
			tempJsonMap[k] = v
		}
	}
	if bytes, err := json.Marshal(tempJsonMap); err != nil {
		return err
	} else {
		json.Unmarshal(bytes, param)
	}

	pv := reflect.ValueOf(param)
	fileMap := request.MultipartForm.File
	for k, v := range fileMap {
		for i := 0; i < pv.Type().Elem().NumField(); i++ {
			tagName := pv.Type().Elem().Field(i).Tag.Get("json")
			if k == tagName {

				if len(v) == 1 {
					pv.Elem().Field(i).Set(reflect.ValueOf(v[0]))
				} else {
					pv.Elem().Field(i).Set(reflect.ValueOf(v))
				}

			}
		}
	}

	return nil
}

func (multipartFormDataConverter *MultipartFormDataConverter) Support(request *http.Request) bool {

	if strings.HasPrefix(request.Header.Get("Content-Type"), "multipart/form-data") {
		return true
	}

	return false
}

func (multipartFormDataConverter *MultipartFormDataConverter) Name() string {
	return "MultipartFormDataConverter"
}
