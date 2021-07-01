package responsewriter

import (
	"github.com/lyzzz123/illusion/handler/response"
	"io"
	"net/http"
	"reflect"
	"strconv"
)

var fileResponseType = reflect.TypeOf(new(response.FileResponse))

type FileResponseWriter struct {
}

func (fileResponseWriter *FileResponseWriter) Write(writer http.ResponseWriter, returnValue interface{}) error {

	fileReturnValue := returnValue.(*response.FileResponse)
	if fileReturnValue != nil {
		writer.Header().Set("Content-Disposition", "attachment; filename="+fileReturnValue.Name)
		writer.Header().Set("Content-Length", strconv.FormatInt(fileReturnValue.Size, 10))
		io.Copy(writer, fileReturnValue.Reader)
	}
	return nil
}

func (fileResponseWriter *FileResponseWriter) GetSupportResponseType() reflect.Type {
	return fileResponseType
}
