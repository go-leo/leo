package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-leo/gox/netx/httpx"

	"codeup.aliyun.com/qimao/leo/leo/actuator/health"
)

type HealthCheckerHandler struct {
	HealthChecker          health.Checker
	HttpHealthStatusMapper health.HttpHealthStatusMapper
}

func (h *HealthCheckerHandler) Pattern() string {
	return "/actuator/health/check/" + h.HealthChecker.Name()
}

func (h *HealthCheckerHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	heal := h.Health(request.Context())
	st := heal.Status()
	writer.WriteHeader(h.HttpHealthStatusMapper.MapStatus(st))
	resp := &HealthResponse{
		Status: HealthStatus{
			Code:        string(st.Code()),
			Description: st.Description(),
		},
		Details: heal.Details(),
	}
	_ = httpx.WriteJSON(writer, resp)
}

func (h *HealthCheckerHandler) Health(ctx context.Context) (r health.Health) {
	defer func() {
		if p := recover(); p != nil {
			r = health.DownHealthWithError(fmt.Errorf("%v", p))
		}
	}()
	return h.HealthChecker.Check(ctx)
}

type HealthStatus struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

type HealthResponse struct {
	Status  HealthStatus   `json:"status,omitempty"`
	Details map[string]any `json:"details,omitempty"`
}
