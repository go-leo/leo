package core

import (
	"github.com/go-leo/leo/v3/cmd/internal"
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
		if err := f.GenerateClientTransports(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateFactories(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateEndpointers(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateServerEndpoints(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateClientEndpoints(service, g); err != nil {
			return err
		}
	}
	for _, service := range f.Services {
		if err := f.GenerateClientTransportsImplements(service, g); err != nil {
			return err
		}
	}
	for _, service := range f.Services {
		if err := f.GenerateEndpointersImplements(service, g); err != nil {
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
		g.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ") ", internal.EndpointPackage.Ident("Endpoint"))
	}
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClientTransports(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.ClientTransportsName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "() ", internal.TransportxPackage.Ident("ClientTransport"))
	}
	g.P("}")
	g.P("type ", service.ClientTransportsNameV2(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", instance string) (", internal.EndpointPackage.Ident("Endpoint"), ", ", internal.IOPackage.Ident("Closer"), ", error)")
	}
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateFactories(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.FactoriesName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ") ", internal.SdPackage.Ident("Factory"))
	}
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateEndpointers(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.EndpointersName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ") ", internal.SdPackage.Ident("Endpointer"))
	}
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateServerEndpoints(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedServerEndpointsName(), " struct {")
	g.P("svc ", service.ServiceName())
	g.P("middlewares []", internal.EndpointPackage.Ident("Middleware"))
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (e *", service.UnexportedServerEndpointsName(), ") ", endpoint.Name(), "(", internal.ContextPackage.Ident("Context"), ") ", internal.EndpointPackage.Ident("Endpoint"), "{")
		g.P("component := func(ctx ", internal.ContextPackage.Ident("Context"), ", request any) (any, error) {")
		g.P("return e.svc.", endpoint.Name(), "(ctx, request.(*", endpoint.InputGoIdent(), "))")
		g.P("}")
		g.P("return ", internal.EndpointxPackage.Ident("Chain"), "(component, e.middlewares...)")
		g.P("}")
	}
	g.P("func new", service.ServerEndpointsName(), "(svc ", service.ServiceName(), ", middlewares ...", internal.EndpointPackage.Ident("Middleware"), ") ", service.EndpointsName(), "{")
	g.P("return &", service.UnexportedServerEndpointsName(), "{svc: svc, middlewares: middlewares}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClientEndpoints(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedClientEndpointsName(), " struct {")
	g.P("transports ", service.ClientTransportsName())
	g.P("middlewares []", internal.EndpointPackage.Ident("Middleware"))
	g.P("}")

	for _, endpoint := range service.Endpoints {
		g.P("func (e *", service.UnexportedClientEndpointsName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ") ", internal.EndpointPackage.Ident("Endpoint"), "{")
		g.P("return ", internal.EndpointxPackage.Ident("Chain"), "(e.transports.", endpoint.Name(), "().Endpoint(ctx), e.middlewares...)")
		g.P("}")
	}
	g.P("func new", service.ClientEndpointsName(), "(transports ", service.ClientTransportsName(), ", middlewares ...", internal.EndpointPackage.Ident("Middleware"), ") ", service.EndpointsName(), " {")
	g.P("return &", service.UnexportedClientEndpointsName(), "{transports: transports, middlewares: middlewares}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClientTransportsImplements(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedFactoriesName(), " struct {")
	g.P("transports ", service.ClientTransportsNameV2())
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (f *", service.UnexportedFactoriesName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ") ", internal.SdPackage.Ident("Factory"), "{")
		g.P("return func(instance string) (", internal.EndpointPackage.Ident("Endpoint"), ", ", internal.IOPackage.Ident("Closer"), ", error) {")
		g.P("return f.transports.", endpoint.Name(), "(ctx, instance)")
		g.P("}")
		g.P("}")
	}
	g.P("func new", service.FactoriesName(), "(transports ", service.ClientTransportsNameV2(), ") ", service.FactoriesName(), " {")
	g.P("return &", service.UnexportedFactoriesName(), "{transports: transports}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateEndpointersImplements(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedEndpointersName(), " struct {")
	g.P("instancer ", internal.SdPackage.Ident("Instancer"))
	g.P("factories ", service.FactoriesName())
	g.P("logger ", internal.LogPackage.Ident("Logger"))
	g.P("options []", internal.SdPackage.Ident("EndpointerOption"))
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (e *", service.UnexportedEndpointersName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ") ", internal.SdPackage.Ident("Endpointer"), " {")
		g.P("return ", internal.SdPackage.Ident("NewEndpointer"), "(e.instancer, e.factories.", endpoint.Name(), "(ctx), e.logger, e.options...)")
		g.P("}")
	}
	g.P("func new", service.EndpointersName(), "(instancer ", internal.SdPackage.Ident("Instancer"), ", factories ", service.FactoriesName(), ", logger ", internal.LogPackage.Ident("Logger"), ", options ...", internal.SdPackage.Ident("EndpointerOption"), ") ", service.EndpointersName(), " {")
	g.P("return &", service.UnexportedEndpointersName(), "{instancer: instancer, factories: factories, logger: logger, options: options}")
	g.P("}")
	g.P()
	return nil
}
