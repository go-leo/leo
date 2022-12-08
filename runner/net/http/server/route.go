package server

import "github.com/gin-gonic/gin"

type Route interface {
	Method() string
	Path() string
	Handler() gin.HandlerFunc
}

func NewRoute(method string, path string, handler gin.HandlerFunc) Route {
	return &route{method: method, path: path, handler: handler}
}

type RichRoute interface {
	Methods() []string
	Path() string
	Handlers() []gin.HandlerFunc
}

func NewRichRoute(methods []string, path string, handlers ...gin.HandlerFunc) RichRoute {
	return &richRoute{methods: methods, path: path, handlers: handlers}
}

type route struct {
	method  string
	path    string
	handler gin.HandlerFunc
}

func (r *route) Method() string {
	return r.method
}

func (r *route) Path() string {
	return r.path
}

func (r *route) Handler() gin.HandlerFunc {
	return r.handler
}

type richRoute struct {
	methods  []string
	path     string
	handlers []gin.HandlerFunc
}

func (r *richRoute) Methods() []string {
	return r.methods
}

func (r *richRoute) Path() string {
	return r.path
}

func (r *richRoute) Handlers() []gin.HandlerFunc {
	return r.handlers
}
