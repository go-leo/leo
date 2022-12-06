package server

import "github.com/gin-gonic/gin"

type Route interface {
	Methods() []string
	Path() string
	Handlers() []gin.HandlerFunc
}

func NewRoute(methods []string, path string, handlers ...gin.HandlerFunc) Route {
	return &route{methods: methods, path: path, handlers: handlers}
}

type route struct {
	methods  []string
	path     string
	handlers []gin.HandlerFunc
}

func (r *route) Methods() []string {
	return r.methods
}

func (r *route) Path() string {
	return r.path
}

func (r *route) Handlers() []gin.HandlerFunc {
	return r.handlers
}
