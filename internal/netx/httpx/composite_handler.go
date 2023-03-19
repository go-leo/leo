package httpx

import (
	"net/http"
)

type CompositeHandler struct {
	matchableHandlers []*matchableHandler
}

func (h *CompositeHandler) AddHandler(handler http.Handler, match func(request *http.Request) bool) {
	h.matchableHandlers = append(h.matchableHandlers, &matchableHandler{handler: handler, match: match})
}

func (h *CompositeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, matchableHandler := range h.matchableHandlers {
		if matchableHandler.match(request) {
			matchableHandler.handler.ServeHTTP(writer, request)
			return
		}
	}
}

type matchableHandler struct {
	handler http.Handler
	match   func(request *http.Request) bool
}

func (h *matchableHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.handler.ServeHTTP(writer, request)
}
