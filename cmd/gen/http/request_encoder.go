package http

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

type RequestEncoderGenerator struct {
	service *internal.Service
	g       *protogen.GeneratedFile
}

func (f *RequestEncoderGenerator) GenerateRequestEncoder() error {
	f.g.P("type ", f.service.HttpClientRequestEncoderName(), " interface {")
	for _, endpoint := range f.service.Endpoints {
		f.g.P(endpoint.Name(), "(instance string) ", internal.HttpTransportPackage.Ident("CreateRequestFunc"))
	}
	f.g.P("}")
	f.g.P()
	return nil
}

func (f *RequestEncoderGenerator) GenerateClientRequestEncoderImplements() error {
	f.g.P("type ", f.service.Unexported(f.service.HttpClientRequestEncoderName()), " struct {")
	f.g.P("marshalOptions ", internal.ProtoJsonMarshalOptionsIdent)
	f.g.P("unmarshalOptions ", internal.ProtoJsonUnmarshalOptionsIdent)
	f.g.P("router *", internal.Router)
	f.g.P("scheme string")
	f.g.P("}")
	for _, endpoint := range f.service.Endpoints {
		f.g.P("func (encoder ", f.service.Unexported(f.service.HttpClientRequestEncoderName()), ")", endpoint.Name(), "(instance string) ", internal.HttpTransportPackage.Ident("CreateRequestFunc"), " {")
		httpRule := endpoint.HttpRule()
		f.g.P("return func(ctx context.Context, obj any) (*", internal.HttpPackage.Ident("Request"), ", error) {")
		f.g.P("if obj == nil {")
		f.g.P("return nil, ", internal.ErrorsPackage.Ident("New"), "(", strconv.Quote("request is nil"), ")")
		f.g.P("}")
		f.g.P("req, ok := obj.(*", endpoint.InputGoIdent(), ")")
		f.g.P("if !ok {")
		f.g.P("return nil, ", internal.FmtPackage.Ident("Errorf"), "(", strconv.Quote("invalid request type, %T"), ", obj)")
		f.g.P("}")
		f.g.P("_ = req")

		bodyMessage, bodyField, namedPathFields, pathFields, queryFields, err := endpoint.ParseParameters()
		if err != nil {
			return err
		}

		f.g.P("var body ", internal.IOPackage.Ident("Reader"))
		f.g.P("var contentType string")
		//if bodyMessage != nil {
		//	switch bodyMessage.Desc.FullName() {
		//	case "google.api.HttpBody":
		//		f.PrintReaderBlock(g, internal.BytesPackage, []any{"body"}, []any{"req.GetData()"})
		//		f.g.P("contentType := req.GetContentType()")
		//	default:
		//		f.g.P("var bodyBuf ", internal.BytesPackage.Ident("Buffer"))
		//		encoder := internal.JsonxPackage.Ident("NewEncoder")
		//		f.PrintEncodeBlock(g, encoder, []any{"&bodyBuf"}, []any{"req"})
		//		f.g.P("body = &bodyBuf")
		//		f.g.P("contentType := ", strconv.Quote(internal.JsonContentType))
		//	}
		//} else if bodyField != nil {
		//	if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
		//		f.PrintReaderBlock(g, internal.BytesPackage, []any{"body"}, []any{"req.Get", bodyField.GoName, "()", ".GetData()"})
		//		f.g.P("contentType := req.Get", bodyField.GoName, "()", ".GetContentType()")
		//	} else {
		//		f.g.P("var bodyBuf ", internal.BytesPackage.Ident("Buffer"))
		//		encoder := internal.JsonxPackage.Ident("NewEncoder")
		//		tgtValue := []any{"&bodyBuf"}
		//		srcValue := []any{"req.Get", bodyField.GoName, "()"}
		//		f.PrintEncodeBlock(g, encoder, tgtValue, srcValue)
		//		f.g.P("body = &bodyBuf")
		//		f.g.P("contentType := ", strconv.Quote(internal.JsonContentType))
		//	}
		//}
		if bodyMessage != nil {
			srcValue := []any{"req"}
			switch bodyMessage.Desc.FullName() {
			case "google.api.HttpBody":
				f.PrintEncodeHttpBodyToRequest(srcValue)
			case "google.rpc.HttpRequest":
				//f.PrintDecodeHttpRequestFromRequest(g, []any{"req"})
			default:
				f.PrintEncodeMessageToRequest(srcValue)
			}
		} else if bodyField != nil {
			switch bodyField.Desc.Kind() {
			case protoreflect.MessageKind:
				srcValue := []any{"req.Get", bodyField.GoName, "()"}
				switch bodyField.Message.Desc.FullName() {
				case "google.api.HttpBody":
					f.PrintEncodeHttpBodyToRequest(srcValue)
				default:
					//f.PrintDecodeMessageFromRequest(g, []any{"req.", bodyField.GoName})
					f.PrintEncodeMessageToRequest(srcValue)
				}
			}
		}

		f.g.P("var pairs []string")
		if len(namedPathFields) > 0 {
			f.PrintNamedPathField(namedPathFields, httpRule)
		}

		if len(pathFields) > 0 {
			f.PrintPathField(pathFields)
		}

		f.g.P("path, err := encoder.router.Get(", strconv.Quote(endpoint.FullName()), ").URLPath(pairs...)")
		f.g.P("if err != nil {")
		f.g.P("return nil, err")
		f.g.P("}")

		f.g.P("queries := ", internal.UrlPackage.Ident("Values"), "{}")
		if len(queryFields) > 0 {
			f.PrintQueryField(queryFields)
		}

		f.g.P("target := &", internal.UrlPackage.Ident("URL"), "{")
		f.g.P("Scheme:   encoder.scheme,")
		f.g.P("Host:     instance,")
		f.g.P("Path:     path.Path,")
		f.g.P("RawQuery: queries.Encode(),")
		f.g.P("}")

		f.g.P("r, err := ", internal.HttpPackage.Ident("NewRequestWithContext"), "(ctx, ", strconv.Quote(httpRule.Method()), ", target.String(), body)")
		f.g.P("if err != nil {")
		f.g.P("return nil, err")
		f.g.P("}")
		f.g.P("r.Header.Set(", strconv.Quote(internal.ContentTypeKey), ", contentType)")
		f.g.P("return r, nil")
		f.g.P("}")
		f.g.P("}")
	}
	f.g.P()
	return nil
}

