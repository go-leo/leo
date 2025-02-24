package http

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

type ServerResponseEncoderGenerator struct{}

func (f *ServerResponseEncoderGenerator) GenerateServerResponseEncoder(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.HttpServerResponseEncoderName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("EncodeResponseFunc"))
	}
	g.P("}")
	g.P()
	return nil
}

func (f *ServerResponseEncoderGenerator) GenerateServerResponseEncoderImplements(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.HttpServerResponseEncoderName()), " struct {")
	g.P("marshalOptions ", internal.ProtoJsonMarshalOptionsIdent)
	g.P("unmarshalOptions ", internal.ProtoJsonUnmarshalOptionsIdent)
	g.P("responseTransformer ", internal.ResponseTransformerIdent)
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (encoder ", service.Unexported(service.HttpServerResponseEncoderName()), ")", endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("EncodeResponseFunc"), "{")
		httpRule := endpoint.HttpRule()
		g.P("return func ", "(ctx ", internal.ContextPackage.Ident("Context"), ", w ", internal.HttpPackage.Ident("ResponseWriter"), ", obj any) error {")
		g.P("resp := obj.(*", endpoint.Output().GoIdent, ")")
		bodyParameter := httpRule.ResponseBody()
		switch bodyParameter {
		case "", "*":
			message := endpoint.Output()
			switch message.Desc.FullName() {
			case "google.api.HttpBody":
				srcValue := []any{"resp"}
				f.PrintHttpBodyEncodeBlock(g, srcValue)
			case "google.rpc.HttpResponse":
				srcValue := []any{"resp"}
				f.PrintHttpResponseEncodeBlock(g, srcValue)
			default:
				srcValue := []any{"encoder.responseTransformer(ctx, resp)"}
				f.PrintResponseEncodeBlock(g, srcValue)
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
					f.PrintHttpBodyEncodeBlock(g, srcValue)
				default:
					f.PrintResponseEncodeBlock(g, srcValue)
				}
			}
		}
		g.P("}")
		g.P("}")
	}
	g.P()
	return nil
}

func (f *ServerResponseEncoderGenerator) PrintGoogleApiHttpBodyEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"w.Header().Set(", strconv.Quote("Content-Type"), ", "}, srcValue...), ".GetContentType())")...)
	g.P(append(append([]any{"for _, src := range "}, srcValue...), ".GetExtensions() {")...)
	g.P("dst, err := ", internal.AnypbPackage.Ident("UnmarshalNew"), "(src, ", internal.ProtoPackage.Ident("UnmarshalOptions"), "{})")
	g.P("if err != nil {")
	g.P("return err")
	g.P("}")
	g.P("metadata, ok := dst.(*", internal.StructpbPackage.Ident("Struct"), ")")
	g.P("if !ok {")
	g.P("continue")
	g.P("}")
	g.P("for key, value := range metadata.GetFields() {")
	g.P("w.Header().Add(key, string(", internal.ErrorxPackage.Ident("Ignore"), "(", internal.JsonxPackage.Ident("Marshal"), "(value))))")
	g.P("}")
	g.P("}")
	g.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
	g.P(append(append([]any{"if ", "_, err := w.Write("}, srcValue...), ".GetData())", "; err != nil {")...)
	g.P("return err")
	g.P("}")
}

//func (f *ServerResponseEncoderGenerator) PrintJsonEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
//	g.P("w.Header().Set(", strconv.Quote("Content-Type"), ", ", strconv.Quote(internal.JsonContentType), ")")
//	g.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
//	g.P(append(append([]any{"if err := ", internal.JsonxPackage.Ident("NewEncoder"), "(w).Encode("}, srcValue...), "); err != nil {")...)
//	g.P("return err")
//	g.P("}")
//}

func (f *ServerResponseEncoderGenerator) PrintHttpBodyEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"return ", internal.HttpBodyEncoderIdent, "(ctx, w, "}, srcValue...), ")")...)
}

func (f *ServerResponseEncoderGenerator) PrintHttpResponseEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"return ", internal.HttpResponseEncoderIdent, "(ctx, w, "}, srcValue...), ")")...)
}

func (f *ServerResponseEncoderGenerator) PrintResponseEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"return ", internal.ResponseEncoderIdent, "(ctx, w, "}, srcValue...), ", encoder.marshalOptions)")...)
}
