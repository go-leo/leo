package http

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-leo/stringx"
)

type options struct {
	TLSConf           *tls.Config
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
	ReadHeaderTimeout time.Duration
	TLSNextProto      map[string]func(*http.Server, *tls.Conn, http.Handler)
	ConnState         func(net.Conn, http.ConnState)
	ErrorLog          *log.Logger
	BaseContext       func(net.Listener) context.Context
	ConnContext       func(ctx context.Context, c net.Conn) context.Context
	HealthCheckPath   string
}

type Option func(o *options)

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if stringx.IsBlank(o.HealthCheckPath) {
		o.HealthCheckPath = "/health/check"
	}
}

func TLS(conf *tls.Config) Option {
	return func(o *options) {
		o.TLSConf = conf
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

func TLSNextProto(fs map[string]func(*http.Server, *tls.Conn, http.Handler)) Option {
	return func(o *options) {
		o.TLSNextProto = fs
	}
}

func ConnState(f func(net.Conn, http.ConnState)) Option {
	return func(o *options) {
		o.ConnState = f
	}
}

func ErrorLog(l *log.Logger) Option {
	return func(o *options) {
		o.ErrorLog = l
	}
}

func BaseContext(f func(net.Listener) context.Context) Option {
	return func(o *options) {
		o.BaseContext = f
	}
}

func ConnContext(f func(ctx context.Context, c net.Conn) context.Context) Option {
	return func(o *options) {
		o.ConnContext = f
	}
}

func HealthCheckPath(path string) Option {
	return func(o *options) {
		o.HealthCheckPath = path
	}
}
