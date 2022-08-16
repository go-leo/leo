package server

import (
	"crypto/tls"
	"time"

	"github.com/gin-gonic/gin"
)

type Router struct {
	// HTTPMethods 如果为空，则注册到所有Method
	HTTPMethods  []string
	Path         string
	HandlerFuncs []gin.HandlerFunc
}

type options struct {
	GRPCClient     any
	ServiceDesc    *ServiceDesc
	TLSConf        *tls.Config
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	MaxHeaderBytes int
	GinMiddlewares []gin.HandlerFunc
	Routers        []Router
}

type Option func(o *options)

func (o *options) apply(opts []Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {}

func GRPCClient(cli any) Option {
	return func(o *options) {
		o.GRPCClient = cli
	}
}

func ServiceDescription(serviceDesc *ServiceDesc) Option {
	return func(o *options) {
		o.ServiceDesc = serviceDesc
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ReadTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.WriteTimeout = timeout
	}
}

func IdleTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.IdleTimeout = timeout
	}
}

func MaxHeaderBytes(size int) Option {
	return func(o *options) {
		o.MaxHeaderBytes = size
	}
}

func TLS(conf *tls.Config) Option {
	return func(o *options) {
		o.TLSConf = conf
	}
}

func Middlewares(middlewares ...gin.HandlerFunc) Option {
	return func(o *options) {
		o.GinMiddlewares = append(o.GinMiddlewares, middlewares...)
	}
}

func Routers(routers ...Router) Option {
	return func(o *options) {
		o.Routers = routers
	}
}
