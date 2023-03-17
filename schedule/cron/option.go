package cron

import (
	"time"

	"github.com/robfig/cron/v3"

	"codeup.aliyun.com/qimao/leo/leo/schedule"
)

type options struct {
	Middlewares         []schedule.Middleware
	Location            *time.Location
	Seconds             bool
	DelayIfStillRunning bool
	SkipIfStillRunning  bool
	Task                []schedule.Task
}

func (o *options) init() {
	if o.Location == nil {
		o.Location = time.Local
	}
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) cronOptions() []cron.Option {
	var opts []cron.Option
	if o.Seconds {
		opts = append(opts, cron.WithSeconds())
	}
	if o.Location != nil {
		opts = append(opts, cron.WithLocation(o.Location))
	}
	var chains []cron.JobWrapper
	if o.DelayIfStillRunning {
		chains = append(chains, cron.DelayIfStillRunning(cron.DiscardLogger))
	}
	if o.SkipIfStillRunning {
		chains = append(chains, cron.SkipIfStillRunning(cron.DiscardLogger))
	}
	if len(chains) > 0 {
		opts = append(opts, cron.WithChain(chains...))
	}
	return opts
}

type Option func(*options)

// Location 时区
func Location(location *time.Location) Option {
	return func(options *options) {
		options.Location = location
	}
}

// Seconds cron表达式支持秒位
func Seconds() Option {
	return func(o *options) {
		o.Seconds = true
	}
}

// DelayIfStillRunning 延迟后续的invoker，直到前一个invoker完成。
func DelayIfStillRunning() Option {
	return func(o *options) {
		o.DelayIfStillRunning = true
	}
}

// SkipIfStillRunning 如果之前的invoker仍在运行,跳过当前的invoker。
func SkipIfStillRunning() Option {
	return func(o *options) {
		o.SkipIfStillRunning = true
	}
}

func Middlewares(middlewares ...schedule.Middleware) Option {
	return func(o *options) {
		o.Middlewares = append(o.Middlewares, middlewares...)
	}
}

func Tasks(task ...schedule.Task) Option {
	return func(o *options) {
		o.Task = append(o.Task, task...)
	}
}
