package cron

import (
	"github.com/robfig/cron/v3"

	"github.com/hmldd/leo/log"
	crontask "github.com/hmldd/leo/runner/task/cron"
)

func SkipIfStillRunning(l log.Logger) cron.JobWrapper {
	return cron.SkipIfStillRunning(crontask.NewLogger(l))
}

func DelayIfStillRunning(l log.Logger) cron.JobWrapper {
	return cron.DelayIfStillRunning(crontask.NewLogger(l))
}
