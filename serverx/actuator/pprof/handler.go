package pprof

import (
	"github.com/go-leo/leo/v3/server/actuator"
	"net/http"
	"net/http/pprof"
)

func init() {
	actuator.RegisterHandler(IndexHandler{})
	actuator.RegisterHandler(CmdlineHandler{})
	actuator.RegisterHandler(ProfileHandler{})
	actuator.RegisterHandler(SymbolHandler{})
	actuator.RegisterHandler(TraceHandler{})
}

type IndexHandler struct{}

func (IndexHandler) Pattern() string { return "/debug/pprof/" }

func (IndexHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	pprof.Index(writer, request)
}

type CmdlineHandler struct{}

func (CmdlineHandler) Pattern() string { return "/debug/pprof/cmdline" }

func (CmdlineHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	pprof.Cmdline(writer, request)
}

type ProfileHandler struct{}

func (ProfileHandler) Pattern() string { return "/debug/pprof/profile" }

func (ProfileHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	pprof.Profile(writer, request)
}

type SymbolHandler struct{}

func (SymbolHandler) Pattern() string { return "/debug/pprof/symbol" }

func (SymbolHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	pprof.Symbol(writer, request)
}

type TraceHandler struct{}

func (TraceHandler) Pattern() string { return "/debug/pprof/trace" }

func (TraceHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	pprof.Trace(writer, request)
}
