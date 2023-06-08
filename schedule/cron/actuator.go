package cron

import (
	"context"
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/netx/httpx/render"

	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
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
	_ = render.JSON(w, resp, render.PureJSON())
}

type healthChecker struct {
	scheduler *Scheduler
}

func (h *healthChecker) Check(ctx context.Context) health.Health {
	if h.scheduler.isAlive() {
		return health.Up()
	}
	return health.Down()
}

func (h *healthChecker) Name() string {
	return "/schedule/cron"
}
