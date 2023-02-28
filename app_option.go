package leo

import (
	"os"

	"codeup.aliyun.com/qimao/leo/leo/console"
	"codeup.aliyun.com/qimao/leo/leo/controller"
	"codeup.aliyun.com/qimao/leo/leo/log"
	"codeup.aliyun.com/qimao/leo/leo/resource"
	"codeup.aliyun.com/qimao/leo/leo/rpc"
	"codeup.aliyun.com/qimao/leo/leo/runner"
	"codeup.aliyun.com/qimao/leo/leo/scheduler"
	"codeup.aliyun.com/qimao/leo/leo/stream"
)

type options struct {
	InfoLogger      interface{ Infof(string, ...any) }
	Runners         []runner.Runner
	Commanders      []console.Commander
	Controllers     []controller.Server
	Resources       []resource.Server
	Routers         []stream.Router
	Providers       []rpc.Provider
	Schedulers      []scheduler.Scheduler
	ShutdownSignals []os.Signal
}

func (o *options) init() {
	if o.InfoLogger == nil {
		o.InfoLogger = log.Discard{}
	}
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func Runner(runners ...runner.Runner) Option {
	return func(o *options) {
		o.Runners = append(o.Runners, runners...)
	}
}

func Commander(commanders ...console.Commander) Option {
	return func(o *options) {
		o.Commanders = append(o.Commanders, commanders...)
	}
}

func Controller(controllers ...controller.Server) Option {
	return func(o *options) {
		o.Controllers = append(o.Controllers, controllers...)
	}
}

func Resource(resources ...resource.Server) Option {
	return func(o *options) {
		o.Resources = append(o.Resources, resources...)
	}
}

func Stream(routers ...stream.Router) Option {
	return func(o *options) {
		o.Routers = append(o.Routers, routers...)
	}
}

func Provider(providers ...rpc.Provider) Option {
	return func(o *options) {
		o.Providers = append(o.Providers, providers...)
	}
}

func Scheduler(schedulers ...scheduler.Scheduler) Option {
	return func(o *options) {
		o.Schedulers = append(o.Schedulers, schedulers...)
	}
}

// ShutdownSignal 关闭信号
func ShutdownSignal(signals []os.Signal) Option {
	return func(o *options) {
		o.ShutdownSignals = signals
	}
}
