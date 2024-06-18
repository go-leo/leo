package grpc

import (
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo/generator/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"strconv"
)

type Generator struct {
	Plugin   *protogen.Plugin
	File     *protogen.File
	Services []*internal.Service
}

func NewGenerator(plugin *protogen.Plugin, file *protogen.File) (*Generator, error) {
	services, err := internal.NewServices(file)
	if err != nil {
		return nil, err
	}
	return &Generator{Plugin: plugin, File: file, Services: services}, nil
}

func (f *Generator) GenerateServer(g *protogen.GeneratedFile) error {
	for _, service := range f.Services {
		if err := f.GenerateServerTransports(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateServerService(service, g); err != nil {
			return err
		}
	}
	return nil
}

func (f *Generator) GenerateClient(g *protogen.GeneratedFile) error {
	for _, service := range f.Services {
		if err := f.GenerateClientTransports(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateClientService(service, g); err != nil {
			return err
		}
	}

	return nil
}

func (f *Generator) GenerateServerTransports(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.GrpcServerTransportsName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "() *", internal.GrpcTransportPackage.Ident("Server"))
	}
	g.P("}")
	g.P()

	g.P("type ", service.UnexportedGrpcServerTransportsName(), " struct {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), " *", internal.GrpcTransportPackage.Ident("Server"))
	}
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (t *", service.UnexportedGrpcServerTransportsName(), ") ", endpoint.Name(), "() *", internal.GrpcTransportPackage.Ident("Server"), "{")
		g.P("return t.", endpoint.UnexportedName())
		g.P("}")
		g.P()
	}

	g.P("func New", service.GrpcServerTransportsName(), "(endpoints ", service.EndpointsName(), ") ", service.GrpcServerTransportsName(), " {")
	g.P("return &", service.UnexportedGrpcServerTransportsName(), "{")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), ":", internal.GrpcTransportPackage.Ident("NewServer"), "(")
		g.P("endpoints.", endpoint.Name(), "(", internal.ContextPackage.Ident("TODO"), "()), ")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil },")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil },")

		g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxTransportxPackage.Ident("ServerEndpointInjector"), "(", strconv.Quote(endpoint.FullName()), ")),")
		g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxTransportxPackage.Ident("ServerTransportInjector"), "),")
		g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxTransportxPackage.Ident("IncomingMetadata"), "),")

		g.P("),")
	}
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateServerService(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedGrpcServerName(), " struct {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), " *", internal.GrpcTransportPackage.Ident("Server"))
	}
	g.P("}")
	g.P()

	for _, endpoint := range service.Endpoints {
		g.P("func (s *", service.UnexportedGrpcServerName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		g.P("ctx, rep, err := s.", endpoint.UnexportedName(), ".ServeGRPC(ctx, request)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("_ = ctx")
		g.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
		g.P("}")
		g.P()
	}
	g.P()

	g.P("func New", service.GrpcServerName(), "(svc ", service.ServiceName(), ", middlewares ...", internal.EndpointPackage.Ident("Middleware"), ") ", service.ServiceName(), " {")
	g.P("endpoints := New", service.ServerEndpointsName(), "(svc, middlewares...)")
	g.P("transports := New", service.GrpcServerTransportsName(), "(endpoints)")
	g.P("return &", service.UnexportedGrpcServerName(), "{")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), ": transports.", endpoint.Name(), "(),")
	}
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClientTransports(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedGrpcClientTransportsName(), " struct {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), " ", internal.TransportxPackage.Ident("ClientTransport"))
	}
	g.P("}")
	g.P()

	for _, endpoint := range service.Endpoints {
		g.P("func (t *", service.UnexportedGrpcClientTransportsName(), ") ", endpoint.Name(), "() ", internal.TransportxPackage.Ident("ClientTransport"), "{")
		g.P("return t.", endpoint.UnexportedName())
		g.P("}")
		g.P()
	}

	g.P("func New", service.GrpcClientTransportsName(), "(")
	g.P("target string,")
	g.P("dialOption []", internal.GrpcPackage.Ident("DialOption"), ",")
	g.P("options ...", internal.TransportxPackage.Ident("ClientTransportOption"), ",")
	g.P(") (", service.ClientTransportsName(), ", error) {")
	g.P("t := &", service.UnexportedGrpcClientTransportsName(), "{}")
	g.P("var err error")
	for _, endpoint := range service.Endpoints {
		g.P("t.", endpoint.UnexportedName(), ", err = ", internal.ErrorxPackage.Ident("Break"), "[", internal.TransportxPackage.Ident("ClientTransport"), "](err)(func() (", internal.TransportxPackage.Ident("ClientTransport"), ", error) {")
		g.P("return ", internal.TransportxPackage.Ident("NewClientTransport"), "(")
		g.P("target,")
		g.P(internal.GrpcxTransportxPackage.Ident("ClientFactory"), "(")
		g.P("dialOption, ")
		g.P(strconv.Quote(service.FullName()), ",")
		g.P(strconv.Quote(endpoint.Name()), ", ")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil }", ", ")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil }", ", ")
		g.P(endpoint.OutputGoIdent(), "{},")
		g.P(internal.GrpcTransportPackage.Ident("ClientBefore"), "(", internal.GrpcxTransportxPackage.Ident("OutgoingMetadata"), "),")
		g.P("),")
		g.P("options...,")
		g.P(")")
		g.P("})")
	}
	g.P("return t, err")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClientService(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedGrpcClientName(), " struct {")
	g.P("endpoints ", service.EndpointsName())
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (c *", service.UnexportedGrpcClientName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		g.P("ctx = ", internal.EndpointxPackage.Ident("InjectName"), "(ctx, ", strconv.Quote(endpoint.FullName()), ")")
		g.P("ctx = ", internal.TransportxPackage.Ident("InjectName"), "(ctx, ", internal.GrpcxTransportxPackage.Ident("GrpcClient"), ")")
		g.P("rep, err := c.endpoints.", endpoint.Name(), "(ctx)(ctx, request)")
		g.P("if err != nil {")
		g.P("return nil, ", internal.StatusxPackage.Ident("FromGrpcError"), "(err)")
		g.P("}")
		g.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
		g.P("}")
		g.P()
	}

	g.P("func New", service.GrpcClientName(), "(transports ", service.ClientTransportsName(), ", middlewares ...", internal.EndpointPackage.Ident("Middleware"), ") ", service.ServiceName(), " {")
	g.P("endpoints := New", service.ClientEndpointsName(), "(transports, middlewares...)")
	g.P("return &", service.UnexportedGrpcClientName(), "{endpoints:endpoints}")
	g.P("}")
	g.P()
	return nil
}
