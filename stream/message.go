package stream

import "context"

type Header interface {
	Get(key string) string
	Set(key string, val string)
}

type Message interface {
	ID() string
	Payload() []byte
	Header() Header
	Context() context.Context
	WithContext(ctx context.Context) Message
}
