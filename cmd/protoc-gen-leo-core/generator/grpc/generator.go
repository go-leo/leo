package grpc

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
	services, err := internal.NewServices(file)
	if err != nil {
		return nil, err
	}
	return &Generator{Plugin: plugin, File: file, Services: services}, nil
}

func (f *Generator) Generate(g *protogen.GeneratedFile) error {
	for _, service := range f.Services {
		if err := f.GenerateServer(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateNewServer(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateClient(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateNewClient(service, g); err != nil {
			return err
		}
	}

	return nil
}

func (f *Generator) GenerateServer(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("type ", service.UnexportedGRPCServerName(), " struct {")
	generatedFile.P()
	if *Conf.RequireUnimplemented {
		generatedFile.P(service.UnimplementedServerName())
		generatedFile.P()
	}
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.UnexportedName(), " ", internal.GrpcTransportPackage.Ident("Handler"))
		generatedFile.P()
	}
	generatedFile.P("}")
	generatedFile.P()

	for _, endpoint := range service.Endpoints {
		generatedFile.P("func (s *", service.UnexportedGRPCServerName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		generatedFile.P("ctx, rep, err := s.", endpoint.UnexportedName(), ".ServeGRPC(ctx, request)")
		generatedFile.P("if err != nil {")
		generatedFile.P("return nil, err")
		generatedFile.P("}")
		generatedFile.P("_ = ctx")
		generatedFile.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
		generatedFile.P("}")
		generatedFile.P()
	}
	return nil
}

func (f *Generator) GenerateNewServer(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("func New", service.GRPCServerName(), "(")
	generatedFile.P("endpoints interface {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.Name(), "() ", internal.EndpointPackage.Ident("Endpoint"))
	}
	generatedFile.P("},")
	generatedFile.P("opts []", internal.GrpcTransportPackage.Ident("ServerOption"), ",")
	generatedFile.P("middlewares ...", internal.EndpointPackage.Ident("Middleware"), ",")
	generatedFile.P(") interface {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error)")
	}
	generatedFile.P("} {")
	generatedFile.P("return &", service.UnexportedGRPCServerName(), "{")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.UnexportedName(), ":    ", internal.GrpcTransportPackage.Ident("NewServer"), "(")
		generatedFile.P(internal.EndpointxPackage.Ident("Chain"), "(endpoints.", endpoint.Name(), "(), middlewares...), ")
		generatedFile.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil },")
		generatedFile.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil },")
		generatedFile.P("opts...,")
		generatedFile.P("),")
	}
	generatedFile.P("}")
	generatedFile.P("}")
	generatedFile.P()
	return nil
}

func (f *Generator) GenerateClient(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("type ", service.UnexportedGRPCClientName(), " struct {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.UnexportedName(), " ", internal.EndpointPackage.Ident("Endpoint"))
	}
	generatedFile.P("}")
	generatedFile.P()
	for _, endpoint := range service.Endpoints {
		generatedFile.P("func (c *", service.UnexportedGRPCClientName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		generatedFile.P("rep, err := c.", endpoint.UnexportedName(), "(ctx, request)")
		generatedFile.P("if err != nil {")
		generatedFile.P("return nil, err")
		generatedFile.P("}")
		generatedFile.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
		generatedFile.P("}")
		generatedFile.P()
	}
	return nil
}

func (f *Generator) GenerateNewClient(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("func New", service.GRPCClientName(), "(")
	generatedFile.P("conn *", internal.GrpcPackage.Ident("ClientConn"), ",")
	generatedFile.P("opts []", internal.GrpcTransportPackage.Ident("ClientOption"), ",")
	generatedFile.P("middlewares ...", internal.EndpointPackage.Ident("Middleware"), ",")
	generatedFile.P(") interface {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error)")
	}
	generatedFile.P("} {")
	generatedFile.P("return &", service.UnexportedGRPCClientName(), "{")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.UnexportedName(), ":    ", internal.EndpointxPackage.Ident("Chain"), "(")
		generatedFile.P(internal.GrpcTransportPackage.Ident("NewClient"), "(")
		generatedFile.P("conn, ")
		generatedFile.P(strconv.Quote(service.FullName()), ",")
		generatedFile.P(strconv.Quote(endpoint.Name()), ", ")
		generatedFile.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil }", ", ")
		generatedFile.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil }", ", ")
		generatedFile.P(endpoint.OutputGoIdent(), "{},")
		generatedFile.P("opts...,")
		generatedFile.P(").Endpoint(),")
		generatedFile.P("middlewares...),")
	}
	generatedFile.P("}")
	generatedFile.P("}")
	generatedFile.P()
	return nil
}
