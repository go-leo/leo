package gochan

import (
	"codeup.aliyun.com/qimao/leo/leo/log"
	"codeup.aliyun.com/qimao/leo/leo/stream"
)

//
//import (
//	"codeup.aliyun.com/qimao/leo/leo/stream"
//	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
//	"time"
//)

type options struct {
	Logger      log.Logger
	Marshaler   Marshaler
	NackHandler func(msg *stream.Message)
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
}

func Logger(l log.Logger) Option {
	return func(o *options) {
		o.Logger = l
	}
}

func MessageMarshaler(m Marshaler) Option {
	return func(o *options) {
		o.Marshaler = m
	}
}

func NackHandler(h func(msg *stream.Message)) Option {
	return func(o *options) {
		o.NackHandler = h
	}
}
