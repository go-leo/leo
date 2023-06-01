package stream

import (
	"context"
	"net/http"

	"github.com/go-leo/gox/netx/httpx/render"

	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
)

type actuatorHandler struct {
	streamer *Streamer
}

func (h *actuatorHandler) Pattern() string {
	return "/stream"
}

func (h *actuatorHandler) Handle(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{}
	var handlers []map[string]any
	for _, handler := range h.streamer.options.Handlers {
		handlerMap := map[string]any{}
		if handler.Subscriber != nil {
			handlerMap["subscriber"] = map[string]any{
				"topic": handler.Subscriber.Topic(),
				"queue": handler.Subscriber.Queue(),
			}
		}
		if handler.Publisher != nil {
			handlerMap["publisher"] = map[string]any{
				"topic": handler.Publisher.Topic(),
				"queue": handler.Publisher.Queue(),
			}
		}
		handlers = append(handlers, handlerMap)
	}
	resp["handlers"] = handlers
	_ = render.JSON(w, resp, render.PureJSON())
}

type healthChecker struct {
	streamer *Streamer
}

// TODO
func (h *healthChecker) Check(ctx context.Context) health.Health {
	return health.Up()
	return health.Down()
}

func (h *healthChecker) Name() string {
	return "/stream"
}
