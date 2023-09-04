package recovery

import (
	"github.com/robfig/cron/v3"

	"github.com/hmldd/leo/log"
	crontask "github.com/hmldd/leo/runner/task/cron"
)

func CronMiddleware(l log.Logger) cron.JobWrapper {
	return cron.Recover(crontask.NewLogger(l))
}
