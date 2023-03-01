package handler

import (
	"net/http"
	"net/http/pprof"
)

type PProfIndexHandler struct{}

func (h *PProfIndexHandler) Pattern() string {
	return "/actuator/debug/pprof/"
}

func (h *PProfIndexHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	pprof.Index(writer, request)
}

type PProfCmdlineHandler struct{}

func (h *PProfCmdlineHandler) Pattern() string {
	return "/actuator/debug/pprof/cmdline"
}

func (h *PProfCmdlineHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	pprof.Cmdline(writer, request)
}

type PProfProfileHandler struct{}

func (h *PProfProfileHandler) Pattern() string {
	return "/actuator/debug/pprof/profile"
}

func (h *PProfProfileHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	pprof.Profile(writer, request)
}

type PProfSymbolHandler struct{}

func (h *PProfSymbolHandler) Pattern() string {
	return "/actuator/debug/pprof/symbol"
}

func (h *PProfSymbolHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	pprof.Symbol(writer, request)
}

type PProfTraceHandler struct{}

func (h *PProfTraceHandler) Pattern() string {
	return "/actuator/debug/pprof/trace"
}

func (h *PProfTraceHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	pprof.Trace(writer, request)
}