func (f *RequestEncoderGenerator) PrintReaderBlock(readerPkg protogen.GoImportPath, tgtValue []any, srcValue []any) {
	f.g.P(append(append(append(append([]any{}, tgtValue...), []any{" = ", readerPkg.Ident("NewReader"), "("}...), srcValue...), ")")...)
}

func (f *RequestEncoderGenerator) PrintEncodeBlock(encoder protogen.GoIdent, tgtValue []any, srcValue []any) {
	f.g.P(append(append(append(append([]any{"if err := ", encoder, "("}, tgtValue...), []any{").Encode("}...), srcValue...), []any{"); err != nil {"}...)...)
	f.g.P("return nil, err")
	f.g.P("}")
}

func (f *RequestEncoderGenerator) PrintNamedPathField(namedPathFields []*protogen.Field, httpRule *internal.HttpRule) {
	fullFieldGetterName := internal.FullFieldGetterName(namedPathFields)
	_, _, _, namedPathParameters := httpRule.RegularizePath(httpRule.Path())
	lastField := namedPathFields[len(namedPathFields)-1]
	switch lastField.Desc.Kind() {
	case protoreflect.StringKind:
		f.g.P("namedPathParameter := req.", fullFieldGetterName)
	case protoreflect.MessageKind:
		f.g.P("namedPathParameter := req.", fullFieldGetterName, ".GetValue()")
	}

	f.g.P("namedPathValues := ", internal.StringsPackage.Ident("Split"), "(namedPathParameter, ", strconv.Quote("/"), ")")
	f.g.P("if len(namedPathValues) != ", strconv.Itoa(len(namedPathParameters)*2), " {")
	f.g.P("return nil, ", internal.ErrorsPackage.Ident("New"), "(", strconv.Quote("invalid named path parameter, %s"), ", namedPathParameter)")
	f.g.P("}")

	pairs := []any{"pairs = append(pairs"}
	for i, parameter := range namedPathParameters {
		pairs = append(pairs, ",", strconv.Quote(parameter), ",", fmt.Sprintf("namedPathValues[%d]", i*2+1))
	}
	pairs = append(pairs, ")")
	f.g.P(pairs...)
}

func (f *RequestEncoderGenerator) PrintPathField(pathFields []*protogen.Field) {
	pairs := []any{"pairs = append(pairs"}
	for _, field := range pathFields {
		pairs = append(append(pairs, ",", strconv.Quote(string(field.Desc.Name())), ","), f.PathFieldFormat(field)...)
	}
	pairs = append(pairs, ")")
	f.g.P(pairs...)
}

