package stream

import (
	"context"
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/netx/httpx/render"

	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
)

type actuatorHandler struct {
	streamer *Streamer
}

func (h *actuatorHandler) Pattern() string {
	return "/stream"
}

func (h *actuatorHandler) Handle(w http.ResponseWriter, r *http.Request) {
	resp := h.actuatorResponse()
	_ = render.JSON(w, resp, render.PureJSON())
}

func (h *actuatorHandler) actuatorResponse() map[string]any {
	resp := map[string]any{}
	var handlers []map[string]any
	for _, handler := range h.streamer.handlerWrappers {
		handlerMap := map[string]any{}
		if handler.subscriber != nil {
			handlerMap["subscriber"] = map[string]any{
				"topic": handler.subscriber.Topic(),
				"queue": handler.subscriber.Queue(),
			}
		}
		if handler.publisher != nil {
			handlerMap["publisher"] = map[string]any{
				"topic": handler.publisher.Topic(),
				"queue": handler.publisher.Queue(),
			}
		}
		handlers = append(handlers, handlerMap)
	}
	resp["handlers"] = handlers
	return resp
}

type healthChecker struct {
	streamer *Streamer
}

func (h *healthChecker) Check(ctx context.Context) health.Health {
	if h.streamer.isAlive() {
		return health.Up()
	}
	return health.Down()
}

func (h *healthChecker) Name() string {
	return "/stream"
}
