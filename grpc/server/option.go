package server

import (
	"crypto/tls"

	"google.golang.org/grpc"
)

type options struct {
	unaryInterceptors []grpc.UnaryServerInterceptor
	serverOptions     []grpc.ServerOption
	tlsConf           *tls.Config
}

type Option func(o *options)

func (o *options) apply(opts []Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {

}

func UnaryInterceptors(unaryInterceptors ...grpc.UnaryServerInterceptor) Option {
	return func(o *options) {
		o.unaryInterceptors = unaryInterceptors
	}
}

func ServerOptions(serverOptions ...grpc.ServerOption) Option {
	return func(o *options) {
		o.serverOptions = serverOptions
	}
}

func TLS(conf *tls.Config) Option {
	return func(o *options) {
		o.tlsConf = conf
	}
}
