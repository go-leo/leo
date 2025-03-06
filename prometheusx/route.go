package prometheusx

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Append(router *mux.Router) *mux.Router {
	handel := promhttp.Handler()
	router.NewRoute().Path("/metrics").HandlerFunc(handel.ServeHTTP)
	return router
}
