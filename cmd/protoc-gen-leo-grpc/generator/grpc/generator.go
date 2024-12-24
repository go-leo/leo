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
	g.P("type ", service.UnexportedGrpcServerTransportsName(), " struct {")
	g.P("endpoints ", service.EndpointsName())
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (t *", service.UnexportedGrpcServerTransportsName(), ") ", endpoint.Name(), "() ", internal.GrpcTransportPackage.Ident("Handler"), "{")
		g.P("return ", internal.GrpcTransportPackage.Ident("NewServer"), "(")
		g.P("t.endpoints.", endpoint.Name(), "(", internal.ContextPackage.Ident("TODO"), "()), ")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil },")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil },")
		g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxTransportxPackage.Ident("ServerEndpointInjector"), "(", strconv.Quote(endpoint.FullName()), ")),")
		g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxTransportxPackage.Ident("ServerTransportInjector"), "),")
		g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxTransportxPackage.Ident("IncomingMetadataInjector"), "),")
		g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxTransportxPackage.Ident("IncomingStain"), "),")
		g.P(")")
		g.P("}")
		g.P()
	}
	return nil
}

func (f *Generator) GenerateServerService(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedGrpcServerName(), " struct {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), " ", internal.GrpcTransportPackage.Ident("Handler"))
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
	g.P("func New", service.GrpcServerName(), "(svc ", service.ServiceName(), ", middlewares ...", internal.EndpointPackage.Ident("Middleware"), ") ", service.ServiceName(), " {")
	g.P("endpoints := new", service.ServerEndpointsName(), "(svc, middlewares...)")
	g.P("transports := &", service.UnexportedGrpcServerTransportsName(), "{endpoints: endpoints}")
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

func (f *Generator) GenerateClientService(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("func New", service.GrpcClientName(), "(target string, opts ...", internal.GrpcxTransportxPackage.Ident("ClientOption"), ") ", service.ServiceName(), " {")
	g.P("options := ", internal.GrpcxTransportxPackage.Ident("NewClientOptions"), "(opts...)")
	g.P("transports := new", service.GrpcClientTransportsName(), "(options.DialOptions(), options.ClientTransportOptions(), options.Middlewares())")
	g.P("endpoints := new", service.ClientEndpointsName(), "(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())")
	g.P("return new", service.ClientServiceName(), "(endpoints, ", internal.GrpcxTransportxPackage.Ident("GrpcClient"), ")")
	g.P("}")
	g.P()
	return nil
}
