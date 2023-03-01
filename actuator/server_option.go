package actuator

import (
	"crypto/tls"

	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	"codeup.aliyun.com/qimao/leo/leo/console"
	"codeup.aliyun.com/qimao/leo/leo/resource"
	"codeup.aliyun.com/qimao/leo/leo/rpc"
	"codeup.aliyun.com/qimao/leo/leo/schedule"
	"codeup.aliyun.com/qimao/leo/leo/stream"
	"codeup.aliyun.com/qimao/leo/leo/view"
)

type options struct {
	TLSConf                *tls.Config
	ConsoleCommander       console.Commander
	ViewController         view.Controller
	ResourceServer         resource.Server
	SteamRouter            stream.Router
	RPCProvider            rpc.Provider
	Scheduler              schedule.Scheduler
	HealthCheckers         []health.Checker
	HttpHealthStatusMapper health.HttpHealthStatusMapper
	Handlers               []Handler
	PProfEnabled           bool
}

func (o *options) init() {

}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func TLSConf(conf *tls.Config) Option {
	return func(o *options) {
		o.TLSConf = conf
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

func HealthCheck(checkers []health.Checker, mapper health.HttpHealthStatusMapper) Option {
	return func(o *options) {
		o.HealthCheckers = checkers
		o.HttpHealthStatusMapper = mapper
	}
}

func Handlers(handlers ...Handler) Option {
	return func(o *options) {
		o.Handlers = append(o.Handlers, handlers...)
	}
}
