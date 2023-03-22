package lgrpc

import (
	"time"

	"google.golang.org/grpc"

	"codeup.aliyun.com/qimao/leo/leo/registry"
)

type options struct {
	ID                 string
	Name               string
	MetaData           map[string]string
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	ServerOptions      []grpc.ServerOption
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

func ID(id string) Option {
	return func(o *options) {
		o.ID = id
	}
}

func Name(name string) Option {
	return func(o *options) {
		o.Name = name
	}
}

func MetaData(m map[string]string) Option {
	return func(o *options) {
		o.MetaData = m
	}
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
