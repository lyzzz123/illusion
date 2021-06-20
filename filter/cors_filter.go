package filter

import (
	"net/http"
)

type CorsFilter struct {
}

func (corsFilter *CorsFilter) PreHandle(writer http.ResponseWriter, request *http.Request) error {

	return nil
}

func (corsFilter *CorsFilter) PostHandle(writer http.ResponseWriter) error {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, DELETE, PUT")
	writer.Header().Set("Access-Control-Max-Age", "3600")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	return nil
}

func (corsFilter *CorsFilter) GetPriority() int {
	return 5
}

func (corsFilter *CorsFilter) GetPathPattern() string {
	return "/**"
}
