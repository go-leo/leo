package core

import (
	"github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo/gen/internal"
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

func (f *Generator) Generate() error {
	// 没有定义服务，则不生成代码
	if len(f.Services) <= 0 {
		return nil
	}

	file := f.File
	filename := file.GeneratedFilenamePrefix + "_leo.core.pb.go"
	g := f.Plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-leo. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	services, err := internal.NewServices(file)
	if err != nil {
		return err
	}

	for _, service := range services {
		f.GenerateServices(service, g)
		//f.GenerateHandlers(service, g)
		f.GenerateServerEndpoints(service, g)
		f.GenerateClientEndpoints(service, g)
		f.GenerateClientTransports(service, g)
		f.GenerateFactories(service, g)
		f.GenerateEndpointers(service, g)
		f.GenerateBalancers(service, g)
		f.GenerateServerEndpointsImplements(service, g)
		f.GenerateClientTransportsImplements(service, g)
		f.GenerateEndpointersImplements(service, g)
		f.GenerateBalancersImplements(service, g)
		f.GenerateClientEndpointsImplements(service, g)
		f.GenerateClientService(service, g)
	}
	return nil
}

func (f *Generator) GenerateServices(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.ServiceName() + " is a service"))
	g.P("type ", service.ServiceName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.Context, ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error)")
	}
	g.P("}")
	g.P()
}

func (f *Generator) GenerateHandlers(service *internal.Service, g *protogen.GeneratedFile) {
	g.P("type (")
	for _, endpoint := range service.Endpoints {
		g.P("", endpoint.HandlerName(), " interface {")
		g.P("Handle(ctx ", internal.Context, ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error)")
		g.P("}")
	}
	g.P(")")
}

func (f *Generator) GenerateServerEndpoints(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.ServerEndpointsName() + " is server endpoints"))
	g.P("type ", service.ServerEndpointsName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.Context, ") ", internal.EndpointPackage.Ident("Endpoint"))
	}
	g.P("}")
	g.P()
}

func (f *Generator) GenerateClientEndpoints(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.ClientEndpointsName() + " is client endpoints"))
	g.P("type ", service.ClientEndpointsName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.Context, ") (", internal.EndpointPackage.Ident("Endpoint"), ", error)")
	}
	g.P("}")
	g.P()
}

func (f *Generator) GenerateClientTransports(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.ClientTransportsName() + " is client transports"))
	g.P("type ", service.ClientTransportsName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.Context, ", instance string) (", internal.EndpointPackage.Ident("Endpoint"), ", ", internal.IOPackage.Ident("Closer"), ", error)")
	}
	g.P("}")
	g.P()
}

func (f *Generator) GenerateFactories(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.FactoriesName() + " is client factories"))
	g.P("type ", service.FactoriesName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.Context, ") ", internal.SdPackage.Ident("Factory"))
	}
	g.P("}")
	g.P()
}

func (f *Generator) GenerateEndpointers(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.EndpointersName() + " is client endpointers"))
	g.P("type ", service.EndpointersName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.Context, ", color string) (", internal.SdPackage.Ident("Endpointer"), ", error)")
	}
	g.P("}")
	g.P()
}

func (f *Generator) GenerateBalancers(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.BalancersName() + " is client balancers"))
	g.P("type ", service.BalancersName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(ctx ", internal.Context, ") (", internal.LbPackage.Ident("Balancer"), ", error)")
	}
	g.P("}")
	g.P()
}

func (f *Generator) GenerateServerEndpointsImplements(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.Unexported(service.ServerEndpointsName())+" implements "), service.ServerEndpointsName())
	g.P("type ", service.Unexported(service.ServerEndpointsName()), " struct {")
	g.P("svc ", service.ServiceName())
	g.P("middlewares []", internal.EndpointPackage.Ident("Middleware"))
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (e *", service.Unexported(service.ServerEndpointsName()), ") ", endpoint.Name(), "(", internal.Context, ") ", internal.EndpointPackage.Ident("Endpoint"), "{")
		g.P("component := func(ctx ", internal.Context, ", request any) (any, error) {")
		g.P("return e.svc.", endpoint.Name(), "(ctx, request.(*", endpoint.InputGoIdent(), "))")
		g.P("}")
		g.P("return ", internal.EndpointxPackage.Ident("Chain"), "(component, e.middlewares...)")
		g.P("}")
		g.P()
	}
}

