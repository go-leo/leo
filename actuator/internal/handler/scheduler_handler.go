package handler

import (
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/schedule"
)

type SchedulerHandler struct {
	Scheduler schedule.Scheduler
}

func (h *SchedulerHandler) Pattern() string {
	return "/actuator/schedule/scheduler"
}

func (h *SchedulerHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	// TODO implement me
	panic("implement me")
}
