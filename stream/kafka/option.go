package kafka

import (
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"time"
)

type options struct {
	Marshaler       Marshaler
	ShutdownTimeout time.Duration
	RebalanceCb     kafka.RebalanceCb
	NackHandler     func(msg *stream.Message)
	PollTimeout     time.Duration
}

type Option func(o *options)

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.Marshaler == nil {
		o.Marshaler = DefaultMarshaler{}
	}
	if o.ShutdownTimeout <= 0 {
		o.ShutdownTimeout = 10 * time.Second
	}
	if o.PollTimeout <= 0 {
		o.PollTimeout = 100 * time.Millisecond
	}
}

func MessageMarshaler(m Marshaler) Option {
	return func(o *options) {
		o.Marshaler = m
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ShutdownTimeout = timeout
	}
}

func RebalanceCb(cb kafka.RebalanceCb) Option {
	return func(o *options) {
		o.RebalanceCb = cb
	}
}

func PollTimeout(t time.Duration) Option {
	return func(o *options) {
		o.PollTimeout = t
	}
}

func NackHandler(h func(msg *stream.Message)) Option {
	return func(o *options) {
		o.NackHandler = h
	}
}
