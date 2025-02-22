package http

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

type ClientRequestEncoderGenerator struct{}

func (f *ClientRequestEncoderGenerator) GenerateClientRequestEncoder(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.HttpClientRequestEncoderName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "(instance string) ", internal.HttpTransportPackage.Ident("CreateRequestFunc"))
	}
	g.P("}")
	g.P()
	return nil
}

func (f *ClientRequestEncoderGenerator) GenerateClientRequestEncoderImplements(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.Unexported(service.HttpClientRequestEncoderName()), " struct {")
	g.P("router *", internal.MuxPackage.Ident("Router"))
	g.P("scheme string")
	g.P("}")
	for _, endpoint := range service.Endpoints {
		g.P("func (e ", service.Unexported(service.HttpClientRequestEncoderName()), ")", endpoint.Name(), "(instance string) ", internal.HttpTransportPackage.Ident("CreateRequestFunc"), " {")
		httpRule := endpoint.HttpRule()
		g.P("return func(ctx context.Context, obj any) (*", internal.HttpPackage.Ident("Request"), ", error) {")
		g.P("if obj == nil {")
		g.P("return nil, ", internal.ErrorsPackage.Ident("New"), "(", strconv.Quote("request is nil"), ")")
		g.P("}")
		g.P("req, ok := obj.(*", endpoint.InputGoIdent(), ")")
		g.P("if !ok {")
		g.P("return nil, ", internal.FmtPackage.Ident("Errorf"), "(", strconv.Quote("invalid request type, %T"), ", obj)")
		g.P("}")
		g.P("_ = req")

		bodyMessage, bodyField, namedPathFields, pathFields, queryFields, err := endpoint.ParseParameters()
		if err != nil {
			return err
		}

		g.P("var body ", internal.IOPackage.Ident("Reader"))
		if bodyMessage != nil {
			switch bodyMessage.Desc.FullName() {
			case "google.api.HttpBody":
				f.PrintReaderBlock(g, internal.BytesPackage, []any{"body"}, []any{"req.GetData()"})
				g.P("contentType := req.GetContentType()")
			default:
				g.P("var bodyBuf ", internal.BytesPackage.Ident("Buffer"))
				encoder := internal.JsonxPackage.Ident("NewEncoder")
				f.PrintEncodeBlock(g, encoder, []any{"&bodyBuf"}, []any{"req"})
				g.P("body = &bodyBuf")
				g.P("contentType := ", strconv.Quote(internal.JsonContentType))
			}
		} else if bodyField != nil {
			if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
				f.PrintReaderBlock(g, internal.BytesPackage, []any{"body"}, []any{"req.Get", bodyField.GoName, "()", ".GetData()"})
				g.P("contentType := req.Get", bodyField.GoName, "()", ".GetContentType()")
			} else {
				g.P("var bodyBuf ", internal.BytesPackage.Ident("Buffer"))
				encoder := internal.JsonxPackage.Ident("NewEncoder")
				tgtValue := []any{"&bodyBuf"}
				srcValue := []any{"req.Get", bodyField.GoName, "()"}
				f.PrintEncodeBlock(g, encoder, tgtValue, srcValue)
				g.P("body = &bodyBuf")
				g.P("contentType := ", strconv.Quote(internal.JsonContentType))
			}
		}

		g.P("var pairs []string")
		if len(namedPathFields) > 0 {
			f.PrintNamedPathField(g, namedPathFields, httpRule)
		}

		if len(pathFields) > 0 {
			f.PrintPathField(g, pathFields)
		}

		g.P("path, err := e.router.Get(", strconv.Quote(endpoint.FullName()), ").URLPath(pairs...)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")

		g.P("queries := ", internal.UrlPackage.Ident("Values"), "{}")
		if len(queryFields) > 0 {
			f.PrintQueryField(g, queryFields)
		}

		g.P("target := &", internal.UrlPackage.Ident("URL"), "{")
		g.P("Scheme:   e.scheme,")
		g.P("Host:     instance,")
		g.P("Path:     path.Path,")
		g.P("RawQuery: queries.Encode(),")
		g.P("}")

		g.P("r, err := ", internal.HttpPackage.Ident("NewRequestWithContext"), "(ctx, ", strconv.Quote(httpRule.Method()), ", target.String(), body)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		if bodyMessage != nil || bodyField != nil {
			g.P("r.Header.Set(", strconv.Quote(internal.ContentTypeKey), ", contentType)")
		}
		g.P("return r, nil")
		g.P("}")
		g.P("}")
	}
	g.P()
	return nil
}

func (f *ClientRequestEncoderGenerator) PrintReaderBlock(g *protogen.GeneratedFile, readerPkg protogen.GoImportPath, tgtValue []any, srcValue []any) {
	g.P(append(append(append(append([]any{}, tgtValue...), []any{" = ", readerPkg.Ident("NewReader"), "("}...), srcValue...), ")")...)
}

