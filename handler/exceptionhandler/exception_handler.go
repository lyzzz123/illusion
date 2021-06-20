package exceptionhandler

import "net/http"

var ExceptionHandlerInstance ExceptionHandler

func RegisterExceptionHandler(exceptionHandler ExceptionHandler) {
	ExceptionHandlerInstance = exceptionHandler
}

type ExceptionHandler interface {
	Handle(writer http.ResponseWriter, err error) error
}

type DefaultExceptionHandler struct {
}

func (defaultExceptionHandler *DefaultExceptionHandler) Handle(writer http.ResponseWriter, err error) error {
	writer.WriteHeader(500)
	return nil
}
