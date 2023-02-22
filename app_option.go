package leo

import (
	"os"

	"codeup.aliyun.com/qimao/leo/leo/console"
	"codeup.aliyun.com/qimao/leo/leo/controller"
	"codeup.aliyun.com/qimao/leo/leo/pubsub"
	"codeup.aliyun.com/qimao/leo/leo/resource"
	"codeup.aliyun.com/qimao/leo/leo/rpc"
	"codeup.aliyun.com/qimao/leo/leo/runner"
	"codeup.aliyun.com/qimao/leo/leo/scheduler"
)

type options struct {
	InfoLogger      interface{ Infof(string, ...any) }
	Runners         []runner.Runner
	Consoles        []console.Console
	Controllers     []controller.Controller
	Resources       []resource.Resource
	Subscribers     []pubsub.Subscriber
	Providers       []rpc.Provider
	Schedulers      []scheduler.Scheduler
	ShutdownSignals []os.Signal
}

func (o *options) init() {}

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

func Console(consoles ...console.Console) Option {
	return func(o *options) {
		o.Consoles = append(o.Consoles, consoles...)
	}
}

func Controller(controllers ...controller.Controller) Option {
	return func(o *options) {
		o.Controllers = append(o.Controllers, controllers...)
	}
}

func Resource(resources ...resource.Resource) Option {
	return func(o *options) {
		o.Resources = append(o.Resources, resources...)
	}
}

func Subscriber(subscribers ...pubsub.Subscriber) Option {
	return func(o *options) {
		o.Subscribers = append(o.Subscribers, subscribers...)
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
