package cron

import (
	"github.com/robfig/cron/v3"

	leocron "github.com/go-leo/leo/v2/cron"
	"github.com/go-leo/leo/v2/log"
)

func SkipIfStillRunning(l log.Logger) cron.JobWrapper {
	return cron.SkipIfStillRunning(leocron.NewLogger(l))
}

func DelayIfStillRunning(l log.Logger) cron.JobWrapper {
	return cron.DelayIfStillRunning(leocron.NewLogger(l))
}

func CronMiddleware(l log.Logger) cron.JobWrapper {
	return cron.Recover(leocron.NewLogger(l))
}
