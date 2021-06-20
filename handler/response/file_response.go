package response

import "io"

type FileResponse struct {
	Name   string
	Size   int64
	Reader io.Reader
}
