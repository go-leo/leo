package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricHandler struct{}

func (h *MetricHandler) Pattern() string { return "/metrics" }

func (h *MetricHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	handel := promhttp.Handler()
	handel.ServeHTTP(writer, request)
}
