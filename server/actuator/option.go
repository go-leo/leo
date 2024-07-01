package actuator

import (
	"crypto/tls"
	"github.com/gorilla/mux"
	"time"
)

type options struct {
	DisableGeneralOptionsHandler bool
	TLSConfig                    *tls.Config
	ReadTimeout                  time.Duration
	WriteTimeout                 time.Duration
	IdleTimeout                  time.Duration
	MaxHeaderBytes               int
	ReadHeaderTimeout            time.Duration
	ShutdownTimeout              time.Duration
	PathPrefix                   string
	Handlers                     []Handler
	Middlewares                  []mux.MiddlewareFunc
}

func (o *options) init() *options {
	return o
}

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type Option func(o *options)

func TLSConfig(conf *tls.Config) Option {
	return func(o *options) {
		o.TLSConfig = conf
	}
}

func ReadTimeout(d time.Duration) Option {
	return func(o *options) {
		o.ReadTimeout = d
	}
}

func WriteTimeout(d time.Duration) Option {
	return func(o *options) {
		o.WriteTimeout = d
	}
}

func IdleTimeout(d time.Duration) Option {
	return func(o *options) {
		o.IdleTimeout = d
	}
}

func MaxHeaderBytes(n int) Option {
	return func(o *options) {
		o.MaxHeaderBytes = n
	}
}

func PathPrefix(prefix string) Option {
	return func(o *options) {
		o.PathPrefix = prefix
	}
}

func Handlers(handlers ...Handler) Option {
	return func(o *options) {
		o.Handlers = append(o.Handlers, handlers...)
	}
}

func ShutdownTimeout(d time.Duration) Option {
	return func(o *options) {
		o.ShutdownTimeout = d
	}
}
