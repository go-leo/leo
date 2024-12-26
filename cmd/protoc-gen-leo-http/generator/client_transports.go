package generator

import (
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

type ClientTransportsGenerator struct{}

func (f *ClientTransportsGenerator) GenerateTransports(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.HttpClientTransportsName()), " struct {")
	g.P("clientOptions []", internal.HttpTransportPackage.Ident("ClientOption"))
	g.P("middlewares []", internal.EndpointPackage.Ident("Middleware"))
	g.P("requestEncoder ", service.HttpClientRequestEncoderName())
	g.P("responseDecoder ", service.HttpClientResponseDecoderName())
	g.P("}")
	g.P()

	for _, endpoint := range service.Endpoints {
		g.P("func (t *", service.Unexported(service.HttpClientTransportsName()), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", instance string) (", internal.EndpointPackage.Ident("Endpoint"), ", ", internal.IOPackage.Ident("Closer"), ", error) {")
		g.P("opts := []", internal.HttpTransportPackage.Ident("ClientOption"), "{")
		g.P(internal.HttpTransportPackage.Ident("ClientBefore"), "(", internal.HttpxTransportxPackage.Ident("OutgoingMetadataInjector"), "),")
		g.P(internal.HttpTransportPackage.Ident("ClientBefore"), "(", internal.HttpxTransportxPackage.Ident("OutgoingTimeLimitInjector"), "),")
		g.P(internal.HttpTransportPackage.Ident("ClientBefore"), "(", internal.HttpxTransportxPackage.Ident("OutgoingStainInjector"), "),")
		g.P("}")
		g.P("opts = append(opts, t.clientOptions...)")
		g.P("client := ", internal.HttpTransportPackage.Ident("NewExplicitClient"), "(")
		g.P("t.requestEncoder.", endpoint.Name(), "(instance),")
		g.P("t.responseDecoder.", endpoint.Name(), "(),")
		g.P("opts...,")
		g.P(")")
		g.P("return ", internal.EndpointxPackage.Ident("Chain"), "(client.Endpoint(), t.middlewares...), nil, nil")
		g.P("}")
		g.P()
	}

	g.P("func new", service.HttpClientTransportsName(), "(scheme string, clientOptions []", internal.HttpTransportPackage.Ident("ClientOption"), ", middlewares []", internal.EndpointPackage.Ident("Middleware"), ") ", service.ClientTransportsName(), " {")
	g.P("return &", service.Unexported(service.HttpClientTransportsName()), "{")
	g.P("clientOptions: clientOptions,")
	g.P("middlewares:   middlewares,")
	g.P("requestEncoder:  ", service.Unexported(service.HttpClientRequestEncoderName()), "{")
	g.P("scheme:        scheme,")
	g.P("router:        append", service.HttpRoutesName(), "(", internal.MuxPackage.Ident("NewRouter"), "()),")
	g.P("},")
	g.P("responseDecoder: ", service.Unexported(service.HttpClientResponseDecoderName()), "{},")
	g.P("}")
	g.P("}")
	g.P()
	return nil
}
