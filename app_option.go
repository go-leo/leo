package leo

import (
	"os"
)

type options struct {
	Runners         []Runner
	ShutdownSignals []os.Signal
}

func (o *options) init() {}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func Runners(runners ...Runner) Option {
	return func(o *options) {
		o.Runners = append(o.Runners, runners...)
	}
}

// ShutdownSignals 关闭信号
func ShutdownSignals(signals ...os.Signal) Option {
	return func(o *options) {
		o.ShutdownSignals = signals
	}
}