func (f *Generator) GenerateClientTransportsImplements(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.Unexported(service.FactoriesName())+" implements "), service.FactoriesName())
	g.P("type ", service.Unexported(service.FactoriesName()), " struct {")
	g.P("transports ", service.ClientTransportsName())
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (f *", service.Unexported(service.FactoriesName()), ") ", endpoint.Name(), "(ctx ", internal.Context, ") ", internal.SdPackage.Ident("Factory"), "{")
		g.P("return func(instance string) (", internal.EndpointPackage.Ident("Endpoint"), ", ", internal.IOPackage.Ident("Closer"), ", error) {")
		g.P("return f.transports.", endpoint.Name(), "(ctx, instance)")
		g.P("}")
		g.P("}")
		g.P()
	}
}

func (f *Generator) GenerateEndpointersImplements(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.Unexported(service.EndpointersName())+" implements "), service.EndpointersName())
	g.P("type ", service.Unexported(service.EndpointersName()), " struct {")
	g.P("target string")
	g.P("builder ", internal.SdxPackage.Ident("Builder"))
	g.P("factories ", service.FactoriesName())
	g.P("logger ", internal.LogPackage.Ident("Logger"))
	g.P("options []", internal.SdPackage.Ident("EndpointerOption"))
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (e *", service.Unexported(service.EndpointersName()), ") ", endpoint.Name(), "(ctx ", internal.Context, ", color string) (", internal.SdPackage.Ident("Endpointer"), ", error) {")
		g.P("return ", internal.SdxPackage.Ident("NewEndpointer"), "(ctx, e.target, color, e.builder, e.factories.", endpoint.Name(), "(ctx), e.logger, e.options...)")
		g.P("}")
		g.P()
	}
}

func (f *Generator) GenerateBalancersImplements(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.Unexported(service.BalancersName())+" implements "), service.BalancersName())
	g.P("type ", service.Unexported(service.BalancersName()), " struct {")
	g.P("factory ", internal.LbxPackage.Ident("BalancerFactory"))
	g.P("endpointer ", service.EndpointersName())
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Unexported(endpoint.Name()), " ", internal.LazyLoadxPackage.Ident("Group"), "[", internal.LbPackage.Ident("Balancer"), "]")
	}
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (b *", service.Unexported(service.BalancersName()), ") ", endpoint.Name(), "(ctx ", internal.Context, ") (", internal.LbPackage.Ident("Balancer"), ", error) {")
		g.P("color, _ := ", internal.StainxColorExtractor, "(ctx)")
		g.P("balancer, err, _ := b.", endpoint.Unexported(endpoint.Name()), ".LoadOrNew(color, ", internal.LbxPackage.Ident("NewBalancer"), "(ctx, b.factory, b.endpointer.", endpoint.Name(), "))")
		g.P("return balancer, err")
		g.P("}")
	}
	g.P("func new", service.BalancersName(), "(factory ", internal.LbxPackage.Ident("BalancerFactory"), ", endpointer ", service.EndpointersName(), ") ", service.BalancersName(), " {")
	g.P("return &", service.Unexported(service.BalancersName()), "{")
	g.P("factory: factory,")
	g.P("endpointer: endpointer,")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Unexported(endpoint.Name()), ": ", internal.LazyLoadxPackage.Ident("Group"), "[", internal.LbPackage.Ident("Balancer"), "]{},")
	}
	g.P("}")
	g.P("}")
}

func (f *Generator) GenerateClientEndpointsImplements(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.Unexported(service.ClientEndpointsName())+" implements "), service.ClientEndpointsName())
	g.P("type ", service.Unexported(service.ClientEndpointsName()), " struct {")
	g.P("balancers ", service.BalancersName())
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (e *", service.Unexported(service.ClientEndpointsName()), ") ", endpoint.Name(), "(ctx ", internal.Context, ") (", internal.EndpointPackage.Ident("Endpoint"), ", error) {")
		g.P("balancer, err := e.balancers.", endpoint.Name(), "(ctx)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("return balancer.Endpoint()")
		g.P("}")
		g.P()
	}
}

func (f *Generator) GenerateClientService(service *internal.Service, g *protogen.GeneratedFile) {
	g.P(internal.Comments(service.Unexported(service.ClientServiceName())+" implements "), service.ClientServiceName())
	g.P("type ", service.Unexported(service.ClientServiceName()), " struct {")
	g.P("endpoints ", service.ClientEndpointsName())
	g.P("transportName string")
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (c *", service.Unexported(service.ClientServiceName()), ") ", endpoint.Name(), "(ctx ", internal.Context, ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		g.P("ctx = ", internal.EndpointxNameInjector, "(ctx, ", strconv.Quote(endpoint.FullName()), ")")
		g.P("ctx = ", internal.TransportxPackage.Ident("NameInjector"), "(ctx, c.transportName)")
		g.P("endpoint, err := c.endpoints.", endpoint.Name(), "(ctx)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("rep, err := endpoint(ctx, request)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
		g.P("}")
		g.P()
	}
}
