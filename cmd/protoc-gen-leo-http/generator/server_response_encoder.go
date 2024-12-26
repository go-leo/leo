package generator

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/internal"
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
	g.P("type ", service.Unexported(service.HttpServerResponseEncoderName()), " struct {}")
	for _, endpoint := range service.Endpoints {
		g.P("func (", service.Unexported(service.HttpServerResponseEncoderName()), ")", endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("EncodeResponseFunc"), "{")
		httpRule := endpoint.HttpRule()
		g.P("return func ", "(ctx ", internal.ContextPackage.Ident("Context"), ", w ", internal.HttpPackage.Ident("ResponseWriter"), ", obj any) error {")
		g.P("resp := obj.(*", endpoint.Output().GoIdent, ")")
		bodyParameter := httpRule.ResponseBody()
		switch bodyParameter {
		case "", "*":
			srcValue := []any{"resp"}
			message := endpoint.Output()
			switch message.Desc.FullName() {
			case "google.api.HttpBody":
				f.PrintGoogleApiHttpBodyEncodeBlock(g, srcValue)
			default:
				f.PrintJsonEncodeBlock(g, srcValue)
			}
		default:
			bodyField := internal.FindField(bodyParameter, endpoint.Output())
			if bodyField == nil {
				return fmt.Errorf("%s, failed to find body response field %s", endpoint.FullName(), bodyParameter)
			}
			srcValue := []any{"resp.Get", bodyField.GoName, "()"}
			if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
				f.PrintGoogleApiHttpBodyEncodeBlock(g, srcValue)
			} else {
				f.PrintJsonEncodeBlock(g, srcValue)
			}
		}
		g.P("return nil")
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
	g.P("return ", internal.StatusxPackage.Ident("ErrInternal"), ".With(", internal.StatusxPackage.Ident("Wrap"), "(err))")
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
	g.P("return ", internal.StatusxPackage.Ident("ErrInternal"), ".With(", internal.StatusxPackage.Ident("Wrap"), "(err))")
	g.P("}")
}

func (f *ServerResponseEncoderGenerator) PrintJsonEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P("w.Header().Set(", strconv.Quote("Content-Type"), ", ", strconv.Quote(internal.JsonContentType), ")")
	g.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
	g.P(append(append([]any{"if err := ", internal.JsonxPackage.Ident("NewEncoder"), "(w).Encode("}, srcValue...), "); err != nil {")...)
	g.P("return ", internal.StatusxPackage.Ident("ErrInternal"), ".With(", internal.StatusxPackage.Ident("Wrap"), "(err))")
	g.P("}")
}
