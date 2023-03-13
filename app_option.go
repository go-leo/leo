package leo

import (
	"os"

	"codeup.aliyun.com/qimao/leo/leo/actuator"
	"codeup.aliyun.com/qimao/leo/leo/log"
)

type options struct {
	Logger          log.Logger
	Runners         []Runner
	ActuatorServer  *actuator.Server
	ShutdownSignals []os.Signal
}

func (o *options) init() {
	if o.Logger == nil {
		o.Logger = log.Discard{}
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

func ActuatorServer(server *actuator.Server) Option {
	return func(o *options) {
		o.ActuatorServer = server
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
