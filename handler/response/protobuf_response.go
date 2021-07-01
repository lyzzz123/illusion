package response

import "github.com/golang/protobuf/proto"

type ProtobufResponse struct {
	Data proto.Message
}
