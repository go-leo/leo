package leo

import (
	"os"

	"codeup.aliyun.com/qimao/leo/leo/actuator"
	"codeup.aliyun.com/qimao/leo/leo/console"
	"codeup.aliyun.com/qimao/leo/leo/log"
	"codeup.aliyun.com/qimao/leo/leo/resource"
	"codeup.aliyun.com/qimao/leo/leo/rpc"
	"codeup.aliyun.com/qimao/leo/leo/runner"
	"codeup.aliyun.com/qimao/leo/leo/schedule"
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"codeup.aliyun.com/qimao/leo/leo/view"
)

type options struct {
	InfoLogger       interface{ Infof(string, ...any) }
	Runners          []runner.Runner
	ConsoleCommander console.Commander
	ViewController   view.Controller
	ResourceServer   resource.Server
	SteamRouter      stream.Router
	RPCProvider      rpc.Provider
	Scheduler        schedule.Scheduler
	ActuatorServer   *actuator.Server
	ShutdownSignals  []os.Signal
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

func ConsoleCommander(commander console.Commander) Option {
	return func(o *options) {
		o.ConsoleCommander = commander
	}
}

func ViewController(controller view.Controller) Option {
	return func(o *options) {
		o.ViewController = controller
	}
}

func ResourceServer(resource resource.Server) Option {
	return func(o *options) {
		o.ResourceServer = resource
	}
}

func StreamRouter(router stream.Router) Option {
	return func(o *options) {
		o.SteamRouter = router
	}
}

func RPCProvider(provider rpc.Provider) Option {
	return func(o *options) {
		o.RPCProvider = provider
	}
}

func Scheduler(scheduler schedule.Scheduler) Option {
	return func(o *options) {
		o.Scheduler = scheduler
	}
}

func ActuatorServer(server *actuator.Server) Option {
	return func(o *options) {
		o.ActuatorServer = server
	}
}

// ShutdownSignal 关闭信号
func ShutdownSignal(signals []os.Signal) Option {
	return func(o *options) {
		o.ShutdownSignals = signals
	}
}

// Logger 注入日志
func Logger(l log.Logger) Option {
	return func(o *options) {
		o.InfoLogger = l
	}
}
