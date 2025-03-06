package healthx

import (
	"context"
	"github.com/go-leo/gox/netx/httpx/render"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net/http"
)

func Append(router *mux.Router) *mux.Router {
	router.NewRoute().
		Path("/health").
		HandlerFunc(System)
	router.NewRoute().
		Path("/health/{name}").
		HandlerFunc(Component)
	return router
}

func System(writer http.ResponseWriter, request *http.Request) {
	var statuses []grpc_health_v1.HealthCheckResponse_ServingStatus
	components := make(map[string]string)
	checkers := GetCheckers()
	for _, checker := range checkers {
		status := check(request.Context(), checker)
		statuses = append(statuses, status)
		components[checker.Name()] = status.String()
	}
	status := GetAggregator().Aggregate(statuses...)
	writer.WriteHeader(GetMapper().Map(status))
	_ = render.JSON(writer, &Response{Status: status.String(), Components: components})
}

func Component(writer http.ResponseWriter, request *http.Request) {
	var status grpc_health_v1.HealthCheckResponse_ServingStatus
	checker, ok := GetChecker(mux.Vars(request)["name"])
	if ok {
		status = check(request.Context(), checker)
	} else {
		status = grpc_health_v1.HealthCheckResponse_UNKNOWN
	}
	writer.WriteHeader(GetMapper().Map(status))
	_ = render.JSON(writer, &Response{Status: status.String()})
}

func check(ctx context.Context, checker Checker) (r grpc_health_v1.HealthCheckResponse_ServingStatus) {
	if checker == nil {
		return grpc_health_v1.HealthCheckResponse_UNKNOWN
	}
	defer func() {
		if p := recover(); p != nil {
			r = grpc_health_v1.HealthCheckResponse_UNKNOWN
		}
	}()
	checkResp, err := checker.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
	if err != nil {
		return grpc_health_v1.HealthCheckResponse_UNKNOWN
	}
	return checkResp.GetStatus()
}

type Response struct {
	Status     string            `json:"status"`
	Components map[string]string `json:"components,omitempty"`
}
