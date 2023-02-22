package leo

import (
	"github.com/go-kit/kit/log"
	"os"
)

type options struct {
	Logger          log.Logger
	Runners         []Runner
	ShutdownSignals []os.Signal
}

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *options) init() *options {
	return o
}

type Option func(o *options)

func Runners(runners ...Runner) Option {
	return func(o *options) {
		o.Runners = append(o.Runners, runners...)
	}
}

func Logger(l log.Logger) Option {
	return func(o *options) {
		o.Logger = l
	}
}

// ShutdownSignals 关闭信号
func ShutdownSignals(signals ...os.Signal) Option {
	return func(o *options) {
		o.ShutdownSignals = signals
	}
}
