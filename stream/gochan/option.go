package gochan

import (
	"codeup.aliyun.com/qimao/leo/leo/log"
	"codeup.aliyun.com/qimao/leo/leo/stream"
)

type options struct {
	Logger      log.Logger
	NackHandler func(msg *stream.Message)
}

type Option func(o *options)

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {}

func Logger(l log.Logger) Option {
	return func(o *options) {
		o.Logger = l
	}
}

func NackHandler(h func(msg *stream.Message)) Option {
	return func(o *options) {
		o.NackHandler = h
	}
}
