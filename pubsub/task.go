package pubsub

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/go-leo/leo/v2/runner"
)

var _ runner.Runnable = new(Task)

type Task struct {
	o         *options
	jobs      []*Job
	startOnce sync.Once
	stopOnce  sync.Once
	router    *message.Router
}

func New(jobs []*Job, opts ...Option) *Task {
	o := new(options)
	o.apply(opts...)
	o.init()
	return &Task{jobs: jobs, o: o}
}

func (task *Task) Start(ctx context.Context) error {
	err := errors.New("pubsub already started")
	task.startOnce.Do(func() {
		err = nil
		config := message.RouterConfig{
			CloseTimeout: task.o.CloseTimeout,
		}
		task.router, err = message.NewRouter(config, &logger{l: task.o.Logger})
		if err != nil {
			return
		}
		for _, mdw := range task.o.Middlewares {
			task.router.AddMiddleware(mdw)
		}
		for _, p := range task.o.Plugins {
			task.router.AddPlugin(p)
		}
		for _, job := range task.jobs {
			if job.subscriber == nil {
				err = fmt.Errorf("%s subscriber is nil", job.Name())
				return
			}
			if job.publisher == nil {
				job.handler = task.router.AddNoPublisherHandler(job.Name(), job.SubscribeTopic(), job.Subscriber(), job.noPublishHandlerFunc)
				for _, middleware := range job.middlewares {
					job.handler.AddMiddleware(middleware)
				}
				continue
			}
			job.handler = task.router.AddHandler(job.Name(), job.SubscribeTopic(), job.Subscriber(), job.PublishTopic(), job.Publisher(), job.handlerFunc)
			for _, middleware := range job.middlewares {
				job.handler.AddMiddleware(middleware)
			}
		}
		for _, decorator := range task.o.SubscriberDecorators {
			task.router.AddSubscriberDecorators(decorator)
		}
		for _, decorator := range task.o.PublisherDecorators {
			task.router.AddPublisherDecorators(decorator)
		}
		err = task.router.Run(ctx)
	})
	return err
}

func (task *Task) Stop(ctx context.Context) error {
	err := errors.New("pubsub already stopped")
	task.stopOnce.Do(func() {
		err = nil
		// 停止pubsub
		err = task.router.Close()
	})
	return err
}

func (task *Task) String() string {
	return "pubsub"
}

func (task *Task) Jobs() []*Job {
	return task.jobs
}
