package http

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

type ResponseEncoderGenerator struct {
	service *internal.Service
	g       *protogen.GeneratedFile
}

func (f *ResponseEncoderGenerator) GenerateResponseEncoder() error {
	f.g.P("type ", f.service.HttpServerResponseEncoderName(), " interface {")
	for _, endpoint := range f.service.Endpoints {
		f.g.P(endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("EncodeResponseFunc"))
	}
	f.g.P("}")
	f.g.P()
	return nil
}

func (f *ResponseEncoderGenerator) GenerateServerResponseEncoderImplements() error {
	f.g.P("type ", f.service.Unexported(f.service.HttpServerResponseEncoderName()), " struct {")
	f.g.P("marshalOptions ", internal.ProtoJsonMarshalOptionsIdent)
	f.g.P("unmarshalOptions ", internal.ProtoJsonUnmarshalOptionsIdent)
	f.g.P("responseTransformer ", internal.ResponseTransformer)
	f.g.P("}")
	f.g.P()
	for _, endpoint := range f.service.Endpoints {
		f.g.P("func (encoder ", f.service.Unexported(f.service.HttpServerResponseEncoderName()), ")", endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("EncodeResponseFunc"), "{")
		httpRule := endpoint.HttpRule()
		f.g.P("return func ", "(ctx ", internal.ContextPackage.Ident("Context"), ", w ", internal.HttpPackage.Ident("ResponseWriter"), ", obj any) error {")
		f.g.P("resp := obj.(*", endpoint.Output().GoIdent, ")")
		bodyParameter := httpRule.ResponseBody()
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
				srcValue := []any{"encoder.responseTransformer(ctx, resp)"}
				f.PrintEncodeResponseToResponse(srcValue)
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
					f.PrintEncodeResponseToResponse(srcValue)
				}
			}
		}
		f.g.P("}")
		f.g.P("}")
	}
	f.g.P()
	return nil
}

func (f *ResponseEncoderGenerator) PrintGoogleApiHttpBodyEncodeBlock(srcValue []any) {
	f.g.P(append(append([]any{"w.Header().Set(", strconv.Quote("Content-Type"), ", "}, srcValue...), ".GetContentType())")...)
	f.g.P(append(append([]any{"for _, src := range "}, srcValue...), ".GetExtensions() {")...)
	f.g.P("dst, err := ", internal.AnypbPackage.Ident("UnmarshalNew"), "(src, ", internal.ProtoPackage.Ident("UnmarshalOptions"), "{})")
	f.g.P("if err != nil {")
	f.g.P("return err")
	f.g.P("}")
	f.g.P("metadata, ok := dst.(*", internal.StructpbPackage.Ident("Struct"), ")")
	f.g.P("if !ok {")
	f.g.P("continue")
	f.g.P("}")
	f.g.P("for key, value := range metadata.GetFields() {")
	f.g.P("w.Header().Add(key, string(", internal.ErrorxPackage.Ident("Ignore"), "(", internal.JsonxPackage.Ident("Marshal"), "(value))))")
	f.g.P("}")
	f.g.P("}")
	f.g.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
	f.g.P(append(append([]any{"if ", "_, err := w.Write("}, srcValue...), ".GetData())", "; err != nil {")...)
	f.g.P("return err")
	f.g.P("}")
}

func (f *ResponseEncoderGenerator) PrintEncodeResponseToResponse(srcValue []any) {
	f.g.P(append(append([]any{"return ", internal.EncodeResponseToResponse, "(ctx, w, "}, srcValue...), ", encoder.marshalOptions)")...)
}

func (f *ResponseEncoderGenerator) PrintEncodeHttpBodyToResponse(srcValue []any) {
	f.g.P(append(append([]any{"return ", internal.EncodeHttpBodyToResponse, "(ctx, w, "}, srcValue...), ")")...)
}

func (f *ResponseEncoderGenerator) PrintEncodeHttpResponseToResponse(srcValue []any) {
	f.g.P(append(append([]any{"return ", internal.EncodeHttpResponseToResponse, "(ctx, w, "}, srcValue...), ")")...)
}
