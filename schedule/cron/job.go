package cron

import (
	"context"

	"github.com/robfig/cron/v3"

	"codeup.aliyun.com/qimao/leo/leo/schedule"
)

type cronJob struct {
	Task schedule.Task
	// cron里一个Entry
	EntryID cron.EntryID
	// Middlewares 中间件
	Middlewares []schedule.Middleware
}

func (j *cronJob) Run() {
	schedule.Chain(j.Task, j.Middlewares...).Invoke(context.Background())
}
