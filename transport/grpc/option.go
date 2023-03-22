package grpc

import (
	"crypto/tls"
	"time"

	"google.golang.org/grpc"

	"codeup.aliyun.com/qimao/leo/leo/registry"
)

type options struct {
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	ServerOptions      []grpc.ServerOption
	TLSConf            *tls.Config
	Registers          []any
	Registrar          registry.Registrar
	ShutdownTimeout    time.Duration
}

type Option func(o *options)

func (o *options) apply(opts []Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {

}

func UnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(o *options) {
		o.UnaryInterceptors = append(o.UnaryInterceptors, interceptors...)
	}
}

func StreamInterceptors(interceptors ...grpc.StreamServerInterceptor) Option {
	return func(o *options) {
		o.StreamInterceptors = append(o.StreamInterceptors, interceptors...)
	}
}

func ServerOptions(serverOptions ...grpc.ServerOption) Option {
	return func(o *options) {
		o.ServerOptions = serverOptions
	}
}

func TLS(conf *tls.Config) Option {
	return func(o *options) {
		o.TLSConf = conf
	}
}

func Registers(registers ...any) Option {
	return func(o *options) {
		o.Registers = append(o.Registers, registers...)
	}
}

func Registrar(registrar registry.Registrar) Option {
	return func(o *options) {
		o.Registrar = registrar
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ShutdownTimeout = timeout
	}
}
