package stream

import (
	"time"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

type options struct {
	Handlers          []Handler
	PubSubHandlers    []PubSubHandler
	Interceptors      []Interceptor
	MessageBufferSize int
	ErrorHandler      func(err error)
	Logger            log.Logger
	ShutdownTimeout   time.Duration
}

type Option func(o *options)

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.MessageBufferSize <= 0 {
		o.MessageBufferSize = 1
	}
	if o.ErrorHandler == nil {
		o.ErrorHandler = func(err error) {}
	}
	if o.Logger == nil {
		o.Logger = log.L()
	}
}

func Handlers(h ...Handler) Option {
	return func(o *options) {
		o.Handlers = append(o.Handlers, h...)
	}
}

func PubSubHandlers(h ...PubSubHandler) Option {
	return func(o *options) {
		o.PubSubHandlers = append(o.PubSubHandlers, h...)
	}
}

func Interceptors(f ...Interceptor) Option {
	return func(o *options) {
		o.Interceptors = append(o.Interceptors, f...)
	}
}

func MessageBufferSize(size int) Option {
	return func(o *options) {
		o.MessageBufferSize = size
	}
}

func ErrorHandler(f func(err error)) Option {
	return func(o *options) {
		o.ErrorHandler = f
	}
}

func Logger(l log.Logger) Option {
	return func(o *options) {
		o.Logger = l
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ShutdownTimeout = timeout
	}
}
