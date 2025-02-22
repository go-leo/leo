package http

import (
	"github.com/go-leo/leo/v3/cmd/gen/internal"
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
	g.P("endpoints := &", service.Unexported(service.ServerEndpointsName()), "{")
	g.P("svc:         svc,")
	g.P("middlewares: middlewares,")
	g.P("}")
	g.P("transports := &", service.Unexported(service.HttpServerTransportsName()), "{")
	g.P("endpoints:       endpoints,")
	g.P("requestDecoder:  ", service.Unexported(service.HttpServerRequestDecoderName()), "{},")
	g.P("responseEncoder: ", service.Unexported(service.HttpServerResponseEncoderName()), "{},")
	g.P("}")
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

	g.P("requestEncoder := &", service.Unexported(service.HttpClientRequestEncoderName()), "{")
	g.P("router: append", service.HttpRoutesName(), "(", internal.MuxPackage.Ident("NewRouter"), "()),")
	g.P("scheme: options.Scheme(),")
	g.P("}")
	g.P("responseDecoder := &", service.Unexported(service.HttpClientResponseDecoderName()), "{}")
	g.P("transports :=  &", service.Unexported(service.HttpClientTransportsName()), "{")
	g.P("clientOptions: options.ClientTransportOptions(),")
	g.P("middlewares:   options.Middlewares(),")
	g.P("requestEncoder:  requestEncoder,")
	g.P("responseDecoder: responseDecoder,")
	g.P("}")
	g.P("factories := &", service.Unexported(service.FactoriesName()), "{")
	g.P("transports: transports,")
	g.P("}")
	g.P("endpointer := &", service.Unexported(service.EndpointersName()), "{")
	g.P("target:    target,")
	g.P("builder:   options.Builder(),")
	g.P("factories: factories,")
	g.P("logger:    options.Logger(),")
	g.P("options:   options.EndpointerOptions(),")
	g.P("}")
	g.P("balancers := &", service.Unexported(service.BalancersName()), "{")
	g.P("factory:    options.BalancerFactory(),")
	g.P("endpointer: endpointer,")
	g.P("sayHello:   ", internal.LazyLoadxPackage.Ident("Group"), "[", internal.LbPackage.Ident("Balancer"), "]{},")
	g.P("}")
	g.P("endpoints := &", service.Unexported(service.ClientEndpointsName()), "{")
	g.P("balancers: balancers,")
	g.P("}")

	g.P("return &", service.Unexported(service.ClientServiceName()), "{")
	g.P("endpoints:     endpoints,")
	g.P("transportName: ", internal.HttpxTransportxPackage.Ident("HttpClient"), ",")
	g.P("}")
	g.P("}")
	g.P()
	return nil
}
