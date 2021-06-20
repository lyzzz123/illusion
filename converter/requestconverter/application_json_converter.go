package requestconverter

import (
	"encoding/json"
	"illusion/handler/handlerwrapper"
	"io/ioutil"
	"net/http"
)

type ApplicationJSONConverter struct {
}

func (applicationJSONConverter *ApplicationJSONConverter) Convert(writer http.ResponseWriter, request *http.Request, param interface{}, hw *handlerwrapper.HandlerWrapper) error {

	if bodyBytes, err := ioutil.ReadAll(request.Body); err != nil {
		return err
	} else {
		json.Unmarshal(bodyBytes, param)
	}
	FillInPathValue(request, param, hw)
	return nil
}

func (applicationJSONConverter *ApplicationJSONConverter) Support(request *http.Request) bool {

	if request.Header.Get("Content-Type") == "application/json" {
		return true
	}

	return false
}

func (applicationJSONConverter *ApplicationJSONConverter) Name() string {
	return "ApplicationJSONConverter"
}
