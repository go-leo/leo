package cron

import (
	"github.com/robfig/cron/v3"

	"github.com/go-leo/leo/v2/log"
	crontask "github.com/go-leo/leo/v2/runner/task/cron"
)

func SkipIfStillRunning(l log.Logger) cron.JobWrapper {
	return cron.SkipIfStillRunning(crontask.NewLogger(l))
}

func DelayIfStillRunning(l log.Logger) cron.JobWrapper {
	return cron.DelayIfStillRunning(crontask.NewLogger(l))
}
