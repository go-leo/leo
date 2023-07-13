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
	var channels []map[string]any
	for _, channel := range h.streamer.channels {
		handlerMap := map[string]any{}
		if channel.Subscriber() != nil {
			handlerMap["subscriber"] = map[string]any{
				"topic": channel.Subscriber().Topic(),
				"queue": channel.Subscriber().Queue(),
			}
		}
		if channel.Publisher() != nil {
			handlerMap["publisher"] = map[string]any{
				"topic": channel.Publisher().Topic(),
				"queue": channel.Publisher().Queue(),
			}
		}
		channels = append(channels, handlerMap)
	}
	resp["channels"] = channels
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
