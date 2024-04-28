package leo

import (
	"github.com/go-leo/leo/v3/runner"
	"os"
)

type options struct {
	Runners []runner.Runner
	Signals []os.Signal
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

func Runners(runners ...runner.Runner) Option {
	return func(o *options) {
		o.Runners = append(o.Runners, runners...)
	}
}

// Signals 信号
func Signals(signals ...os.Signal) Option {
	return func(o *options) {
		o.Signals = signals
	}
}