func (f *ClientRequestEncoderGenerator) PrintEncodeBlock(g *protogen.GeneratedFile, encoder protogen.GoIdent, tgtValue []any, srcValue []any) {
	g.P(append(append(append(append([]any{"if err := ", encoder, "("}, tgtValue...), []any{").Encode("}...), srcValue...), []any{"); err != nil {"}...)...)
	g.P("return nil, err")
	g.P("}")
}

func (f *ClientRequestEncoderGenerator) PrintNamedPathField(g *protogen.GeneratedFile, namedPathFields []*protogen.Field, httpRule *internal.HttpRule) {
	fullFieldGetterName := internal.FullFieldGetterName(namedPathFields)
	_, _, _, namedPathParameters := httpRule.RegularizePath(httpRule.Path())
	lastField := namedPathFields[len(namedPathFields)-1]
	switch lastField.Desc.Kind() {
	case protoreflect.StringKind:
		g.P("namedPathParameter := req.", fullFieldGetterName)
	case protoreflect.MessageKind:
		g.P("namedPathParameter := req.", fullFieldGetterName, ".GetValue()")
	}

	g.P("namedPathValues := ", internal.StringsPackage.Ident("Split"), "(namedPathParameter, ", strconv.Quote("/"), ")")
	g.P("if len(namedPathValues) != ", strconv.Itoa(len(namedPathParameters)*2), " {")
	g.P("return nil, ", internal.ErrorsPackage.Ident("New"), "(", strconv.Quote("invalid named path parameter, %s"), ", namedPathParameter)")
	g.P("}")

	pairs := []any{"pairs = append(pairs"}
	for i, parameter := range namedPathParameters {
		pairs = append(pairs, ",", strconv.Quote(parameter), ",", fmt.Sprintf("namedPathValues[%d]", i*2+1))
	}
	pairs = append(pairs, ")")
	g.P(pairs...)
}

func (f *ClientRequestEncoderGenerator) PrintPathField(g *protogen.GeneratedFile, pathFields []*protogen.Field) {
	pairs := []any{"pairs = append(pairs"}
	for _, field := range pathFields {
		pairs = append(append(pairs, ",", strconv.Quote(string(field.Desc.Name())), ","), f.PathFieldFormat(field)...)
	}
	pairs = append(pairs, ")")
	g.P(pairs...)
}

func (f *ClientRequestEncoderGenerator) PrintQueryField(g *protogen.GeneratedFile, queryFields []*protogen.Field) {
	for _, field := range queryFields {
		srcValue := []any{"req.Get", field.GoName, "()"}
		fieldName := string(field.Desc.Name())
		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.BoolSliceFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.BoolValueFormat(srcValue))
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.IntSliceFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.IntValueFormat(srcValue))
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.UintSliceFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.UintValueFormat(srcValue))
			}
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.IntSliceFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.IntValueFormat(srcValue))
			}
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.UintSliceFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.UintValueFormat(srcValue))
			}
		case protoreflect.FloatKind: // float32
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.FloatSliceFormat(srcValue, "32"), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.FloatValueFormat(srcValue, "32"))
			}
		case protoreflect.DoubleKind: // float64
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.FloatSliceFormat(srcValue, "64"), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.FloatValueFormat(srcValue, "64"))
			}
		case protoreflect.StringKind: // string
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.StringKindFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.StringKindFormat(srcValue))
			}
		case protoreflect.EnumKind: // enum int32
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.IntSliceFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.IntValueFormat(srcValue))
			}
		case protoreflect.MessageKind:
			switch field.Message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapFloatSliceFormat(srcValue, "64"), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapFloatValueFormat(srcValue, "64"))
				}
			case "google.protobuf.FloatValue":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapFloatSliceFormat(srcValue, "32"), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapFloatValueFormat(srcValue, "32"))
				}
			case "google.protobuf.Int64Value":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapIntSliceFormat(srcValue, "64"), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapIntValueFormat(srcValue))
				}
			case "google.protobuf.UInt64Value":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapUintSliceFormat(srcValue, "64"), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapUintValueFormat(srcValue))
				}
			case "google.protobuf.Int32Value":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapIntSliceFormat(srcValue, "32"), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapIntValueFormat(srcValue))
				}
			case "google.protobuf.UInt32Value":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapUintSliceFormat(srcValue, "32"), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapUintValueFormat(srcValue))
				}
			case "google.protobuf.BoolValue":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapBoolSliceFormat(srcValue), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapBoolValueFormat(srcValue))
				}
			case "google.protobuf.StringValue":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapStringSliceFormat(srcValue), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapStringValueFormat(srcValue))
				}
			}
		}
	}
}

