package actuator

import (
	"crypto/tls"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	"codeup.aliyun.com/qimao/leo/leo/actuator/pprof"
)

type options struct {
	TLSConf                *tls.Config
	HealthCheckers         []health.Checker
	HttpHealthStatusMapper health.HttpHealthStatusMapper
	PProfDisabled          bool
	PathPrefix             string
	Handlers               []Handler

	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	CloseTimeout   time.Duration
	MaxHeaderBytes int
}

func (o *options) init() {
	if o.HttpHealthStatusMapper == nil {
		o.HttpHealthStatusMapper = health.DefaultHttpHealthStatusMapper()
	}
	if !o.PProfDisabled {
		o.Handlers = append(o.Handlers, &pprof.IndexHandler{})
		o.Handlers = append(o.Handlers, &pprof.CmdlineHandler{})
		o.Handlers = append(o.Handlers, &pprof.ProfileHandler{})
		o.Handlers = append(o.Handlers, &pprof.SymbolHandler{})
		o.Handlers = append(o.Handlers, &pprof.TraceHandler{})
	}
	for _, checker := range o.HealthCheckers {
		handler := &health.CheckerHandler{
			HealthChecker:          checker,
			HttpHealthStatusMapper: o.HttpHealthStatusMapper,
		}
		o.Handlers = append(o.Handlers, handler)
	}
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func TLSConfig(conf *tls.Config) Option {
	return func(o *options) {
		o.TLSConf = conf
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

func CloseTimeout(d time.Duration) Option {
	return func(o *options) {
		o.CloseTimeout = d
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

func HttpHealthStatusMapper(mapper health.HttpHealthStatusMapper) Option {
	return func(o *options) {
		o.HttpHealthStatusMapper = mapper
	}
}

func HealthCheckers(checkers ...health.Checker) Option {
	return func(o *options) {
		o.HealthCheckers = append(o.HealthCheckers, checkers...)
	}
}

func Handlers(handlers ...Handler) Option {
	return func(o *options) {
		o.Handlers = append(o.Handlers, handlers...)
	}
}

func DisablePProf() Option {
	return func(o *options) {
		o.PProfDisabled = true
	}
}
