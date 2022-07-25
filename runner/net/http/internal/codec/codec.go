package codec

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var JSONCodec = json{}

var ProtobufCodec = protobuf{}

// Codec defines the interface leo uses to encode and decode messages.  Note
// that implementations of this interface must be thread safe; a Codec's
// methods can be called from concurrent goroutines.
type Codec interface {
	// Marshal returns the wire format of v.
	Marshal(v any) ([]byte, error)
	// Unmarshal parses the wire format into v.
	Unmarshal(data []byte, v any) error
	// Name returns the name of the Codec implementation. The returned string
	// will be used as part of content type in transmission.  The result must be
	// static; the result cannot change between calls.
	Name() string
	// ContentType returns the Content-Type which this marshaler is responsible for.
	ContentType() string
}

type json struct{}

func (json) Marshal(v any) ([]byte, error) {
	vv, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("failed to marshal, message is %T, want proto.Message", v)
	}
	return protojson.Marshal(vv)
}

func (json) Unmarshal(data []byte, v any) error {
	vv, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("failed to unmarshal, message is %T, want proto.Message", v)
	}
	return protojson.Unmarshal(data, vv)
}

func (json) Name() string {
	return "json"
}

func (json) ContentType() string {
	return "application/json"
}

type protobuf struct{}

func (protobuf) Marshal(v any) ([]byte, error) {
	vv, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("failed to marshal, message is %T, want proto.Message", v)
	}
	return proto.Marshal(vv)
}

func (protobuf) Unmarshal(data []byte, v any) error {
	vv, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("failed to unmarshal, message is %T, want proto.Message", v)
	}
	return proto.Unmarshal(data, vv)
}

func (protobuf) Name() string {
	return "protobuf"
}

func (protobuf) ContentType() string {
	return "application/x-protobuf"
}

func GetCodec(contentType string) Codec {
	switch contentType {
	case "application/json":
		return JSONCodec
	case "application/x-protobuf":
		return ProtobufCodec
	default:
		return JSONCodec
	}
}
