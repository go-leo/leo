package cron

import (
	"context"
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	"codeup.aliyun.com/qimao/leo/leo/internal/netx/httpx"
)

type actuatorHandler struct {
	scheduler *Scheduler
}

func (h *actuatorHandler) Pattern() string {
	return "/schedule/cron"
}

func (h *actuatorHandler) Handle(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{
		"vendor": "github.com/robfig/cron/v3",
	}
	var tasks []map[string]any
	for task, job := range h.scheduler.taskJobMap {
		tasks = append(tasks, map[string]any{
			"name":       task.Name(),
			"expression": task.Expression(),
			"entry_id":   job.EntryID,
		})
	}
	resp["tasks"] = tasks
	_ = httpx.WriteJSON(w, resp)
}

type healthChecker struct {
	scheduler *Scheduler
}

func (h *healthChecker) Check(ctx context.Context) health.Health {
	if h.scheduler.isRunning() {
		return health.UpHealth()
	}
	return health.DownHealth()
}

func (h *healthChecker) Name() string {
	return "/schedule/cron"
}
