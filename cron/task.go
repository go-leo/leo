package cron

import (
	"context"
	"errors"
	"sync"

	"github.com/robfig/cron/v3"

	"github.com/go-leo/leo/v2/runner"
)

var _ runner.Runnable = new(Task)

type Task struct {
	// jobs 任务清单
	jobs      []*Job
	opts      *options
	exec      *cron.Cron
	startOnce sync.Once
	stopOnce  sync.Once
}

func New(jobs []*Job, opts ...Option) *Task {
	o := new(options)
	o.apply(opts...)
	o.init()
	var cronOpts []cron.Option
	if o.Location != nil {
		cronOpts = append(cronOpts, cron.WithLocation(o.Location))
	}
	if o.Seconds {
		cronOpts = append(cronOpts, cron.WithSeconds())
	}
	if o.Parser != nil {
		cronOpts = append(cronOpts, cron.WithParser(o.Parser))
	}
	if len(o.Middlewares) > 0 {
		cronOpts = append(cronOpts, cron.WithChain(o.Middlewares...))
	}
	cronOpts = append(cronOpts, cron.WithLogger(NewLogger(o.Logger)))
	exec := cron.New(cronOpts...)
	return &Task{opts: o, exec: exec, jobs: jobs}
}

func (task *Task) Start(_ context.Context) error {
	err := errors.New("cron already started")
	task.startOnce.Do(func() {
		err = nil
		// 遍历Jobs，并将其添加到cron实例中
		for _, job := range task.jobs {
			id, e := task.exec.AddFunc(job.Spec(), job.cmd)
			if e != nil {
				err = e
				return
			}
			job.entryID = id
			job.exec = task.exec
		}
		task.exec.Run()
	})
	return err
}

func (task *Task) Stop(_ context.Context) error {
	err := errors.New("cron already stopped")
	task.stopOnce.Do(func() {
		err = nil
		// 停止cron
		task.exec.Stop()
	})
	return err
}

func (task *Task) String() string {
	return "cron task"
}

func (task *Task) Jobs() []*Job {
	return task.jobs
}
