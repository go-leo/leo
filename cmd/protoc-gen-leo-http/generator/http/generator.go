package http

import (
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"strconv"
)

type Generator struct {
	Plugin   *protogen.Plugin
	File     *protogen.File
	Services []*internal.Service
}

func NewGenerator(plugin *protogen.Plugin, file *protogen.File) (*Generator, error) {
	services, err := internal.NewHttpServices(file)
	if err != nil {
		return nil, err
	}
	return &Generator{Plugin: plugin, File: file, Services: services}, nil
}

func (f *Generator) GenerateFunc(g *protogen.GeneratedFile) error {
	for _, service := range f.Services {
		if err := f.GenerateAppendRoutesFunc(service, g); err != nil {
			return err
		}
		if err := f.GenerateAppendServerFunc(service, g); err != nil {
			return err
		}
		if err := f.GenerateNewClientFunc(service, g); err != nil {
			return err
		}
	}
	return nil
}

func (f *Generator) GenerateAppendRoutesFunc(service *internal.Service, g *protogen.GeneratedFile) error {
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

func (f *Generator) GenerateAppendServerFunc(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("func Append", service.HttpRoutesName(), "(router *", internal.MuxPackage.Ident("Router"), ", svc ", service.ServiceName(), ", middlewares ...", internal.EndpointPackage.Ident("Middleware"), ") ", "*", internal.MuxPackage.Ident("Router"), " {")
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

func (f *Generator) GenerateNewClientFunc(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("func New", service.HttpClientName(), "(target string, opts ...", internal.HttpxTransportxPackage.Ident("ClientOption"), ") ", service.ServiceName(), " {")
	g.P("options := ", internal.HttpxTransportxPackage.Ident("NewClientOptions"), "(opts...)")
	g.P("transports := new", service.HttpClientTransportsName(), "(options.Scheme(), options.ClientTransportOptions(), options.Middlewares())")
	g.P("endpoints := new", service.ClientEndpointsName(), "(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())")
	g.P("return new", service.ClientServiceName(), "(endpoints, ", internal.HttpxTransportxPackage.Ident("HttpClient"), ")")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateServer(g *protogen.GeneratedFile) error {
	serverTransportsGenerator := ServerTransportsGenerator{}
	serverRequestDecoderGenerator := ServerRequestDecoderGenerator{}
	serverResponseEncoderGenerator := ServerResponseEncoderGenerator{}
	client := ClientGenerator{}
	clientRequestEncoderGenerator := ClientRequestEncoderGenerator{}
	clientResponseDecoderGenerator := ClientResponseDecoderGenerator{}
	for _, service := range f.Services {
		if err := serverTransportsGenerator.GenerateTransports(service, g); err != nil {
			return err
		}
		if err := serverRequestDecoderGenerator.GenerateServerRequestDecoder(service, g); err != nil {
			return err
		}
		if err := serverResponseEncoderGenerator.GenerateServerResponseEncoder(service, g); err != nil {
			return err
		}
		if err := clientRequestEncoderGenerator.GenerateClientRequestEncoder(service, g); err != nil {
			return err
		}
		if err := clientResponseDecoderGenerator.GenerateClientResponseDecoder(service, g); err != nil {
			return err
		}
		if err := serverTransportsGenerator.GenerateTransportsImplements(service, g); err != nil {
			return err
		}
		if err := serverRequestDecoderGenerator.GenerateServerRequestDecoderImplements(service, g); err != nil {
			return err
		}
		if err := serverResponseEncoderGenerator.GenerateServerResponseEncoderImplements(service, g); err != nil {
			return err
		}
		if err := client.GenerateTransports(service, g); err != nil {
			return err
		}
		if err := clientResponseDecoderGenerator.GenerateClientResponseDecoderImplements(service, g); err != nil {
			return err
		}
	}
	return nil
}

func (f *Generator) GenerateCoder(g *protogen.GeneratedFile) error {
	client := ClientGenerator{}
	for _, service := range f.Services {
		for _, endpoint := range service.Endpoints {
			if err := client.PrintEncodeRequestFunc(g, endpoint); err != nil {
				return err
			}
		}
	}
	return nil
}
