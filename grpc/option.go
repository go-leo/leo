package grpc

import (
	"crypto/tls"

	"google.golang.org/grpc"

	"github.com/go-leo/leo/v2/registry"
)

type options struct {
	unaryInterceptors []grpc.UnaryServerInterceptor
	serverOptions     []grpc.ServerOption
	tlsConf           *tls.Config
	Registrar         registry.Registrar
}

type Option func(o *options)

func (o *options) apply(opts []Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {}

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

func Registrar(reg registry.Registrar) Option {
	return func(o *options) {
		o.Registrar = reg
	}
}
