package http

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

type RequestEncoderGenerator struct {
	service *internal.Service
	g       *protogen.GeneratedFile
}

func (f *RequestEncoderGenerator) GenerateRequestEncoder() {
	f.g.P("type ", f.service.HttpClientRequestEncoderName(), " interface {")
	for _, endpoint := range f.service.Endpoints {
		f.g.P(endpoint.Name(), "(instance string) ", internal.HttpTransportPackage.Ident("CreateRequestFunc"))
	}
	f.g.P("}")
	f.g.P()
}

func (f *RequestEncoderGenerator) GenerateClientRequestEncoderImplements() {
	f.g.P("type ", f.service.Unexported(f.service.HttpClientRequestEncoderName()), " struct {")
	f.g.P("marshalOptions ", internal.ProtoJsonMarshalOptionsIdent)
	f.g.P("router *", internal.Router)
	f.g.P("scheme string")
	f.g.P("}")
	for _, endpoint := range f.service.Endpoints {
		f.g.P("func (encoder ", f.service.Unexported(f.service.HttpClientRequestEncoderName()), ")", endpoint.Name(), "(instance string) ", internal.HttpTransportPackage.Ident("CreateRequestFunc"), " {")
		f.g.P("return func(ctx context.Context, obj any) (*", internal.Request, ", error) {")
		f.g.P("if obj == nil {")
		f.g.P("return nil, ", internal.ErrorsPackage.Ident("New"), "(", strconv.Quote("request is nil"), ")")
		f.g.P("}")
		f.g.P("req, ok := obj.(*", endpoint.InputGoIdent(), ")")
		f.g.P("if !ok {")
		f.g.P("return nil, ", internal.Errorf, "(", strconv.Quote("invalid request type, %T"), ", obj)")
		f.g.P("}")
		f.g.P("_ = req")

		f.g.P("method := ", endpoint.HttpMethod())
		f.g.P("target := &", internal.UrlPackage.Ident("URL"), "{")
		f.g.P("Scheme:   encoder.scheme,")
		f.g.P("Host:     instance,")
		f.g.P("}")
		f.g.P("header := ", internal.Header, "{}")
		f.g.P("var body ", internal.Buffer)

		if bodyMessage := endpoint.BodyMessage(); bodyMessage != nil {
			srcValue := []any{"req"}
			switch bodyMessage.Desc.FullName() {
			case "google.api.HttpBody":
				f.PrintEncodeHttpBodyToRequest(srcValue)
			case "google.rpc.HttpRequest":
				f.PrintEncodeHttpRequestToRequest(srcValue)
			default:
				f.PrintEncodeMessageToRequest(srcValue)
			}
		} else if bodyField := endpoint.BodyField(); bodyField != nil {
			switch bodyField.Desc.Kind() {
			case protoreflect.MessageKind:
				srcValue := []any{"req.Get", bodyField.GoName, "()"}
				switch bodyField.Message.Desc.FullName() {
				case "google.api.HttpBody":
					f.PrintEncodeHttpBodyToRequest(srcValue)
				default:
					f.PrintEncodeMessageToRequest(srcValue)
				}
			}
		}

		f.g.P("var pairs []string")
		f.PrintNamedPathField(endpoint.NamedPathFields(), endpoint)
		f.PrintPathField(endpoint.PathFields())
		f.g.P("path, err := encoder.router.Get(", strconv.Quote(endpoint.FullName()), ").URLPath(pairs...)")
		f.g.P("if err != nil {")
		f.g.P("return nil, err")
		f.g.P("}")
		f.g.P("target.Path = path.Path")

		f.PrintQueryField(endpoint.QueryFields())

		f.g.P("r, err := ", internal.NewRequestWithContext, "(ctx, method, target.String(), &body)")
		f.g.P("if err != nil {")
		f.g.P("return nil, err")
		f.g.P("}")
		f.g.P(internal.CopyHeader, "(r.Header, header)")
		f.g.P("return r, nil")
		f.g.P("}")
		f.g.P("}")
	}
	f.g.P()
}

func (f *RequestEncoderGenerator) PrintEncodeMessageToRequest(srcValue []any) {
	f.g.P(append(append([]any{"if err := ", internal.EncodeMessageToRequest, "(ctx, "}, srcValue...), ", header, &body, encoder.marshalOptions); err!= nil {")...)
	f.g.P("return nil, err")
	f.g.P("}")
}

func (f *RequestEncoderGenerator) PrintEncodeHttpBodyToRequest(srcValue []any) {
	f.g.P(append(append([]any{"if err := ", internal.EncodeHttpBodyToRequest, "(ctx, "}, srcValue...), ", header, &body); err!= nil {")...)
	f.g.P("return nil, err")
	f.g.P("}")
}

func (f *RequestEncoderGenerator) PrintEncodeHttpRequestToRequest(srcValue []any) {
	f.g.P(append(append([]any{"if err := ", internal.EncodeHttpRequestToRequest, "(ctx, "}, srcValue...), ", header, &body); err!= nil {")...)
	f.g.P("return nil, err")
	f.g.P("}")
}

func (f *RequestEncoderGenerator) PrintNamedPathField(namedPathFields []*protogen.Field, endpoint *internal.Endpoint) {
	if len(namedPathFields) <= 0 {
		return
	}
	namedPathParameters := endpoint.NamedPathFieldsParameters()
	fullFieldGetterName := internal.FullFieldGetterName(namedPathFields)
	lastField := namedPathFields[len(namedPathFields)-1]
	switch lastField.Desc.Kind() {
	case protoreflect.StringKind:
		f.g.P("namedPathParameter := req.", fullFieldGetterName)
	case protoreflect.MessageKind:
		f.g.P("namedPathParameter := req.", fullFieldGetterName, ".GetValue()")
	}

	f.g.P("namedPathValues := ", internal.StringsPackage.Ident("Split"), "(namedPathParameter, ", strconv.Quote("/"), ")")
	f.g.P("if len(namedPathValues) != ", strconv.Itoa(len(namedPathParameters)*2), " {")
	f.g.P("return nil, ", internal.Errorf, "(", strconv.Quote("invalid named path parameter, %s"), ", namedPathParameter)")
	f.g.P("}")

	f.g.P("pairs = append(pairs, ")
	for i, parameter := range namedPathParameters {
		f.g.P(append(append([]any{strconv.Quote(parameter), ", "}, fmt.Sprintf("namedPathValues[%d]", i*2+1)), ",")...)
	}
	f.g.P(")")
}

func (f *RequestEncoderGenerator) PrintPathField(pathFields []*protogen.Field) {
	if len(pathFields) <= 0 {
		return
	}
	f.g.P("pairs = append(pairs, ")
	for _, field := range pathFields {
		f.g.P(append(append([]any{strconv.Quote(string(field.Desc.Name())), ", "}, f.PathFieldFormat(field)...), ",")...)
	}
	f.g.P(")")
}

func (f *RequestEncoderGenerator) PrintQueryField(queryFields []*protogen.Field) {
	if len(queryFields) <= 0 {
		return
	}
	f.g.P("queries := ", internal.UrlPackage.Ident("Values"), "{}")
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
	f.g.P("target.RawQuery = queries.Encode()")
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
