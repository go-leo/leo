package pprof

import (
	"net/http"
	"net/http/pprof"
)

type IndexHandler struct{}

func (h *IndexHandler) Pattern() string { return "/debug/pprof/" }

func (h *IndexHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	pprof.Index(writer, request)
}

type CmdlineHandler struct{}

func (h *CmdlineHandler) Pattern() string { return "/debug/pprof/cmdline" }

func (h *CmdlineHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	pprof.Cmdline(writer, request)
}

type ProfileHandler struct{}

func (h *ProfileHandler) Pattern() string { return "/debug/pprof/profile" }

func (h *ProfileHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	pprof.Profile(writer, request)
}

type SymbolHandler struct{}

func (h *SymbolHandler) Pattern() string { return "/debug/pprof/symbol" }

func (h *SymbolHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	pprof.Symbol(writer, request)
}

type TraceHandler struct{}

func (h *TraceHandler) Pattern() string { return "/debug/pprof/trace" }

func (h *TraceHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	pprof.Trace(writer, request)
}
