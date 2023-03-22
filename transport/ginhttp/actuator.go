package ginhttp

import (
	"context"
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	"codeup.aliyun.com/qimao/leo/leo/internal/netx/httpx"
)

type actuatorHandler struct {
	server *Server
}

func (h *actuatorHandler) Pattern() string {
	return "/transport/http"
}

func (h *actuatorHandler) Handle(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{
		"vendor": "github.com/gin-gonic/gin",
	}
	var routes []map[string]any
	for _, route := range h.server.engine.Routes() {
		routes = append(routes, map[string]any{
			"method":  route.Method,
			"path":    route.Path,
			"handler": route.Handler,
		})
	}
	resp["routes"] = routes
	_ = httpx.WriteJSON(w, resp)
}

type healthChecker struct {
	server *Server
}

func (h *healthChecker) Check(ctx context.Context) health.Health {
	if h.server.healthSrv.IsRunning() {
		return health.Up()
	}
	return health.Down()
}

func (h *healthChecker) Name() string {
	return "/transport/http"
}
