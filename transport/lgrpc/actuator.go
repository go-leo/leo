package lgrpc

import (
	"context"
	"net/http"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
	"codeup.aliyun.com/qimao/leo/leo/internal/netx/httpx"
)

type actuatorHandler struct {
	server *Server
}

func (h *actuatorHandler) Pattern() string {
	return "/transport/grpc"
}

func (h *actuatorHandler) Handle(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{
		"vendor": "google.golang.org/grpc",
	}

	services := make(map[string]map[string]any)
	for name, info := range h.server.gRPCSrv.GetServiceInfo() {
		serviceInfo := map[string]any{
			"metadata": info.Metadata,
		}

		methods := make([]map[string]any, 0, len(info.Methods))
		for _, method := range info.Methods {
			methods = append(methods, map[string]any{
				"name":             method.Name,
				"is_server_stream": method.IsServerStream,
				"is_client_stream": method.IsClientStream,
			})
		}
		serviceInfo["methods"] = methods

		services[name] = serviceInfo
	}
	resp["services"] = services
	_ = httpx.WriteJSON(w, resp)
}

type healthChecker struct {
	server *Server
}

func (h *healthChecker) Check(ctx context.Context) health.Health {
	checkResponse, err := h.server.healthSrv.Check(ctx, &healthpb.HealthCheckRequest{})
	if err != nil {
		return health.Unknown()
	}
	switch checkResponse.Status {
	case healthpb.HealthCheckResponse_UNKNOWN:
		return health.Unknown()
	case healthpb.HealthCheckResponse_SERVING:
		return health.Up()
	case healthpb.HealthCheckResponse_NOT_SERVING:
		return health.Down()
	case healthpb.HealthCheckResponse_SERVICE_UNKNOWN:
		return health.OutOfService()
	default:
		return health.Unknown()
	}
}

func (h *healthChecker) Name() string {
	return "/transport/grpc"
}