func (f *RequestEncoderGenerator) PrintQueryField(queryFields []*protogen.Field) {
	for _, field := range queryFields {
		srcValue := []any{"req.Get", field.GoName, "()"}
		fieldName := string(field.Desc.Name())
		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
			if field.Desc.IsList() {
				f.PrintQuery(fieldName, append(f.BoolSliceFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(fieldName, f.BoolValueFormat(srcValue))
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
			if field.Desc.IsList() {
				f.PrintQuery(fieldName, append(f.IntSliceFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(fieldName, f.IntValueFormat(srcValue))
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
			if field.Desc.IsList() {
				f.PrintQuery(fieldName, append(f.UintSliceFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(fieldName, f.UintValueFormat(srcValue))
			}
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
			if field.Desc.IsList() {
				f.PrintQuery(fieldName, append(f.IntSliceFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(fieldName, f.IntValueFormat(srcValue))
			}
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
			if field.Desc.IsList() {
				f.PrintQuery(fieldName, append(f.UintSliceFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(fieldName, f.UintValueFormat(srcValue))
			}
		case protoreflect.FloatKind: // float32
			if field.Desc.IsList() {
				f.PrintQuery(fieldName, append(f.FloatSliceFormat(srcValue, "32"), []any{"..."}...))
			} else {
				f.PrintQuery(fieldName, f.FloatValueFormat(srcValue, "32"))
			}
		case protoreflect.DoubleKind: // float64
			if field.Desc.IsList() {
				f.PrintQuery(fieldName, append(f.FloatSliceFormat(srcValue, "64"), []any{"..."}...))
			} else {
				f.PrintQuery(fieldName, f.FloatValueFormat(srcValue, "64"))
			}
		case protoreflect.StringKind: // string
			if field.Desc.IsList() {
				f.PrintQuery(fieldName, append(f.StringKindFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(fieldName, f.StringKindFormat(srcValue))
			}
		case protoreflect.EnumKind: // enum int32
			if field.Desc.IsList() {
				f.PrintQuery(fieldName, append(f.IntSliceFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(fieldName, f.IntValueFormat(srcValue))
			}
		case protoreflect.MessageKind:
			switch field.Message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
				if field.Desc.IsList() {
					f.PrintQuery(fieldName, append(f.UnwrapFloatSliceFormat(srcValue, "64"), []any{"..."}...))
				} else {
					f.PrintQuery(fieldName, f.UnwrapFloatValueFormat(srcValue, "64"))
				}
			case "google.protobuf.FloatValue":
				if field.Desc.IsList() {
					f.PrintQuery(fieldName, append(f.UnwrapFloatSliceFormat(srcValue, "32"), []any{"..."}...))
				} else {
					f.PrintQuery(fieldName, f.UnwrapFloatValueFormat(srcValue, "32"))
				}
			case "google.protobuf.Int64Value":
				if field.Desc.IsList() {
					f.PrintQuery(fieldName, append(f.UnwrapIntSliceFormat(srcValue, "64"), []any{"..."}...))
				} else {
					f.PrintQuery(fieldName, f.UnwrapIntValueFormat(srcValue))
				}
			case "google.protobuf.UInt64Value":
				if field.Desc.IsList() {
					f.PrintQuery(fieldName, append(f.UnwrapUintSliceFormat(srcValue, "64"), []any{"..."}...))
				} else {
					f.PrintQuery(fieldName, f.UnwrapUintValueFormat(srcValue))
				}
			case "google.protobuf.Int32Value":
				if field.Desc.IsList() {
					f.PrintQuery(fieldName, append(f.UnwrapIntSliceFormat(srcValue, "32"), []any{"..."}...))
				} else {
					f.PrintQuery(fieldName, f.UnwrapIntValueFormat(srcValue))
				}
			case "google.protobuf.UInt32Value":
				if field.Desc.IsList() {
					f.PrintQuery(fieldName, append(f.UnwrapUintSliceFormat(srcValue, "32"), []any{"..."}...))
				} else {
					f.PrintQuery(fieldName, f.UnwrapUintValueFormat(srcValue))
				}
			case "google.protobuf.BoolValue":
				if field.Desc.IsList() {
					f.PrintQuery(fieldName, append(f.UnwrapBoolSliceFormat(srcValue), []any{"..."}...))
				} else {
					f.PrintQuery(fieldName, f.UnwrapBoolValueFormat(srcValue))
				}
			case "google.protobuf.StringValue":
				if field.Desc.IsList() {
					f.PrintQuery(fieldName, append(f.UnwrapStringSliceFormat(srcValue), []any{"..."}...))
				} else {
					f.PrintQuery(fieldName, f.UnwrapStringValueFormat(srcValue))
				}
			}
		}
	}
}

func (f *RequestEncoderGenerator) PathFieldFormat(field *protogen.Field) []any {
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

func (f *RequestEncoderGenerator) PrintQuery(fieldName string, srcValue []any) {
	f.g.P(append(append([]any{"queries[", strconv.Quote(fieldName), "] = append(queries[", strconv.Quote(fieldName), "], "}, srcValue...), []any{")"}...)...)
}

func (f *RequestEncoderGenerator) BoolValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatBool"), "("}, srcValue...), []any{")"}...)
}

func (f *RequestEncoderGenerator) IntValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatInt"), "("}, srcValue...), []any{", 10)"}...)
}

func (f *RequestEncoderGenerator) UintValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatUint"), "("}, srcValue...), []any{", 10)"}...)
}

func (f *RequestEncoderGenerator) FloatValueFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatFloat"), "("}, srcValue...), []any{", 'f', -1, ", bitSize, ")"}...)
}

func (f *RequestEncoderGenerator) StringKindFormat(srcValue []any) []any {
	return append(append([]any{}, srcValue...), []any{}...)
}

func (f *RequestEncoderGenerator) UnwrapFloatValueFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatFloat"), "("}, srcValue...), []any{".GetValue()", ", 'f', -1, ", bitSize, ")"}...)
}

func (f *RequestEncoderGenerator) UnwrapIntValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatInt"), "("}, srcValue...), []any{".GetValue()", ", 10)"}...)
}

func (f *RequestEncoderGenerator) UnwrapUintValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatUint"), "("}, srcValue...), []any{".GetValue()", ", 10)"}...)
}

func (f *RequestEncoderGenerator) UnwrapBoolValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatBool"), "("}, srcValue...), []any{".GetValue()", ")"}...)
}

func (f *RequestEncoderGenerator) UnwrapStringValueFormat(srcValue []any) []any {
	return append(append([]any{}, srcValue...), []any{".GetValue()"}...)
}

func (f *RequestEncoderGenerator) BoolSliceFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatBoolSlice"), "("}, srcValue...), []any{")"}...)
}

func (f *RequestEncoderGenerator) IntSliceFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatIntSlice"), "("}, srcValue...), []any{", 10)"}...)
}

func (f *RequestEncoderGenerator) UintSliceFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatUintSlice"), "("}, srcValue...), []any{", 10)"}...)
}

func (f *RequestEncoderGenerator) FloatSliceFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatFloatSlice"), "("}, srcValue...), []any{", 'f', -1, ", bitSize, ")"}...)
}

func (f *RequestEncoderGenerator) UnwrapBoolSliceFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatBoolSlice"), "(", internal.ProtoxPackage.Ident("UnwrapBoolSlice"), "("}, srcValue...), []any{"))"}...)
}

func (f *RequestEncoderGenerator) UnwrapFloatSliceFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatFloatSlice"), "(", internal.ProtoxPackage.Ident("UnwrapFloat" + bitSize + "Slice"), "("}, srcValue...), []any{"), 'f', -1, ", bitSize, ")"}...)
}

func (f *RequestEncoderGenerator) UnwrapIntSliceFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatIntSlice"), "(", internal.ProtoxPackage.Ident("UnwrapInt" + bitSize + "Slice"), "("}, srcValue...), []any{"), 10)"}...)
}

func (f *RequestEncoderGenerator) UnwrapUintSliceFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatUintSlice"), "(", internal.ProtoxPackage.Ident("UnwrapUint" + bitSize + "Slice"), "("}, srcValue...), []any{"), 10)"}...)
}

func (f *RequestEncoderGenerator) UnwrapStringSliceFormat(srcValue []any) []any {
	return append(append([]any{internal.ProtoxPackage.Ident("UnwrapStringSlice"), "("}, srcValue...), []any{")"}...)
}

func (f *RequestEncoderGenerator) PrintEncodeHttpBodyToRequest(srcValue []any) {
	f.g.P(append(append([]any{"body, contentType = ", internal.EncodeHttpBodyToRequest, "(ctx, "}, srcValue...), ")")...)
}

func (f *RequestEncoderGenerator) PrintEncodeMessageToRequest(srcValue []any) {
	f.g.P(append(append([]any{"body, contentType, err = ", internal.EncodeMessageToRequest, "(ctx, "}, srcValue...), ", encoder.marshalOptions)")...)
}
