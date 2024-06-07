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
		if err := f.GenerateClientFactory(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := f.GenerateClientEndpoints(service, g); err != nil {
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
		g.P("endpoints.", endpoint.Name(), "(), ")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil },")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil },")

		g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxPackage.Ident("ServerEndpointInjector"), "(", strconv.Quote(endpoint.FullName()), ")),")
		g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxPackage.Ident("ServerTransportInjector"), "),")
		g.P(internal.GrpcTransportPackage.Ident("ServerBefore"), "(", internal.GrpcxPackage.Ident("IncomingMetadata"), "),")

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

func (f *Generator) GenerateClientTransports(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.GrpcClientTransportsName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "() *", internal.GrpcTransportPackage.Ident("Client"))
	}
	g.P("}")
	g.P()

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

	g.P("func New", service.GrpcClientTransportsName(), "(conn *", internal.GrpcPackage.Ident("ClientConn"), ") ", service.GrpcClientTransportsName(), " {")
	g.P("return &", service.UnexportedGrpcClientTransportsName(), "{")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), ":", internal.GrpcTransportPackage.Ident("NewClient"), "(")
		g.P("conn, ")
		g.P(strconv.Quote(service.FullName()), ",")
		g.P(strconv.Quote(endpoint.Name()), ", ")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil }", ", ")
		g.P("func(_ ", internal.ContextPackage.Ident("Context"), ", v any) (any, error) { return v, nil }", ", ")
		g.P(endpoint.OutputGoIdent(), "{},")

		g.P(internal.GrpcTransportPackage.Ident("ClientBefore"), "(", internal.GrpcxPackage.Ident("OutgoingMetadata"), "),")

		g.P("),")
	}
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClientFactory(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.GrpcClientFactoriesName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "() ", internal.SdPackage.Ident("Factory"))
	}
	g.P("}")
	g.P()

	g.P("type ", service.UnexportedGrpcClientFactoriesName(), " struct {")
	g.P("endpoints func(transports ", service.GrpcClientTransportsName(), ")", service.EndpointsName())
	g.P("opts      []", internal.GrpcPackage.Ident("DialOption"))
	g.P("}")
	g.P()

	for _, endpoint := range service.Endpoints {
		g.P("func (f *", service.UnexportedGrpcClientFactoriesName(), ") ", endpoint.Name(), "() ", internal.SdPackage.Ident("Factory"), "{")
		g.P("return func(instance string) (", internal.EndpointPackage.Ident("Endpoint"), ", ", internal.IOPackage.Ident("Closer"), ", error) {")
		g.P("conn, err := ", internal.GrpcPackage.Ident("NewClient"), "(instance, f.opts...)")
		g.P("if err != nil {")
		g.P("return nil, nil, err")
		g.P("}")
		g.P("endpoints := f.endpoints(", "New", service.GrpcClientTransportsName(), "(conn))")
		g.P("return endpoints.", endpoint.Name(), "(), conn, nil")
		g.P("}")
		g.P("}")
		g.P()
	}

	g.P("func New", service.GrpcClientFactoriesName(), "(", "endpoints func(transports ", service.GrpcClientTransportsName(), ")", service.EndpointsName(), ", opts ...", internal.GrpcPackage.Ident("DialOption"), ") ", service.GrpcClientFactoriesName(), " {")
	g.P("return &", service.UnexportedGrpcClientFactoriesName(), "{endpoints: endpoints, opts: opts}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClientEndpoints(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedGrpcClientEndpointsName(), " struct {")
	g.P("transports ", service.GrpcClientTransportsName())
	g.P("middlewares []", internal.EndpointPackage.Ident("Middleware"))
	g.P("}")
	g.P()

	for _, endpoint := range service.Endpoints {
		g.P("func (e *", service.UnexportedGrpcClientEndpointsName(), ") ", endpoint.Name(), "() ", internal.EndpointPackage.Ident("Endpoint"), "{")
		g.P("return ", internal.EndpointxPackage.Ident("Chain"), "(e.transports.", endpoint.Name(), "().Endpoint(), e.middlewares...)")
		g.P("}")
		g.P()
	}

	g.P("func New", service.GrpcClientEndpointsName(), "(", "middlewares ...", internal.EndpointPackage.Ident("Middleware"), ") func(transports ", service.GrpcClientTransportsName(), ")", service.EndpointsName(), " {")
	g.P("return func(transports ", service.GrpcClientTransportsName(), ")", service.EndpointsName(), " {")
	g.P("return &", service.UnexportedGrpcClientEndpointsName(), "{transports: transports, middlewares: middlewares}")
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func (f *Generator) GenerateClientService(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedGrpcClientName(), " struct {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), " ", internal.EndpointPackage.Ident("Endpoint"))
	}
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (c *", service.UnexportedGrpcClientName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		g.P("ctx = ", internal.EndpointxPackage.Ident("InjectName"), "(ctx, ", strconv.Quote(endpoint.FullName()), ")")
		g.P("ctx = ", internal.TransportxPackage.Ident("InjectName"), "(ctx, ", internal.GrpcxPackage.Ident("GrpcClient"), ")")
		g.P("rep, err := c.", endpoint.UnexportedName(), "(ctx, request)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
		g.P("}")
		g.P()
	}

	g.P("func New", service.GrpcClientName(), "(endpoints ", service.EndpointsName(), ") ", service.ServiceName(), " {")
	g.P("return &", service.UnexportedGrpcClientName(), "{")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), ": ", "endpoints.", endpoint.Name(), "(),")
	}
	g.P("}")
	g.P("}")
	g.P()
	return nil
}
