package handler

import (
	"errors"
	"github.com/lyzzz123/illusion/converter/requestconverter"
	"github.com/lyzzz123/illusion/converter/responsewriter"
	"github.com/lyzzz123/illusion/handler/exceptionhandler"
	"github.com/lyzzz123/illusion/log"
	"net/http"
	"reflect"
)

func Handle(writer http.ResponseWriter, request *http.Request) error {
	pc := requestconverter.GetRequestConverter(request)
	if pc == nil {
		return errors.New("not found request converter")
	}
	hw := getHandler(request.Method, request.URL.Path)
	if hw == nil {
		return errors.New("not found handler")
	}
	param := reflect.New(hw.InputType).Interface()
	err := pc.Convert(writer, request, param, hw)
	if err != nil {
		return err
	}
	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf(param)
	result := reflect.ValueOf(hw.InnerHandler).Call(args)
	if result[1].Interface() != nil {
		if err := exceptionhandler.ExceptionHandlerInstance.Handle(writer, result[1].Interface().(error)); err != nil {
			return err
		}
		return nil
	}
	returnValue := result[0].Interface()
	responseWriter := responsewriter.GetResponseWriter(returnValue)
	if responseWriter == nil {

		return errors.New("not found responseWriter")
	}
	err = responseWriter.Write(writer, returnValue)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
