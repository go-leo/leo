package http

import (
	"net/http"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/log"
	"codeup.aliyun.com/qimao/leo/leo/stream"
)

type options struct {
	Logger            log.Logger
	Marshaller        Marshaller
	NackHandler       func(msg *stream.Message)
	OnMessageSending  func(*stream.Message, *http.Request) *http.Request
	OnMessageReceived func(*stream.Message, *http.Response) *stream.Message

	HttpClient      *http.Client
	Method          func(topic string) string
	URL             func(topic string) string
	ShutdownTimeout time.Duration
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
	if o.HttpClient == nil {
		o.HttpClient = http.DefaultClient
	}
	if o.ShutdownTimeout <= 0 {
		o.ShutdownTimeout = 10 * time.Second
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

func Method(f func(topic string) string) Option {
	return func(o *options) {
		o.Method = f
	}
}

func URL(f func(topic string) string) Option {
	return func(o *options) {
		o.URL = f
	}
}

func OnMessageSending(f func(*stream.Message, *http.Request) *http.Request) Option {
	return func(o *options) {
		o.OnMessageSending = f
	}
}

func OnMessageReceived(f func(*stream.Message, *http.Response) *stream.Message) Option {
	return func(o *options) {
		o.OnMessageReceived = f
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ShutdownTimeout = timeout
	}
}

func NackHandler(h func(msg *stream.Message)) Option {
	return func(o *options) {
		o.NackHandler = h
	}
}
