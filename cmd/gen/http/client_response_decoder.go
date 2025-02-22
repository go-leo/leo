package http

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

type ClientResponseDecoderGenerator struct {
	service *internal.Service
	g       *protogen.GeneratedFile
}

func (f *ClientResponseDecoderGenerator) GenerateClientResponseDecoder() error {
	f.g.P("type ", f.service.HttpClientResponseDecoderName(), " interface {")
	for _, endpoint := range f.service.Endpoints {
		f.g.P(endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("DecodeResponseFunc"))
	}
	f.g.P("}")
	f.g.P()
	return nil
}

func (f *ClientResponseDecoderGenerator) GenerateClientResponseDecoderImplements() error {
	f.g.P("type ", f.service.Unexported(f.service.HttpClientResponseDecoderName()), " struct {}")
	for _, endpoint := range f.service.Endpoints {
		f.g.P("func (", f.service.Unexported(f.service.HttpClientResponseDecoderName()), ")", endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("DecodeResponseFunc"), " {")
		httpRule := endpoint.HttpRule()
		f.g.P("return func ", "(ctx context.Context, r *", internal.HttpPackage.Ident("Response"), ") (any, error) {")
		f.g.P("if ", internal.HttpxTransportxPackage.Ident("IsErrorResponse"), "(r) {")
		f.g.P("return nil, ", internal.HttpxTransportxPackage.Ident("ErrorDecoder"), "(ctx, r)")
		f.g.P("}")
		f.g.P("resp := &", endpoint.Output().GoIdent, "{}")
		bodyParameter := httpRule.ResponseBody()
		switch bodyParameter {
		case "", "*":
			srcValue := []any{"resp"}
			message := endpoint.Output()
			switch message.Desc.FullName() {
			case "google.api.HttpBody":
				f.PrintGoogleApiHttpBodyDecodeBlock(srcValue)
			default:
				f.PrintJsonDecodeBlock(srcValue)
			}
		default:
			bodyField := internal.FindField(bodyParameter, endpoint.Output())
			if bodyField == nil {
				return fmt.Errorf("%s, failed to find body response field %s", endpoint.FullName(), bodyParameter)
			}
			srcValue := []any{"resp.", bodyField.GoName}
			if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
				f.g.P(append(append([]any{}, srcValue...), " = &", bodyField.Message.GoIdent, "{}")...)
				f.PrintGoogleApiHttpBodyDecodeBlock(srcValue)
			} else {
				_, pointer := internal.FieldGoType(f.g, bodyField)
				if !pointer {
					srcValue = append([]any{"&"}, srcValue...)
				}
				f.PrintJsonDecodeBlock(srcValue)
			}
		}
		f.g.P("return resp, nil")
		f.g.P("}")
		f.g.P("}")
	}
	f.g.P()
	return nil
}

func (f *ClientResponseDecoderGenerator) PrintGoogleApiHttpBodyDecodeBlock(srcValue []any) {
	f.g.P(append(append([]any{}, srcValue...), ".ContentType = r.Header.Get(", strconv.Quote(internal.ContentTypeKey), ")")...)
	f.g.P("body, err := ", internal.IOPackage.Ident("ReadAll"), "(r.Body)")
	f.g.P("if err != nil {")
	f.g.P("return nil, err")
	f.g.P("}")
	f.g.P(append(append([]any{}, srcValue...), ".Data = body")...)

}

func (f *ClientResponseDecoderGenerator) PrintJsonDecodeBlock(srcValue []any) {
	f.g.P(append(append([]any{"if err := ", internal.JsonxPackage.Ident("NewDecoder"), "(r.Body).Decode("}, srcValue...), "); err != nil {")...)
	f.g.P("return nil, err")
	f.g.P("}")
}
