package requestconverter

import (
	"github.com/lyzzz123/illusion/handler/handlerwrapper"
	"net/http"
)

type GetMethodConverter struct {
}

func (getMethodConverter *GetMethodConverter) Convert(writer http.ResponseWriter, request *http.Request, param interface{}, hw *handlerwrapper.HandlerWrapper) error {
	request.ParseForm()
	FillInParamValue(request.Form, param, hw)
	FillInPathValue(request, param, hw)

	return nil
}

func (getMethodConverter *GetMethodConverter) Support(request *http.Request) bool {
	if request.Method == "GET" {
		return true
	}

	return false
}

func (getMethodConverter *GetMethodConverter) Name() string {
	return "GetMethodConverter"
}
