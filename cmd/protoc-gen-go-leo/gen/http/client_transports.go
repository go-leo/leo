package http

import (
	"github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

type ClientTransportsGenerator struct {
	service *internal.Service
	g       *protogen.GeneratedFile
}

func (f *ClientTransportsGenerator) GenerateTransports() {
	f.g.P("type ", f.service.Unexported(f.service.HttpClientTransportsName()), " struct {")
	f.g.P("clientOptions []", internal.HttpTransportPackage.Ident("ClientOption"))
	f.g.P("middlewares []", internal.EndpointPackage.Ident("Middleware"))
	f.g.P("requestEncoder ", f.service.HttpClientRequestEncoderName())
	f.g.P("responseDecoder ", f.service.HttpClientResponseDecoderName())
	f.g.P("}")
	f.g.P()

	for _, endpoint := range f.service.Endpoints {
		f.g.P("func (t *", f.service.Unexported(f.service.HttpClientTransportsName()), ") ", endpoint.Name(), "(ctx ", internal.Context, ", instance string) (", internal.EndpointPackage.Ident("Endpoint"), ", ", internal.IOPackage.Ident("Closer"), ", error) {")
		f.g.P("opts := []", internal.HttpTransportPackage.Ident("ClientOption"), "{")
		f.g.P(internal.HttpTransportPackage.Ident("ClientBefore"), "(", internal.MetadataxHttpOutgoingInjector, "),")
		f.g.P(internal.HttpTransportPackage.Ident("ClientBefore"), "(", internal.TimeoutxOutgoingInjector, "),")
		f.g.P(internal.HttpTransportPackage.Ident("ClientBefore"), "(", internal.StainxHttpOutgoingInjector, "),")
		f.g.P("}")
		f.g.P("opts = append(opts, t.clientOptions...)")
		f.g.P("client := ", internal.HttpTransportPackage.Ident("NewExplicitClient"), "(")
		f.g.P("t.requestEncoder.", endpoint.Name(), "(instance),")
		f.g.P("t.responseDecoder.", endpoint.Name(), "(),")
		f.g.P("opts...,")
		f.g.P(")")
		f.g.P("return ", internal.EndpointxChain, "(client.Endpoint(), t.middlewares...), nil, nil")
		f.g.P("}")
		f.g.P()
	}
}
