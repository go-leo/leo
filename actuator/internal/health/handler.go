package health

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-leo/gox/netx/httpx/render"

	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
)

type CheckerHandler struct {
	HealthChecker          health.Checker
	HttpHealthStatusMapper health.HttpHealthStatusMapper
}

func (h *CheckerHandler) Pattern() string {
	return "/health/check/" + h.HealthChecker.Name()
}

func (h *CheckerHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	health := h.HealthCheck(request.Context())
	st := health.Status()
	writer.WriteHeader(h.HttpHealthStatusMapper.MapStatus(st))
	resp := &Response{
		Status: ResponseStatus{
			Code:        string(st.Code()),
			Description: st.Description(),
		},
		Details: health.Details(),
	}
	_ = render.JSON(writer, resp, render.PureJSON())
}

func (h *CheckerHandler) HealthCheck(ctx context.Context) (r health.Health) {
	defer func() {
		if p := recover(); p != nil {
			r = health.DownWithError(fmt.Errorf("%v", p))
		}
	}()
	return h.HealthChecker.Check(ctx)
}

type MultiCheckerHandler struct {
	HealthCheckers         []health.Checker
	HttpHealthStatusMapper health.HttpHealthStatusMapper
	StatusAggregator       health.StatusAggregator
}

func (h *MultiCheckerHandler) Pattern() string {
	return "/health/check"
}

func (h *MultiCheckerHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	components := make(map[string]*Response)
	var statuses []health.Status
	for _, healthChecker := range h.HealthCheckers {
		healthCheck := h.HealthCheck(request.Context(), healthChecker)
		st := healthCheck.Status()
		statuses = append(statuses, st)
		components[healthChecker.Name()] = &Response{
			Status: ResponseStatus{
				Code:        string(st.Code()),
				Description: st.Description(),
			},
			Details: healthCheck.Details(),
		}
	}
	st := h.StatusAggregator.AggregateStatus(statuses...)
	writer.WriteHeader(h.HttpHealthStatusMapper.MapStatus(st))
	resp := &MultiResponse{
		Status: ResponseStatus{
			Code:        string(st.Code()),
			Description: st.Description(),
		},
		Components: components,
	}
	_ = render.JSON(writer, resp, render.PureJSON())
}

func (h *MultiCheckerHandler) HealthCheck(ctx context.Context, healthChecker health.Checker) (r health.Health) {
	defer func() {
		if p := recover(); p != nil {
			r = health.DownWithError(fmt.Errorf("%v", p))
		}
	}()
	return healthChecker.Check(ctx)
}

type ResponseStatus struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

type Response struct {
	Status  ResponseStatus `json:"status,omitempty"`
	Details map[string]any `json:"details,omitempty"`
}

type MultiResponse struct {
	Status     ResponseStatus       `json:"status,omitempty"`
	Components map[string]*Response `json:"components,omitempty"`
}
