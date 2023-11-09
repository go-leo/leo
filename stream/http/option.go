package http

import (
	"errors"
	"net/http"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/log"
	"codeup.aliyun.com/qimao/leo/leo/stream"
)

var ErrServerIsNil = errors.New("http.Server and http.ServeMux is nil")

type options struct {
	Logger            log.Logger
	Marshaller        Marshaller
	NackHandler       func(msg *stream.Message)
	OnMessageSending  func(*stream.Message, *http.Request) *http.Request
	OnMessageReceived func(*stream.Message, *http.Request) *stream.Message

	AckResponse  func(resp http.ResponseWriter, req *http.Request, msg *stream.Message) (any, error)
	NackResponse func(resp http.ResponseWriter, req *http.Request, msg *stream.Message) (any, error)
	HttpClient   *http.Client
	ServeMux     *http.ServeMux
	HttpServer   *http.Server

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

func OnMessageSending(f func(*stream.Message, *http.Request) *http.Request) Option {
	return func(o *options) {
		o.OnMessageSending = f
	}
}

func OnMessageReceived(f func(*stream.Message, *http.Request) *stream.Message) Option {
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

func AckResponse(h func(resp http.ResponseWriter, req *http.Request, msg *stream.Message) (any, error)) Option {
	return func(o *options) {
		o.AckResponse = h
	}
}
func NackResponse(h func(resp http.ResponseWriter, req *http.Request, msg *stream.Message) (any, error)) Option {
	return func(o *options) {
		o.NackResponse = h
	}
}

func HttpClient(cli *http.Client) Option {
	return func(o *options) {
		o.HttpClient = cli
	}
}

// ServeMux 如果只设置 http.ServeMux，不设置 http.Server，则框架外负责启动 http.Server, 框架内部仅仅只是做handler的处理。
// 如果即设置 http.ServeMux，又设置 http.Server，则框架内即负责启动 http.Server,又负责做handler的处理。
// 如果即不设置 http.ServeMux，也不设置 http.Server，返回 ErrServerIsNil 错误
// 如果不设置 http.ServeMux，但设置了 http.Server， 框架会默认创建一个 http.ServeMux, 框架内即负责启动 http.Server,又负责做handler的处理。
func ServeMux(h *http.ServeMux) Option {
	return func(o *options) {
		o.ServeMux = h
	}
}

// HttpServer 同 ServeMux
func HttpServer(srv *http.Server) Option {
	return func(o *options) {
		o.HttpServer = srv
	}
}
