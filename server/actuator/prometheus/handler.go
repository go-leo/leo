package prometheus

import (
	"github.com/go-leo/leo/v3/server/actuator"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func init() {
	actuator.RegisterHandler(MetricHandler{})
}

type MetricHandler struct{}

func (MetricHandler) Pattern() string { return "/metrics/prometheus" }

func (MetricHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handel := promhttp.Handler()
	handel.ServeHTTP(writer, request)
}
