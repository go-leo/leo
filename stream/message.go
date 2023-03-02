package stream

import "context"

// Header is a mapping from keys to values.
type Header interface {
	// Len returns the number of items in Header.
	Len() int
	// Get obtains the values for a given key.
	Get(key string) []string
	// Set sets the value of a given key with a slice of values.
	Set(key string, values ...string) Header
	// Append appends the value of a given key with a slice of values.
	Append(k string, values ...string) Header
	// Delete delete the values for a given key.
	Delete(key string) Header
	// Range iterates the header
	Range(fn func(key string, values []string)) Header
}

// Message is a message. Message is emitted by Publisher and received by Subscriber.
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
