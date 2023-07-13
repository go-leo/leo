package kafka

import (
	"time"

	"codeup.aliyun.com/qimao/leo/leo/log"
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type options struct {
	Logger            log.Logger
	Marshaller        Marshaller
	NackHandler       func(msg *stream.Message)
	OnMessageSending  func(*stream.Message, *kafka.Message) *kafka.Message
	OnMessageReceived func(*stream.Message, *kafka.Message) *stream.Message

	ShutdownTimeout time.Duration
	RebalanceCb     kafka.RebalanceCb
	PollTimeout     time.Duration
}

type Option func(o *options)

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.Marshaller == nil {
		o.Marshaller = DefaultMarshaller{}
	}
	if o.ShutdownTimeout <= 0 {
		o.ShutdownTimeout = 10 * time.Second
	}
	if o.PollTimeout <= 0 {
		o.PollTimeout = 100 * time.Millisecond
	}
}

func Logger(l log.Logger) Option {
	return func(o *options) {
		o.Logger = l
	}
}

func MessageMarshaller(m Marshaller) Option {
	return func(o *options) {
		o.Marshaller = m
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

func OnMessageSending(f func(*stream.Message, *kafka.Message) *kafka.Message) Option {
	return func(o *options) {
		o.OnMessageSending = f
	}
}

func OnMessageReceived(f func(*stream.Message, *kafka.Message) *stream.Message) Option {
	return func(o *options) {
		o.OnMessageReceived = f
	}
}
