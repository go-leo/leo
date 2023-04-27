package metric

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type MetricHandler struct{}

func (h *MetricHandler) Pattern() string { return "/metrics" }

func (h *MetricHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	handel := promhttp.Handler()
	handel.ServeHTTP(writer, request)
}
