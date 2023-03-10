package health

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-leo/gox/netx/httpx"
)

type CheckerHandler struct {
	HealthChecker          Checker
	HttpHealthStatusMapper HttpHealthStatusMapper
}

func (h *CheckerHandler) Pattern() string {
	return "/health/check/" + h.HealthChecker.Name()
}

func (h *CheckerHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	heal := h.Health(request.Context())
	st := heal.Status()
	writer.WriteHeader(h.HttpHealthStatusMapper.MapStatus(st))
	resp := &Response{
		Status: ResponseStatus{
			Code:        string(st.Code()),
			Description: st.Description(),
		},
		Details: heal.Details(),
	}
	_ = httpx.WriteJSON(writer, resp)
}

func (h *CheckerHandler) Health(ctx context.Context) (r Health) {
	defer func() {
		if p := recover(); p != nil {
			r = DownHealthWithError(fmt.Errorf("%v", p))
		}
	}()
	return h.HealthChecker.Check(ctx)
}

type ResponseStatus struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

type Response struct {
	Status  ResponseStatus `json:"status,omitempty"`
	Details map[string]any `json:"details,omitempty"`
}
