package http

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ResponseEncoderGenerator struct {
	service *internal.Service
	g       *protogen.GeneratedFile
}

func (f *ResponseEncoderGenerator) GenerateResponseEncoder() {
	f.g.P("type ", f.service.HttpServerResponseEncoderName(), " interface {")
	for _, endpoint := range f.service.Endpoints {
		f.g.P(endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("EncodeResponseFunc"))
	}
	f.g.P("}")
	f.g.P()
}

func (f *ResponseEncoderGenerator) GenerateServerResponseEncoderImplements() error {
	f.g.P("type ", f.service.Unexported(f.service.HttpServerResponseEncoderName()), " struct {")
	f.g.P("marshalOptions ", internal.ProtoJsonMarshalOptionsIdent)
	f.g.P("}")
	f.g.P()
	for _, endpoint := range f.service.Endpoints {
		f.g.P("func (encoder ", f.service.Unexported(f.service.HttpServerResponseEncoderName()), ")", endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("EncodeResponseFunc"), "{")
		f.g.P("return func ", "(ctx ", internal.Context, ", w ", internal.ResponseWriter, ", obj any) error {")
		f.g.P("resp := obj.(*", endpoint.Output().GoIdent, ")")
		bodyParameter := endpoint.ResponseBody()
		switch bodyParameter {
		case "", "*":
			message := endpoint.Output()
			switch message.Desc.FullName() {
			case "google.api.HttpBody":
				srcValue := []any{"resp"}
				f.PrintEncodeHttpBodyToResponse(srcValue)
			case "google.rpc.HttpResponse":
				srcValue := []any{"resp"}
				f.PrintEncodeHttpResponseToResponse(srcValue)
			default:
				srcValue := []any{"resp"}
				f.PrintEncodeMessageToResponse(srcValue)
			}
		default:
			bodyField := internal.FindField(bodyParameter, endpoint.Output())
			if bodyField == nil {
				return fmt.Errorf("%s, failed to find body response field %s", endpoint.FullName(), bodyParameter)
			}
			srcValue := []any{"resp.Get", bodyField.GoName, "()"}
			switch bodyField.Desc.Kind() {
			case protoreflect.MessageKind:
				switch bodyField.Message.Desc.FullName() {
				case "google.api.HttpBody":
					f.PrintEncodeHttpBodyToResponse(srcValue)
				default:
					f.PrintEncodeMessageToResponse(srcValue)
				}
			}
		}
		f.g.P("}")
		f.g.P("}")
	}
	f.g.P()
	return nil
}

func (f *ResponseEncoderGenerator) PrintEncodeMessageToResponse(srcValue []any) {
	f.g.P(append(append([]any{"return ", internal.EncodeMessageToResponse, "(ctx, w, "}, srcValue...), ", encoder.marshalOptions)")...)
}

func (f *ResponseEncoderGenerator) PrintEncodeHttpBodyToResponse(srcValue []any) {
	f.g.P(append(append([]any{"return ", internal.EncodeHttpBodyToResponse, "(ctx, w, "}, srcValue...), ")")...)
}

func (f *ResponseEncoderGenerator) PrintEncodeHttpResponseToResponse(srcValue []any) {
	f.g.P(append(append([]any{"return ", internal.EncodeHttpResponseToResponse, "(ctx, w, "}, srcValue...), ")")...)
}
