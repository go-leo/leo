package leo

import (
	"github.com/go-leo/leo/v3/runner"
	"github.com/go-leo/leo/v3/server"
	"os"
)

type options struct {
	Runners []runner.Runner
	Servers []server.Server
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

func Runner(runners ...runner.Runner) Option {
	return func(o *options) {
		o.Runners = append(o.Runners, runners...)
	}
}

func Server(servers ...server.Server) Option {
	return func(o *options) {
		o.Servers = append(o.Servers, servers...)
	}
}

// Signal
func Signal(signals ...os.Signal) Option {
	return func(o *options) {
		o.Signals = signals
	}
}
