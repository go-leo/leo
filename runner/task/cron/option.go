package cron

import (
	"time"

	"github.com/robfig/cron/v3"

	"github.com/go-leo/leo/log"
)

type options struct {
	Location    *time.Location
	Middlewares []cron.JobWrapper
	Seconds     bool
	Parser      cron.ScheduleParser
	Logger      log.Logger
}

func (o *options) init() {
	if o.Location == nil {
		o.Location = time.Local
	}
	if o.Logger == nil {
		o.Logger = log.Discard{}
	}
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(*options)

// Location 时区
func Location(location *time.Location) Option {
	return func(options *options) {
		options.Location = location
	}
}

func Seconds() Option {
	return func(o *options) {
		o.Seconds = true
	}
}

func Parser(p cron.ScheduleParser) Option {
	return func(o *options) {
		o.Parser = p
	}
}

func Middleware(mdw ...cron.JobWrapper) Option {
	return func(o *options) {
		o.Middlewares = append(o.Middlewares, mdw...)
	}
}

func Logger(logger log.Logger) Option {
	return func(o *options) {
		o.Logger = logger
	}
}
