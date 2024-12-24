package core

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
		if err := f.GenerateBalancers(service, g); err != nil {
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
	for _, service := range f.Services {
		if err := f.GenerateBalancersImplements(service, g); err != nil {
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
	g.P("type ", service.ServerEndpointsName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ") ", internal.EndpointPackage.Ident("Endpoint"))
	}
	g.P("}")
	g.P()
	g.P("type ", service.ClientEndpointsName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ") (", internal.EndpointPackage.Ident("Endpoint"), ", error)")
	}
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClientTransports(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.ClientTransportsName(), " interface {")
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
		g.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", color string) (", internal.SdPackage.Ident("Endpointer"), ", error)")
	}
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateBalancers(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.BalancersName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ") (", internal.LbPackage.Ident("Balancer"), ", error)")
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
	g.P("func new", service.ServerEndpointsName(), "(svc ", service.ServiceName(), ", middlewares ...", internal.EndpointPackage.Ident("Middleware"), ") ", service.ServerEndpointsName(), "{")
	g.P("return &", service.UnexportedServerEndpointsName(), "{svc: svc, middlewares: middlewares}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClientEndpoints(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedClientEndpointsName(), " struct {")
	g.P("balancers ", service.BalancersName())
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (e *", service.UnexportedClientEndpointsName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ") (", internal.EndpointPackage.Ident("Endpoint"), ", error) {")
		g.P("balancer, err := e.balancers.", endpoint.Name(), "(ctx)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("return balancer.Endpoint()")
		g.P("}")
	}
	g.P("func new", service.ClientEndpointsName(), "(")
	g.P("target string,")
	g.P("transports ", service.ClientTransportsName(), ",")
	g.P("instancerFactory ", internal.SdxPackage.Ident("InstancerFactory"), ",")
	g.P("endpointerOptions []", internal.SdPackage.Ident("EndpointerOption"), ",")
	g.P("balancerFactory ", internal.LbxPackage.Ident("BalancerFactory"), ",")
	g.P("logger ", internal.LogPackage.Ident("Logger"), ",")
	g.P(") ", service.ClientEndpointsName(), " {")
	g.P("factories := new", service.FactoriesName(), "(transports)")
	g.P("endpointers := new", service.EndpointersName(), "(target, instancerFactory, factories, logger, endpointerOptions...)")
	g.P("balancers := new", service.BalancersName(), "(balancerFactory, endpointers)")
	g.P("return &", service.UnexportedClientEndpointsName(), "{balancers: balancers}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClientTransportsImplements(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedFactoriesName(), " struct {")
	g.P("transports ", service.ClientTransportsName())
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (f *", service.UnexportedFactoriesName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ") ", internal.SdPackage.Ident("Factory"), "{")
		g.P("return func(instance string) (", internal.EndpointPackage.Ident("Endpoint"), ", ", internal.IOPackage.Ident("Closer"), ", error) {")
		g.P("return f.transports.", endpoint.Name(), "(ctx, instance)")
		g.P("}")
		g.P("}")
	}
	g.P("func new", service.FactoriesName(), "(transports ", service.ClientTransportsName(), ") ", service.FactoriesName(), " {")
	g.P("return &", service.UnexportedFactoriesName(), "{transports: transports}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateEndpointersImplements(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedEndpointersName(), " struct {")
	g.P("target string")
	g.P("instancerFactory ", internal.SdxPackage.Ident("InstancerFactory"))
	g.P("factories ", service.FactoriesName())
	g.P("logger ", internal.LogPackage.Ident("Logger"))
	g.P("options []", internal.SdPackage.Ident("EndpointerOption"))
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (e *", service.UnexportedEndpointersName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", color string) (", internal.SdPackage.Ident("Endpointer"), ", error) {")
		g.P("return ", internal.SdxPackage.Ident("NewEndpointer"), "(ctx, e.target, color, e.instancerFactory, e.factories.", endpoint.Name(), "(ctx), e.logger, e.options...)")
		g.P("}")
	}
	g.P("func new", service.EndpointersName(), "(")
	g.P("target string,")
	g.P("instancerFactory ", internal.SdxPackage.Ident("InstancerFactory"), ",")
	g.P("factories ", service.FactoriesName(), ",")
	g.P("logger ", internal.LogPackage.Ident("Logger"), ",")
	g.P("options ...", internal.SdPackage.Ident("EndpointerOption"), ",")
	g.P(")", service.EndpointersName(), " {")
	g.P("return &", service.UnexportedEndpointersName(), "{")
	g.P("target: target,")
	g.P("instancerFactory: instancerFactory,")
	g.P("factories: factories,")
	g.P("logger: logger,")
	g.P("options: options,")
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateBalancersImplements(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedBalancersName(), " struct {")
	g.P("factory ", internal.LbxPackage.Ident("BalancerFactory"))
	g.P("endpointer ", service.EndpointersName())
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), " ", internal.LazyLoadxPackage.Ident("Group"), "[", internal.LbPackage.Ident("Balancer"), "]")
	}
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (b *", service.UnexportedBalancersName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ") (", internal.LbPackage.Ident("Balancer"), ", error) {")
		g.P("color, _ := ", internal.StainxPackage.Ident("ExtractColor"), "(ctx)")
		g.P("balancer, err, _ := b.", endpoint.UnexportedName(), ".LoadOrNew(color, ", internal.LbxPackage.Ident("NewBalancer"), "(ctx, b.factory, b.endpointer.", endpoint.Name(), "))")
		g.P("return balancer, err")
		g.P("}")
	}
	g.P("func new", service.BalancersName(), "(factory ", internal.LbxPackage.Ident("BalancerFactory"), ", endpointer ", service.EndpointersName(), ") ", service.BalancersName(), " {")
	g.P("return &", service.UnexportedBalancersName(), "{")
	g.P("factory: factory,")
	g.P("endpointer: endpointer,")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), ": ", internal.LazyLoadxPackage.Ident("Group"), "[", internal.LbPackage.Ident("Balancer"), "]{},")
	}
	g.P("}")
	g.P("}")
	return nil
}

func (f *Generator) GenerateClientService(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedClientServiceName(), " struct {")
	g.P("endpoints ", service.ClientEndpointsName())
	g.P("transportName string")
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (c *", service.UnexportedClientServiceName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		g.P("ctx = ", internal.EndpointxPackage.Ident("InjectName"), "(ctx, ", strconv.Quote(endpoint.FullName()), ")")
		g.P("ctx = ", internal.TransportxPackage.Ident("InjectName"), "(ctx, c.transportName)")
		g.P("endpoint, err := c.endpoints.", endpoint.Name(), "(ctx)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("rep, err := endpoint(ctx, request)")
		g.P("if err != nil {")
		g.P("return nil, ", internal.StatusxPackage.Ident("From"), "(err)")
		g.P("}")
		g.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
		g.P("}")
	}
	g.P("func new", service.ClientServiceName(), "(endpoints ", service.ClientEndpointsName(), ", transportName string) ", service.ServiceName(), " {")
	g.P("return &", service.UnexportedClientServiceName(), "{endpoints: endpoints, transportName: transportName}")
	g.P("}")
	return nil
}
