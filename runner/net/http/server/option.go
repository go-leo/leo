package server

import (
	"crypto/tls"
	"time"

	"github.com/gin-gonic/gin"
)

type options struct {
	TLSConf          *tls.Config
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	IdleTimeout      time.Duration
	MaxHeaderBytes   int
	GinMiddlewares   []gin.HandlerFunc
	Routes           []Route
	RichRoutes       []RichRoute
	NoRouteHandlers  []gin.HandlerFunc
	NoMethodHandlers []gin.HandlerFunc
}

type Option func(o *options)

func (o *options) apply(opts []Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {}

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

func Routes(routes ...Route) Option {
	return func(o *options) {
		o.Routes = append(o.Routes, routes...)
	}
}

func RichRoutes(routes ...RichRoute) Option {
	return func(o *options) {
		o.RichRoutes = append(o.RichRoutes, routes...)
	}
}

func NoRouteHandlers(handlers ...gin.HandlerFunc) Option {
	return func(o *options) {
		o.NoRouteHandlers = append(o.NoRouteHandlers, handlers...)
	}
}

func NoMethodHandlers(handlers ...gin.HandlerFunc) Option {
	return func(o *options) {
		o.NoMethodHandlers = append(o.NoMethodHandlers, handlers...)
	}
}
