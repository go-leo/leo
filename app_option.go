package leo

import (
	"context"
	"os"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

type options struct {
	Logger          log.Logger
	Runners         []Runner
	ShutdownSignals []os.Signal
}

func (o *options) init() {
	if o.Logger == nil {
		o.Logger = log.L()
	}
}

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

func RunnerFuncs(runners ...func(ctx context.Context) error) Option {
	return func(o *options) {
		for _, runner := range runners {
			o.Runners = append(o.Runners, RunnerFunc(runner))
		}
	}
}

// ShutdownSignals 关闭信号
func ShutdownSignals(signals ...os.Signal) Option {
	return func(o *options) {
		o.ShutdownSignals = signals
	}
}

// Logger 注入日志
func Logger(l log.Logger) Option {
	return func(o *options) {
		o.Logger = l
	}
}
