package generator

import (
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"strconv"
)

type FunctionGenerator struct{}

func (f *FunctionGenerator) GenerateAppendRoutesFunc(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("func append", service.HttpRoutesName(), "(router *", internal.MuxPackage.Ident("Router"), ") *", internal.MuxPackage.Ident("Router"), "{")
	for _, endpoint := range service.Endpoints {
		httpRule := endpoint.HttpRule()
		// 调整路径，来适应 github.com/gorilla/mux 路由规则
		path, _, _, _ := httpRule.RegularizePath(httpRule.Path())
		g.P("router.NewRoute().Name(", strconv.Quote(endpoint.FullName()), ").Methods(", strconv.Quote(httpRule.Method()), ").Path(", strconv.Quote(path), ")")
	}
	g.P("return router")
	g.P("}")
	return nil
}

func (f *FunctionGenerator) GenerateAppendServerFunc(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("func Append", service.HttpServerRoutesName(), "(router *", internal.MuxPackage.Ident("Router"), ", svc ", service.ServiceName(), ", middlewares ...", internal.EndpointPackage.Ident("Middleware"), ") ", "*", internal.MuxPackage.Ident("Router"), " {")
	g.P("transports := new", service.HttpServerTransportsName(), "(svc, middlewares...)")
	g.P("router = append", service.HttpRoutesName(), "(router)")
	for _, endpoint := range service.Endpoints {
		g.P("router.Get(", strconv.Quote(endpoint.FullName()), ").Handler(transports.", endpoint.Name(), "())")
	}
	g.P("return router")
	g.P("}")
	g.P()
	return nil
}

func (f *FunctionGenerator) GenerateNewClientFunc(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("func New", service.HttpClientName(), "(target string, opts ...", internal.HttpxTransportxPackage.Ident("ClientOption"), ") ", service.ServiceName(), " {")
	g.P("options := ", internal.HttpxTransportxPackage.Ident("NewClientOptions"), "(opts...)")
	g.P("transports := new", service.HttpClientTransportsName(), "(options.Scheme(), options.ClientTransportOptions(), options.Middlewares())")
	g.P("endpoints := new", service.ClientEndpointsName(), "(target, transports, options.Builder(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())")
	g.P("return new", service.ClientServiceName(), "(endpoints, ", internal.HttpxTransportxPackage.Ident("HttpClient"), ")")
	g.P("}")
	g.P()
	return nil
}
