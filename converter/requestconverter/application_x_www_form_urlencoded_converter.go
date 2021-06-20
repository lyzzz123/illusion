package requestconverter

import (
	"github.com/lyzzz123/illusion/handler/handlerwrapper"
	"net/http"
)

type ApplicationXWWWFormUrlencodedConverter struct {
}

func (applicationXWWWFormUrlencodedConverter *ApplicationXWWWFormUrlencodedConverter) Convert(writer http.ResponseWriter, request *http.Request, param interface{}, hw *handlerwrapper.HandlerWrapper) error {
	request.ParseForm()
	FillInParamValue(request.Form, param, hw)
	FillInPathValue(request, param, hw)

	return nil
}

func (applicationXWWWFormUrlencodedConverter *ApplicationXWWWFormUrlencodedConverter) Support(request *http.Request) bool {
	if request.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		return true
	}

	return false
}

func (applicationXWWWFormUrlencodedConverter *ApplicationXWWWFormUrlencodedConverter) Name() string {
	return "ApplicationXWWWFormUrlencodedConverter"
}
