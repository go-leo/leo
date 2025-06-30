package pprofx

import (
	"net/http/pprof"

	"github.com/gorilla/mux"
)

func Append(router *mux.Router) *mux.Router {
	router.PathPrefix("/debug/pprof/").HandlerFunc(pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	return router
}
