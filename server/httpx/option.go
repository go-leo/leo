package httpx

import (
	"crypto/tls"
	"time"
)

type options struct {
	Addr                         string
	DisableGeneralOptionsHandler bool
	TLSConfig                    *tls.Config
	ReadTimeout                  time.Duration
	WriteTimeout                 time.Duration
	IdleTimeout                  time.Duration
	MaxHeaderBytes               int
	ReadHeaderTimeout            time.Duration
	ShutdownTimeout              time.Duration
}

type Option func(o *options)

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *options) init() *options {
	if o.Addr == "" {
		o.Addr = ":0"
	}
	return o
}

// Addr set server addr.
func Addr(addr string) Option {
	return func(o *options) {
		o.Addr = addr
	}
}

func TLS(conf *tls.Config) Option {
	return func(o *options) {
		o.TLSConfig = conf
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

func ReadHeaderTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ReadHeaderTimeout = timeout
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.ShutdownTimeout = timeout
	}
}
