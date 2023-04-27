package actuator

import (
	"codeup.aliyun.com/qimao/leo/leo/actuator/internal/metric"
	"crypto/tls"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	internalhealth "codeup.aliyun.com/qimao/leo/leo/actuator/internal/health"
	internallog "codeup.aliyun.com/qimao/leo/leo/actuator/internal/log"
	"codeup.aliyun.com/qimao/leo/leo/actuator/internal/pprof"
	"codeup.aliyun.com/qimao/leo/leo/log"
)

type options struct {
	TLSConf *tls.Config

	HealthCheckers         []health.Checker
	HttpHealthStatusMapper health.HttpHealthStatusMapper
	StatusAggregator       health.StatusAggregator

	PProfDisabled  bool
	MetricDisabled bool
	PathPrefix     string
	Handlers       []Handler

	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	MaxHeaderBytes int

	Logger log.Logger

	ShutdownTimeout time.Duration
}

func (o *options) init() {
	if len(o.HealthCheckers) > 0 {
		if o.HttpHealthStatusMapper == nil {
			o.HttpHealthStatusMapper = health.DefaultHttpHealthStatusMapper()
		}
		if o.StatusAggregator == nil {
			o.StatusAggregator = health.DefaultStatusAggregator()
		}
		handler := &internalhealth.MultiCheckerHandler{
			HealthCheckers:         o.HealthCheckers,
			HttpHealthStatusMapper: o.HttpHealthStatusMapper,
			StatusAggregator:       o.StatusAggregator,
		}
		o.Handlers = append(o.Handlers, handler)
		for _, checker := range o.HealthCheckers {
			handlers := &internalhealth.CheckerHandler{
				HealthChecker:          checker,
				HttpHealthStatusMapper: o.HttpHealthStatusMapper,
			}
			o.Handlers = append(o.Handlers, handlers)
		}
	}

	if !o.PProfDisabled {
		o.Handlers = append(o.Handlers, &pprof.IndexHandler{})
		o.Handlers = append(o.Handlers, &pprof.CmdlineHandler{})
		o.Handlers = append(o.Handlers, &pprof.ProfileHandler{})
		o.Handlers = append(o.Handlers, &pprof.SymbolHandler{})
		o.Handlers = append(o.Handlers, &pprof.TraceHandler{})
	}

	if !o.MetricDisabled {
		o.Handlers = append(o.Handlers, &metric.MetricHandler{})
	}

	if o.Logger != nil {
		o.Handlers = append(o.Handlers, &internallog.Handler{Logger: o.Logger})
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

func HealthCheckers(checkers ...health.Checker) Option {
	return func(o *options) {
		o.HealthCheckers = append(o.HealthCheckers, checkers...)
	}
}

func HttpHealthStatusMapper(mapper health.HttpHealthStatusMapper) Option {
	return func(o *options) {
		o.HttpHealthStatusMapper = mapper
	}
}

func HealthStatusAggregator(aggregator health.StatusAggregator) Option {
	return func(o *options) {
		o.StatusAggregator = aggregator
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
func DisableMetric() Option {
	return func(o *options) {
		o.MetricDisabled = true
	}
}

func Logger(logger log.Logger) Option {
	return func(o *options) {
		o.Logger = logger
	}
}

func ShutdownTimeout(d time.Duration) Option {
	return func(o *options) {
		o.ShutdownTimeout = d
	}
}
