package http

import (
	"github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"strconv"
)

type FunctionGenerator struct {
	service *internal.Service
	g       *protogen.GeneratedFile
}

func (f *FunctionGenerator) GenerateAppendRoutesFunc() {
	f.g.P("func append", f.service.HttpRoutesName(), "(router *", internal.MuxPackage.Ident("Router"), ") *", internal.MuxPackage.Ident("Router"), "{")
	for _, endpoint := range f.service.Endpoints {
		f.g.P("router.NewRoute().")
		f.g.P("Name(", strconv.Quote(endpoint.FullName()), ").")
		f.g.P("Methods(", endpoint.HttpMethod(), ").")
		f.g.P("Path(", strconv.Quote(endpoint.RoutePath()), ")")
	}
	f.g.P("return router")
	f.g.P("}")
}

func (f *FunctionGenerator) GenerateAppendServerFunc() {
	f.g.P("func Append", f.service.HttpServerRoutesName(), "(router *", internal.MuxPackage.Ident("Router"), ", svc ", f.service.ServiceName(), ", opts ...", internal.ServerOption, ") ", "*", internal.MuxPackage.Ident("Router"), " {")
	f.g.P("options := ", internal.NewServerOptions, "(opts...)")
	f.g.P("endpoints := &", f.service.Unexported(f.service.ServerEndpointsName()), "{")
	f.g.P("svc:         svc,")
	f.g.P("middlewares: options.Middlewares(),")
	f.g.P("}")
	f.g.P("requestDecoder := ", f.service.Unexported(f.service.HttpServerRequestDecoderName()), "{")
	f.g.P("unmarshalOptions: options.UnmarshalOptions(),")
	f.g.P("}")
	f.g.P("responseEncoder := ", f.service.Unexported(f.service.HttpServerResponseEncoderName()), "{")
	f.g.P("marshalOptions:      options.MarshalOptions(),")
	f.g.P("}")
	f.g.P("transports := &", f.service.Unexported(f.service.HttpServerTransportsName()), "{")
	f.g.P("endpoints:       endpoints,")
	f.g.P("requestDecoder:  requestDecoder,")
	f.g.P("responseEncoder: responseEncoder,")
	f.g.P("}")
	f.g.P("router = append", f.service.HttpRoutesName(), "(router)")
	for _, endpoint := range f.service.Endpoints {
		f.g.P("router.Get(", strconv.Quote(endpoint.FullName()), ").Handler(transports.", endpoint.Name(), "())")
	}
	f.g.P("return router")
	f.g.P("}")
	f.g.P()
}

func (f *FunctionGenerator) GenerateNewClientFunc() {
	f.g.P("func New", f.service.HttpClientName(), "(target string, opts ...", internal.HttpxTransportxPackage.Ident("ClientOption"), ") ", f.service.ServiceName(), " {")
	f.g.P("options := ", internal.HttpxTransportxPackage.Ident("NewClientOptions"), "(opts...)")

	f.g.P("requestEncoder := &", f.service.Unexported(f.service.HttpClientRequestEncoderName()), "{")
	f.g.P("marshalOptions: options.MarshalOptions(),")
	f.g.P("router: append", f.service.HttpRoutesName(), "(", internal.MuxPackage.Ident("NewRouter"), "()),")
	f.g.P("scheme: options.Scheme(),")
	f.g.P("}")
	f.g.P("responseDecoder := &", f.service.Unexported(f.service.HttpClientResponseDecoderName()), "{")
	f.g.P("unmarshalOptions: options.UnmarshalOptions(),")
	f.g.P("}")
	f.g.P("transports :=  &", f.service.Unexported(f.service.HttpClientTransportsName()), "{")
	f.g.P("clientOptions: options.ClientTransportOptions(),")
	f.g.P("middlewares:   options.Middlewares(),")
	f.g.P("requestEncoder:  requestEncoder,")
	f.g.P("responseDecoder: responseDecoder,")
	f.g.P("}")
	f.g.P("factories := &", f.service.Unexported(f.service.FactoriesName()), "{")
	f.g.P("transports: transports,")
	f.g.P("}")
	f.g.P("endpointer := &", f.service.Unexported(f.service.EndpointersName()), "{")
	f.g.P("target:    target,")
	f.g.P("builder:   options.Builder(),")
	f.g.P("factories: factories,")
	f.g.P("logger:    options.Logger(),")
	f.g.P("options:   options.EndpointerOptions(),")
	f.g.P("}")
	f.g.P("balancers := &", f.service.Unexported(f.service.BalancersName()), "{")
	f.g.P("factory:    options.BalancerFactory(),")
	f.g.P("endpointer: endpointer,")
	f.g.P("}")
	f.g.P("endpoints := &", f.service.Unexported(f.service.ClientEndpointsName()), "{")
	f.g.P("balancers: balancers,")
	f.g.P("}")

	f.g.P("return &", f.service.Unexported(f.service.ClientServiceName()), "{")
	f.g.P("endpoints:     endpoints,")
	f.g.P("transportName: ", internal.HttpxTransportxPackage.Ident("HttpClient"), ",")
	f.g.P("}")
	f.g.P("}")
	f.g.P()
}
