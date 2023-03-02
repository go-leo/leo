package stream

import "context"

type Header interface {
	Get(key string) string
	Set(key string, val string)
	Append(key string, val string)
}

type Message interface {
	// ID return message uniq id
	ID() string
	// Payload return message payload
	Payload() []byte
	// Header return message header
	Header() Header
	// Context return message context
	Context() context.Context
	// WithContext return new message with context
	WithContext(ctx context.Context) Message
	// Ack ack message
	Ack() (any, error)
	// Nack nack message
	Nack() (any, error)
	// Acked returns a channel which is closed when ack message. function will be invoked at Ack
	Acked(f func() (any, error)) <-chan struct{}
	// Nacked returns a channel which is closed when nack message. function will be invoked at Nack
	Nacked(f func() (any, error)) <-chan struct{}
}