func (f *ClientRequestEncoderGenerator) PathFieldFormat(field *protogen.Field) []any {
	srcValue := []any{"req.Get", field.GoName, "()"}
	switch field.Desc.Kind() {
	case protoreflect.BoolKind: // bool
		return f.BoolValueFormat(srcValue)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
		return f.IntValueFormat(srcValue)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
		return f.UintValueFormat(srcValue)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
		return f.IntValueFormat(srcValue)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
		return f.UintValueFormat(srcValue)
	case protoreflect.FloatKind: // float32
		return f.FloatValueFormat(srcValue, "32")
	case protoreflect.DoubleKind: // float64
		return f.FloatValueFormat(srcValue, "64")
	case protoreflect.StringKind: // string
		return f.StringKindFormat(srcValue)
	case protoreflect.EnumKind: //  enum int32
		return f.IntValueFormat(srcValue)
	case protoreflect.MessageKind:
		switch field.Message.Desc.FullName() {
		case "google.protobuf.DoubleValue":
			return f.UnwrapFloatValueFormat(srcValue, "64")
		case "google.protobuf.FloatValue":
			return f.UnwrapFloatValueFormat(srcValue, "32")
		case "google.protobuf.Int64Value":
			return f.UnwrapIntValueFormat(srcValue)
		case "google.protobuf.UInt64Value":
			return f.UnwrapUintValueFormat(srcValue)
		case "google.protobuf.Int32Value":
			return f.UnwrapIntValueFormat(srcValue)
		case "google.protobuf.UInt32Value":
			return f.UnwrapUintValueFormat(srcValue)
		case "google.protobuf.BoolValue":
			return f.UnwrapBoolValueFormat(srcValue)
		case "google.protobuf.StringValue":
			return f.UnwrapStringValueFormat(srcValue)
		}
	}
	return nil
}

func (f *ClientRequestEncoderGenerator) PrintQuery(g *protogen.GeneratedFile, fieldName string, srcValue []any) {
	g.P(append(append([]any{"queries[", strconv.Quote(fieldName), "] = append(queries[", strconv.Quote(fieldName), "], "}, srcValue...), []any{")"}...)...)
}

func (f *ClientRequestEncoderGenerator) BoolValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatBool"), "("}, srcValue...), []any{")"}...)
}

func (f *ClientRequestEncoderGenerator) IntValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatInt"), "("}, srcValue...), []any{", 10)"}...)
}

func (f *ClientRequestEncoderGenerator) UintValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatUint"), "("}, srcValue...), []any{", 10)"}...)
}

func (f *ClientRequestEncoderGenerator) FloatValueFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatFloat"), "("}, srcValue...), []any{", 'f', -1, ", bitSize, ")"}...)
}

func (f *ClientRequestEncoderGenerator) StringKindFormat(srcValue []any) []any {
	return append(append([]any{}, srcValue...), []any{}...)
}

func (f *ClientRequestEncoderGenerator) UnwrapFloatValueFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatFloat"), "("}, srcValue...), []any{".GetValue()", ", 'f', -1, ", bitSize, ")"}...)
}

func (f *ClientRequestEncoderGenerator) UnwrapIntValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatInt"), "("}, srcValue...), []any{".GetValue()", ", 10)"}...)
}

func (f *ClientRequestEncoderGenerator) UnwrapUintValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatUint"), "("}, srcValue...), []any{".GetValue()", ", 10)"}...)
}

func (f *ClientRequestEncoderGenerator) UnwrapBoolValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatBool"), "("}, srcValue...), []any{".GetValue()", ")"}...)
}

func (f *ClientRequestEncoderGenerator) UnwrapStringValueFormat(srcValue []any) []any {
	return append(append([]any{}, srcValue...), []any{".GetValue()"}...)
}

func (f *ClientRequestEncoderGenerator) BoolSliceFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatBoolSlice"), "("}, srcValue...), []any{")"}...)
}

func (f *ClientRequestEncoderGenerator) IntSliceFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatIntSlice"), "("}, srcValue...), []any{", 10)"}...)
}

func (f *ClientRequestEncoderGenerator) UintSliceFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatUintSlice"), "("}, srcValue...), []any{", 10)"}...)
}

func (f *ClientRequestEncoderGenerator) FloatSliceFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatFloatSlice"), "("}, srcValue...), []any{", 'f', -1, ", bitSize, ")"}...)
}

func (f *ClientRequestEncoderGenerator) UnwrapBoolSliceFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatBoolSlice"), "(", internal.ProtoxPackage.Ident("UnwrapBoolSlice"), "("}, srcValue...), []any{"))"}...)
}

func (f *ClientRequestEncoderGenerator) UnwrapFloatSliceFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatFloatSlice"), "(", internal.ProtoxPackage.Ident("UnwrapFloat" + bitSize + "Slice"), "("}, srcValue...), []any{"), 'f', -1, ", bitSize, ")"}...)
}

func (f *ClientRequestEncoderGenerator) UnwrapIntSliceFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatIntSlice"), "(", internal.ProtoxPackage.Ident("UnwrapInt" + bitSize + "Slice"), "("}, srcValue...), []any{"), 10)"}...)
}

func (f *ClientRequestEncoderGenerator) UnwrapUintSliceFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatUintSlice"), "(", internal.ProtoxPackage.Ident("UnwrapUint" + bitSize + "Slice"), "("}, srcValue...), []any{"), 10)"}...)
}

func (f *ClientRequestEncoderGenerator) UnwrapStringSliceFormat(srcValue []any) []any {
	return append(append([]any{internal.ProtoxPackage.Ident("UnwrapStringSlice"), "("}, srcValue...), []any{")"}...)
}
