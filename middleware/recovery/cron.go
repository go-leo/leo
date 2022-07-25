package recovery

import (
	"github.com/robfig/cron/v3"

	"github.com/go-leo/leo/log"
	crontask "github.com/go-leo/leo/runner/task/cron"
)

func CronMiddleware(l log.Logger) cron.JobWrapper {
	return cron.Recover(crontask.NewLogger(l))
}
