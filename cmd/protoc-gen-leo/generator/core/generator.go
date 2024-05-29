package core

import (
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo/generator/internal"
	"google.golang.org/protobuf/compiler/protogen"
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
		if err := f.GenerateServices(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateEndpoints(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateImplementedEndpoints(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateNewEndpoints(service, g); err != nil {
			return err
		}
	}
	return nil
}

func (f *Generator) GenerateServices(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.ServiceName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error)")
	}
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateEndpoints(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.EndpointsName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "() ", internal.EndpointPackage.Ident("Endpoint"))
	}
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateImplementedEndpoints(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedEndpointsName(), " struct {")
	g.P("svc ", service.ServiceName())
	g.P("middlewares []", internal.EndpointPackage.Ident("Middleware"))
	g.P("}")
	g.P()

	for _, endpoint := range service.Endpoints {
		g.P("func (e *", service.UnexportedEndpointsName(), ") ", endpoint.Name(), "() ", internal.EndpointPackage.Ident("Endpoint"), "{")
		g.P("component := func(ctx ", internal.ContextPackage.Ident("Context"), ", request any) (any, error) {")
		g.P("return e.svc.", endpoint.Name(), "(ctx, request.(*", endpoint.InputGoIdent(), "))")
		g.P("}")
		g.P("return ", internal.EndpointxPackage.Ident("Chain"), "(component, e.middlewares...)")
		g.P("}")
		g.P()
	}
	return nil
}

func (f *Generator) GenerateNewEndpoints(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("func New", service.EndpointsName(), "(svc ", service.ServiceName(), ", middlewares ...", internal.EndpointPackage.Ident("Middleware"), ") ", service.EndpointsName(), "{")
	g.P("return &", service.UnexportedEndpointsName(), "{svc: svc, middlewares: middlewares}")
	g.P("}")
	g.P()
	return nil
}
