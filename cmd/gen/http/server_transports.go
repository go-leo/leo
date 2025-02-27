package http

import (
	"github.com/go-leo/leo/v3/cmd/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"strconv"
)

type ServerTransportsGenerator struct {
	service *internal.Service
	g       *protogen.GeneratedFile
}

func (f *ServerTransportsGenerator) GenerateTransports() {
	f.g.P("type ", f.service.HttpServerTransportsName(), " interface {")
	for _, endpoint := range f.service.Endpoints {
		f.g.P(endpoint.Name(), "()", internal.Handler)
	}
	f.g.P("}")
	f.g.P()
}

func (f *ServerTransportsGenerator) GenerateTransportsImplements() {
	f.g.P("type ", f.service.Unexported(f.service.HttpServerTransportsName()), " struct {")
	f.g.P("endpoints ", f.service.ServerEndpointsName())
	f.g.P("requestDecoder ", f.service.HttpServerRequestDecoderName())
	f.g.P("responseEncoder ", f.service.HttpServerResponseEncoderName())
	f.g.P("}")
	f.g.P()
	for _, endpoint := range f.service.Endpoints {
		f.g.P("func (t *", f.service.Unexported(f.service.HttpServerTransportsName()), ")", endpoint.Name(), "()", internal.Handler, " {")
		f.g.P("return ", internal.HttpTransportPackage.Ident("NewServer"), "(")
		f.g.P("t.endpoints.", endpoint.Name(), "(", internal.ContextPackage.Ident("TODO"), "()), ")
		f.g.P("t.requestDecoder.", endpoint.Name(), "(),")
		f.g.P("t.responseEncoder.", endpoint.Name(), "(),")
		f.g.P(internal.HttpTransportPackage.Ident("ServerBefore"), "(", internal.HttpxTransportxPackage.Ident("EndpointInjector"), "(", strconv.Quote(endpoint.FullName()), ")),")
		f.g.P(internal.HttpTransportPackage.Ident("ServerBefore"), "(", internal.HttpxTransportxPackage.Ident("ServerTransportInjector"), "),")
		f.g.P(internal.HttpTransportPackage.Ident("ServerBefore"), "(", internal.HttpxTransportxPackage.Ident("IncomingMetadataInjector"), "),")
		f.g.P(internal.HttpTransportPackage.Ident("ServerBefore"), "(", internal.HttpxTransportxPackage.Ident("IncomingTimeLimitInjector"), "),")
		f.g.P(internal.HttpTransportPackage.Ident("ServerBefore"), "(", internal.HttpxTransportxPackage.Ident("IncomingStainInjector"), "),")
		f.g.P(internal.HttpTransportPackage.Ident("ServerFinalizer"), "(", internal.HttpxTransportxPackage.Ident("CancelInvoker"), "),")
		f.g.P(")")
		f.g.P("}")
		f.g.P()
	}
}
