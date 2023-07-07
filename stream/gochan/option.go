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
	Marshaller  Marshaller
	NackHandler func(msg *stream.Message)
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

func NackHandler(h func(msg *stream.Message)) Option {
	return func(o *options) {
		o.NackHandler = h
	}
}
