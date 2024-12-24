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

	g.P("func new", service.GrpcServerTransportsName(), "(endpoints ", service.EndpointsName(), ") ", service.GrpcServerTransportsName(), " {")
	g.P("return &", service.UnexportedGrpcServerTransportsName(), "{")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), ":", endpoint.GrpcServerTransportName(), "(endpoints),")
	}
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateServerEndpointTransport(service *internal.Service, g *protogen.GeneratedFile, endpoint *internal.Endpoint) error {
	g.P("func ", endpoint.GrpcServerTransportName(), "(endpoints ", service.EndpointsName(), ") *", internal.GrpcTransportPackage.Ident("Server"), " {")
	g.P("return ", internal.GrpcTransportPackage.Ident("NewServer"), "(")
	g.P("endpoints.", endpoint.Name(), "(", internal.ContextPackage.Ident("TODO"), "()), ")
	g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil },")
	g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil },")
	g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxTransportxPackage.Ident("ServerEndpointInjector"), "(", strconv.Quote(endpoint.FullName()), ")),")
	g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxTransportxPackage.Ident("ServerTransportInjector"), "),")
	g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxTransportxPackage.Ident("IncomingMetadataInjector"), "),")
	g.P(")")
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
	g.P("endpoints := new", service.ServerEndpointsName(), "(svc, middlewares...)")
	g.P("transports := new", service.GrpcServerTransportsName(), "(endpoints)")
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
	g.P("dialOptions []", internal.GrpcPackage.Ident("DialOption"))
	g.P("clientOptions []", internal.GrpcTransportPackage.Ident("ClientOption"))
	g.P("middlewares []", internal.EndpointPackage.Ident("Middleware"))
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (t *", service.UnexportedGrpcClientTransportsName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", instance string) (", internal.EndpointPackage.Ident("Endpoint"), ", ", internal.IOPackage.Ident("Closer"), ", error) {")
		g.P("conn, err := ", internal.GrpcPackage.Ident("NewClient"), "(instance, t.dialOptions...)")
		g.P("if err != nil {")
		g.P("return nil, nil, err")
		g.P("}")
		g.P("opts := []", internal.GrpcTransportPackage.Ident("ClientOption"), "{")
		g.P(internal.GrpcTransportPackage.Ident("ClientBefore"), "(", internal.GrpcxTransportxPackage.Ident("OutgoingMetadataInjector"), "),")
		g.P(internal.GrpcTransportPackage.Ident("ClientBefore"), "(", internal.GrpcxTransportxPackage.Ident("OutgoingStain"), "),")
		g.P("}")
		g.P("opts = append(opts, t.clientOptions...)")
		g.P("client := ", internal.GrpcTransportPackage.Ident("NewClient"), "(")
		g.P("conn,")
		g.P(strconv.Quote(service.FullName()), ",")
		g.P(strconv.Quote(endpoint.Name()), ", ")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil }", ", ")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil }", ", ")
		g.P(endpoint.OutputGoIdent(), "{},")
		g.P("opts...)")
		g.P("return ", internal.EndpointxPackage.Ident("Chain"), "(client.Endpoint(), t.middlewares...), conn, nil")
		g.P("}")
		g.P()
	}
	g.P("func new", service.GrpcClientTransportsName(), "(")
	g.P("dialOptions []", internal.GrpcPackage.Ident("DialOption"), ",")
	g.P("clientOptions []", internal.GrpcTransportPackage.Ident("ClientOption"), ",")
	g.P("middlewares []", internal.EndpointPackage.Ident("Middleware"), ",")
	g.P(") ", service.ClientTransportsName(), " {")
	g.P("return &", service.UnexportedGrpcClientTransportsName(), "{")
	g.P("dialOptions: dialOptions,")
	g.P("clientOptions: clientOptions,")
	g.P("middlewares: middlewares,")
	g.P("}")
	g.P("}")
	g.P()

	return nil
}

func (f *Generator) GenerateTransport(g *protogen.GeneratedFile) error {
	for _, service := range f.Services {
		for _, endpoint := range service.Endpoints {
			if err := f.GenerateServerEndpointTransport(service, g, endpoint); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *Generator) GenerateClientService(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedGrpcClientName(), " struct {")
	g.P("balancers ", service.BalancersName())
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (c *", service.UnexportedGrpcClientName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		g.P("ctx = ", internal.EndpointxPackage.Ident("InjectName"), "(ctx, ", strconv.Quote(endpoint.FullName()), ")")
		g.P("ctx = ", internal.TransportxPackage.Ident("InjectName"), "(ctx, ", internal.GrpcxTransportxPackage.Ident("GrpcClient"), ")")
		g.P("balancer, err := c.balancers.", endpoint.Name(), "(ctx)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("endpoint, err := balancer.Endpoint()")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("rep, err := endpoint(ctx, request)")
		g.P("if err != nil {")
		g.P("return nil, ", internal.StatusxPackage.Ident("FromGrpcError"), "(err)")
		g.P("}")
		g.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
		g.P("}")
		g.P()
	}
	g.P("func New", service.GrpcClientName(), "(target string, opts ...", internal.GrpcxTransportxPackage.Ident("ClientOption"), ") ", service.ServiceName(), " {")
	g.P("options := ", internal.GrpcxTransportxPackage.Ident("NewClientOptions"), "(opts...)")
	g.P("transports := new", service.GrpcClientTransportsName(), "(options.DialOptions(), options.ClientTransportOptions(), options.Middlewares())")
	g.P("factories := new", service.FactoriesName(), "(transports)")
	g.P("endpointers := new", service.EndpointersName(), "(target, options.InstancerFactory(), factories, options.Logger(), options.EndpointerOptions()...)")
	g.P("balancers := new", service.BalancersName(), "(options.BalancerFactory(), endpointers)")
	g.P("return &", service.UnexportedHttpClientName(), "{balancers: balancers}")
	g.P("}")
	g.P()
	return nil
}
