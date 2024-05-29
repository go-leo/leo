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

func (f *Generator) Generate(g *protogen.GeneratedFile) error {
	for _, service := range f.Services {
		if err := f.GenerateServerTransports(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateClientTransports(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateImplementedServerTransports(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateImplementedClientTransports(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateServer(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateClient(service, g); err != nil {
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
	return nil
}

func (f *Generator) GenerateClientTransports(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.GrpcClientTransportsName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "() *", internal.GrpcTransportPackage.Ident("Client"))
	}
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateImplementedServerTransports(service *internal.Service, g *protogen.GeneratedFile) error {
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

	g.P("func New", service.GrpcServerTransportsName(), "(endpoints ", service.EndpointsName(), ", serverOptions ...", internal.GrpcTransportPackage.Ident("ServerOption"), ") ", service.GrpcServerTransportsName(), " {")
	g.P("return &", service.UnexportedGrpcServerTransportsName(), "{")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), ":", internal.GrpcTransportPackage.Ident("NewServer"), "(")
		g.P("endpoints.", endpoint.Name(), "(), ")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil },")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil },")
		g.P("serverOptions...,")

		g.P("),")
	}
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateImplementedClientTransports(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedGrpcClientTransportsName(), " struct {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), " *", internal.GrpcTransportPackage.Ident("Client"))
	}
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (t *", service.UnexportedGrpcClientTransportsName(), ") ", endpoint.Name(), "() *", internal.GrpcTransportPackage.Ident("Client"), "{")
		g.P("return t.", endpoint.UnexportedName())
		g.P("}")
		g.P()
	}

	g.P("func New", service.GrpcClientTransportsName(), "(conn *", internal.GrpcPackage.Ident("ClientConn"), ", clientOptions ...", internal.GrpcTransportPackage.Ident("ClientOption"), ") ", service.GrpcClientTransportsName(), " {")
	g.P("return &", service.UnexportedGrpcClientTransportsName(), "{")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), ":", internal.GrpcTransportPackage.Ident("NewClient"), "(")
		g.P("conn, ")
		g.P(strconv.Quote(service.FullName()), ",")
		g.P(strconv.Quote(endpoint.Name()), ", ")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil }", ", ")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil }", ", ")
		g.P(endpoint.OutputGoIdent(), "{},")
		g.P("clientOptions...,")
		g.P("),")
	}
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateServer(service *internal.Service, g *protogen.GeneratedFile) error {
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

	g.P("func New", service.GrpcServerName(), "(transports ", service.GrpcServerTransportsName(), ") ", service.ServiceName(), " {")
	g.P("return &", service.UnexportedGrpcServerName(), "{")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), ": transports.", endpoint.Name(), "(),")
	}
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClient(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedGrpcClientName(), " struct {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), " ", internal.EndpointPackage.Ident("Endpoint"))
	}
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (c *", service.UnexportedGrpcClientName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		g.P("rep, err := c.", endpoint.UnexportedName(), "(ctx, request)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
		g.P("}")
		g.P()
	}

	g.P("func New", service.GrpcClientName(), "(transports ", service.GrpcClientTransportsName(), ", middlewares ...", internal.EndpointPackage.Ident("Middleware"), ") ", service.ServiceName(), " {")
	g.P("return &", service.UnexportedGrpcClientName(), "{")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), ":", internal.EndpointxPackage.Ident("Chain"), "(transports.", endpoint.Name(), "().Endpoint(), middlewares...),")
	}
	g.P("}")
	g.P("}")
	g.P()
	return nil
}
