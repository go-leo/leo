package http

import (
	"github.com/go-leo/leo/v3/cmd/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"strconv"
)

type ServerTransportsGenerator struct{}

func (f *ServerTransportsGenerator) GenerateTransports(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.HttpServerTransportsName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "()", internal.HttpPackage.Ident("Handler"))
	}
	g.P("}")
	g.P()
	return nil
}

func (f *ServerTransportsGenerator) GenerateTransportsImplements(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.HttpServerTransportsName()), " struct {")
	g.P("endpoints ", service.ServerEndpointsName())
	g.P("requestDecoder ", service.HttpServerRequestDecoderName())
	g.P("responseEncoder ", service.HttpServerResponseEncoderName())
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (t *", service.Unexported(service.HttpServerTransportsName()), ")", endpoint.Name(), "()", internal.HttpPackage.Ident("Handler"), " {")
		g.P("return ", internal.HttpTransportPackage.Ident("NewServer"), "(")
		g.P("t.endpoints.", endpoint.Name(), "(", internal.ContextPackage.Ident("TODO"), "()), ")
		g.P("t.requestDecoder.", endpoint.Name(), "(),")
		g.P("t.responseEncoder.", endpoint.Name(), "(),")
		g.P(internal.HttpTransportPackage.Ident("ServerBefore"), "(", internal.HttpxTransportxPackage.Ident("EndpointInjector"), "(", strconv.Quote(endpoint.FullName()), ")),")
		g.P(internal.HttpTransportPackage.Ident("ServerBefore"), "(", internal.HttpxTransportxPackage.Ident("ServerTransportInjector"), "),")
		g.P(internal.HttpTransportPackage.Ident("ServerBefore"), "(", internal.HttpxTransportxPackage.Ident("IncomingMetadataInjector"), "),")
		g.P(internal.HttpTransportPackage.Ident("ServerBefore"), "(", internal.HttpxTransportxPackage.Ident("IncomingTimeLimitInjector"), "),")
		g.P(internal.HttpTransportPackage.Ident("ServerBefore"), "(", internal.HttpxTransportxPackage.Ident("IncomingStainInjector"), "),")
		g.P(internal.HttpTransportPackage.Ident("ServerFinalizer"), "(", internal.HttpxTransportxPackage.Ident("CancelInvoker"), "),")
		g.P(")")
		g.P("}")
		g.P()
	}
	return nil
}
