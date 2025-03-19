package http

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ResponseDecoderGenerator struct {
	service *internal.Service
	g       *protogen.GeneratedFile
}

func (f *ResponseDecoderGenerator) GenerateClientResponseDecoder() {
	f.g.P("type ", f.service.HttpClientResponseDecoderName(), " interface {")
	for _, endpoint := range f.service.Endpoints {
		f.g.P(endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("DecodeResponseFunc"))
	}
	f.g.P("}")
	f.g.P()
}

func (f *ResponseDecoderGenerator) GenerateClientResponseDecoderImplements() error {
	f.g.P("type ", f.service.Unexported(f.service.HttpClientResponseDecoderName()), " struct {")
	f.g.P("unmarshalOptions ", internal.ProtoJsonUnmarshalOptionsIdent)
	f.g.P("}")
	for _, endpoint := range f.service.Endpoints {
		f.g.P("func (decoder ", f.service.Unexported(f.service.HttpClientResponseDecoderName()), ")", endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("DecodeResponseFunc"), " {")
		f.g.P("return func ", "(ctx context.Context, r *", internal.Response, ") (any, error) {")
		f.g.P("if r.StatusCode != ", internal.StatusOK, " {")
		f.g.P("return nil, ", internal.DecodeErrorFromResponse, "(ctx, r)")
		f.g.P("}")
		f.g.P("resp := &", endpoint.Output().GoIdent, "{}")
		bodyParameter := endpoint.ResponseBody()
		switch bodyParameter {
		case "", "*":
			message := endpoint.Output()
			switch message.Desc.FullName() {
			case "google.api.HttpBody":
				srcValue := []any{"resp"}
				f.DecodeHttpBodyFromResponse(srcValue)
			case "google.rpc.HttpResponse":
				srcValue := []any{"resp"}
				f.DecodeHttpResponseFromResponse(srcValue)
			default:
				srcValue := []any{"resp"}
				f.PrintDecodeMessageFromResponse(srcValue)
			}
		default:
			bodyField := internal.FindField(bodyParameter, endpoint.Output())
			if bodyField == nil {
				return fmt.Errorf("%s, failed to find body response field %s", endpoint.FullName(), bodyParameter)
			}
			srcValue := []any{"resp.", bodyField.GoName}
			f.g.P(append(append(append([]any{"if "}, srcValue...), " == nil {"))...)
			f.g.P(append(srcValue, " = &", bodyField.Message.GoIdent, "{}")...)
			f.g.P("}")
			switch bodyField.Desc.Kind() {
			case protoreflect.MessageKind:
				switch bodyField.Message.Desc.FullName() {
				case "google.api.HttpBody":
					f.DecodeHttpBodyFromResponse(srcValue)
				default:
					f.PrintDecodeMessageFromResponse(srcValue)
				}
			}
		}
		f.g.P("return resp, nil")
		f.g.P("}")
		f.g.P("}")
	}
	f.g.P()
	return nil
}

func (f *ResponseDecoderGenerator) PrintDecodeMessageFromResponse(srcValue []any) {
	f.g.P(append(append([]any{"if err := ", internal.DecodeMessageFromResponse, "(ctx, r, "}, srcValue...), " ,decoder.unmarshalOptions); err != nil {")...)
	f.g.P("return nil, err")
	f.g.P("}")
}

func (f *ResponseDecoderGenerator) DecodeHttpBodyFromResponse(srcValue []any) {
	f.g.P(append(append([]any{"if err := ", internal.DecodeHttpBodyFromResponse, "(ctx, r, "}, srcValue...), "); err != nil {")...)
	f.g.P("return nil, err")
	f.g.P("}")
}

func (f *ResponseDecoderGenerator) DecodeHttpResponseFromResponse(srcValue []any) {
	f.g.P(append(append([]any{"if err := ", internal.DecodeHttpResponseFromResponse, "(ctx, r, "}, srcValue...), "); err != nil {")...)
	f.g.P("return nil, err")
	f.g.P("}")
}
