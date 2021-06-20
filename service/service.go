package service

import (
	"illusion/filter"
	"illusion/handler"
	"illusion/log"
	"net/http"
)

type IllusionService struct {
}

func (illusionService *IllusionService) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if err := filter.ExecutePreHandle(writer, request); err != nil {
		log.Error("execute fiter preHandle error:%s", err.Error())
		writer.WriteHeader(500)
	}

	if err := handler.Handle(writer, request); err != nil {
		log.Error("execute handler error:%s", err.Error())
		writer.WriteHeader(500)
		//	io.WriteString(writer, err.Error())
		return
	}

	if err := filter.ExecutePostHandle(writer, request); err != nil {
		log.Error("execute fiter postHandle error:%s", err.Error())
	}
}
