package gochan

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"context"
)

type Marshaler interface {
	Marshal(ctx context.Context, topic string, msg *stream.Message) ([]byte, error)
	Unmarshal(goChanMsg []byte) (*stream.Message, error)
}

type DefaultMarshaler struct{}

func (d DefaultMarshaler) Marshal(ctx context.Context, topic string, msg *stream.Message) ([]byte, error) {
	return msg.Payload, nil
}

func (d DefaultMarshaler) Unmarshal(goChanMsg []byte) (*stream.Message, error) {
	return &stream.Message{Payload: goChanMsg}, nil
}
