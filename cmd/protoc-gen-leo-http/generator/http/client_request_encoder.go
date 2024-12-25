package http

import (
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

type ClientRequestEncoderGenerator struct{}

func (f *ClientRequestEncoderGenerator) GenerateClientRequestEncoder(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.HttpClientRequestEncoderName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("CreateRequestFunc"))
	}
	g.P("}")
	g.P()
	return nil
}
