package client

import (
	"net/http"

	"github.com/hmldd/leo/registry"
	"github.com/hmldd/leo/runner/net/http/internal/codec"
)

type CodecType uint8

const (
	JSON CodecType = iota
	Protobuf
)

type Scheme string

const (
	HTTP  Scheme = registry.TransportHTTP
	HTTPS Scheme = registry.TransportHTTPS
)

type options struct {
	Middlewares []Interceptor
	Codec       codec.Codec
	HttpClient  *http.Client
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.Middlewares == nil {
		o.Middlewares = make([]Interceptor, 0)
	}
	if o.Codec == nil {
		o.Codec = codec.JSONCodec
	}
	if o.HttpClient == nil {
		o.HttpClient = http.DefaultClient
	}
}

type Option func(o *options)

func Middleware(middlewares ...Interceptor) Option {
	return func(o *options) {
		o.Middlewares = append(o.Middlewares, middlewares...)
	}
}

func Codec(c CodecType) Option {
	return func(o *options) {
		switch c {
		case JSON:
			o.Codec = codec.JSONCodec
		case Protobuf:
			o.Codec = codec.ProtobufCodec
		}
	}
}

func HttpClient(cli *http.Client) Option {
	return func(o *options) {
		o.HttpClient = cli
	}
}
