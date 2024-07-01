package health

import (
	"context"
	"github.com/go-leo/gox/netx/httpx/render"
	"github.com/go-leo/leo/v3/healthx"
	"github.com/go-leo/leo/v3/server/actuator"
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	actuator.RegisterHandler(SystemHandler{})
	actuator.RegisterHandler(ComponentHandler{})
}

type SystemHandler struct{}

func (SystemHandler) Pattern() string {
	return "/health"
}

func (SystemHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var statuses []healthx.Status
	components := make(map[string]string)
	checkers := healthx.GetCheckers()
	for _, checker := range checkers {
		status := Check(request.Context(), checker)
		statuses = append(statuses, status)
		components[checker.Name()] = status.Name()
	}
	aggregator := getStatusAggregator()
	status := aggregator.AggregateStatus(statuses...)
	writer.WriteHeader(getHttpStatusMapper().MapStatus(status))
	_ = render.JSON(writer, &Response{Status: status.Name(), Components: components})
}

type ComponentHandler struct{}

func (ComponentHandler) Pattern() string {
	return "/health/{component}"
}

func (ComponentHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var status healthx.Status
	checkers := healthx.GetCheckers()
	value, ok := checkers[mux.Vars(request)["component"]]
	if !ok {
		status = healthx.ServiceUnknown
	} else {
		checker, _ := value.(healthx.Checker)
		status = Check(request.Context(), checker)
	}
	writer.WriteHeader(getHttpStatusMapper().MapStatus(status))
	_ = render.JSON(writer, &Response{Status: status.Name()})
}

func Check(ctx context.Context, checker healthx.Checker) (r healthx.Status) {
	defer func() {
		if p := recover(); p != nil {
			r = healthx.NotServing
		}
	}()
	return checker.Check(ctx)
}

type Response struct {
	Status     string            `json:"status"`
	Components map[string]string `json:"components,omitempty"`
}
