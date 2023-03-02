package stream

import "context"

type Header interface {
	Get(key string) string
	Set(key string, val string)
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
	Ack() bool
	// Nack nack message
	Nack() bool
	// Acked returns channel which is closed when acknowledgement is sent.
	Acked() <-chan struct{}
	// Nacked returns channel which is closed when negative acknowledgement is sent.
	Nacked() <-chan struct{}
}
