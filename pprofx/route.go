package pprofx

import (
	"github.com/gorilla/mux"
	"net/http/pprof"
)

func Append(router *mux.Router) *mux.Router {
	router.NewRoute().Path("/debug/pprof/").HandlerFunc(pprof.Index)
	router.NewRoute().Path("/debug/pprof/cmdline").HandlerFunc(pprof.Cmdline)
	router.NewRoute().Path("/debug/pprof/profile").HandlerFunc(pprof.Profile)
	router.NewRoute().Path("/debug/pprof/symbol").HandlerFunc(pprof.Symbol)
	router.NewRoute().Path("/debug/pprof/trace").HandlerFunc(pprof.Trace)
	return router
}
