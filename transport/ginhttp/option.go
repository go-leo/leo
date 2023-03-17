package ginhttp

import (
	"crypto/tls"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/registry"
)

type options struct {
	ID                string
	Name              string
	MetaData          map[string]string
	TLSConf           *tls.Config
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
	ReadHeaderTimeout time.Duration
	Registrar         registry.Registrar
	ShutdownTimeout   time.Duration
}

type Option func(o *options)

func (o *options) apply(opts ...Option) {
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
