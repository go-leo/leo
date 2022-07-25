package cron

import (
	"github.com/robfig/cron/v3"

	"github.com/go-leo/leo/log"
	crontask "github.com/go-leo/leo/runner/task/cron"
)

func SkipIfStillRunning(l log.Logger) cron.JobWrapper {
	return cron.SkipIfStillRunning(crontask.NewLogger(l))
}

func DelayIfStillRunning(l log.Logger) cron.JobWrapper {
	return cron.DelayIfStillRunning(crontask.NewLogger(l))
}
