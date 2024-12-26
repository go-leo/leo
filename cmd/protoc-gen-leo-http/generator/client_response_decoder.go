package generator

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

type ClientResponseDecoderGenerator struct{}

func (f *ClientResponseDecoderGenerator) GenerateClientResponseDecoder(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.HttpClientResponseDecoderName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("DecodeResponseFunc"))
	}
	g.P("}")
	g.P()
	return nil
}

func (f *ClientResponseDecoderGenerator) GenerateClientResponseDecoderImplements(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.HttpClientResponseDecoderName()), " struct {}")
	for _, endpoint := range service.Endpoints {
		g.P("func (", service.Unexported(service.HttpClientResponseDecoderName()), ")", endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("DecodeResponseFunc"), " {")
		httpRule := endpoint.HttpRule()
		g.P("return func ", "(ctx context.Context, r *", internal.HttpPackage.Ident("Response"), ") (any, error) {")
		g.P("if ", internal.HttpxTransportxPackage.Ident("IsErrorResponse"), "(r) {")
		g.P("return nil, ", internal.HttpxTransportxPackage.Ident("ErrorDecoder"), "(ctx, r)")
		g.P("}")
		g.P("resp := &", endpoint.Output().GoIdent, "{}")
		bodyParameter := httpRule.ResponseBody()
		switch bodyParameter {
		case "", "*":
			srcValue := []any{"resp"}
			message := endpoint.Output()
			switch message.Desc.FullName() {
			case "google.api.HttpBody":
				f.PrintGoogleApiHttpBodyDecodeBlock(g, srcValue)
			default:
				f.PrintJsonDecodeBlock(g, srcValue)
			}
		default:
			bodyField := internal.FindField(bodyParameter, endpoint.Output())
			if bodyField == nil {
				return fmt.Errorf("%s, failed to find body response field %s", endpoint.FullName(), bodyParameter)
			}
			srcValue := []any{"resp.", bodyField.GoName}
			if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
				g.P(append(append([]any{}, srcValue...), " = &", bodyField.Message.GoIdent, "{}")...)
				f.PrintGoogleApiHttpBodyDecodeBlock(g, srcValue)
			} else {
				_, pointer := internal.FieldGoType(g, bodyField)
				if !pointer {
					srcValue = append([]any{"&"}, srcValue...)
				}
				f.PrintJsonDecodeBlock(g, srcValue)
			}
		}
		g.P("return resp, nil")
		g.P("}")
		g.P("}")
	}
	g.P()
	return nil
}

func (f *ClientResponseDecoderGenerator) PrintGoogleApiHttpBodyDecodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{}, srcValue...), ".ContentType = r.Header.Get(", strconv.Quote(internal.ContentTypeKey), ")")...)
	g.P("body, err := ", internal.IOPackage.Ident("ReadAll"), "(r.Body)")
	g.P("if err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P(append(append([]any{}, srcValue...), ".Data = body")...)

}

func (f *ClientResponseDecoderGenerator) PrintJsonDecodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"if err := ", internal.JsonxPackage.Ident("NewDecoder"), "(r.Body).Decode("}, srcValue...), "); err != nil {")...)
	g.P("return nil, err")
	g.P("}")
}
