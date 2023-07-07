package gochan

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"context"
)

type Marshaller interface {
	Marshal(ctx context.Context, topic string, msg *stream.Message) ([]byte, error)
	Unmarshal(goChanMsg []byte) (*stream.Message, error)
}

type DefaultMarshaller struct{}

func (d DefaultMarshaller) Marshal(ctx context.Context, topic string, msg *stream.Message) ([]byte, error) {
	return msg.Payload, nil
}

func (d DefaultMarshaller) Unmarshal(goChanMsg []byte) (*stream.Message, error) {
	return &stream.Message{Payload: goChanMsg}, nil
}
