package http

import (
	"github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

type ServerRequestDecoderGenerator struct {
	service *internal.Service
	g       *protogen.GeneratedFile
}

func (f *ServerRequestDecoderGenerator) GenerateServerRequestDecoder() {
	f.g.P("type ", f.service.HttpServerRequestDecoderName(), " interface {")
	for _, endpoint := range f.service.Endpoints {
		f.g.P(endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("DecodeRequestFunc"))
	}
	f.g.P("}")
	f.g.P()
}

func (f *ServerRequestDecoderGenerator) GenerateServerRequestDecoderImplements() {
	f.g.P("type ", f.service.Unexported(f.service.HttpServerRequestDecoderName()), " struct {")
	f.g.P("unmarshalOptions ", internal.ProtoJsonUnmarshalOptionsIdent)
	f.g.P("}")
	for _, endpoint := range f.service.Endpoints {
		f.g.P("func (decoder ", f.service.Unexported(f.service.HttpServerRequestDecoderName()), ")", endpoint.Name(), "() ", internal.HttpTransportPackage.Ident("DecodeRequestFunc"), "{")
		f.g.P("return func ", "(ctx ", internal.Context, ", r *", internal.Request, ") (any, error) {")
		f.g.P("req := &", endpoint.InputGoIdent(), "{}")

		if bodyMessage := endpoint.BodyMessage(); bodyMessage != nil {
			switch bodyMessage.Desc.FullName() {
			case "google.api.HttpBody":
				f.PrintDecodeHttpBodyFromRequest([]any{"req"})
			case "google.rpc.HttpRequest":
				f.PrintDecodeHttpRequestFromRequest([]any{"req"})
			default:
				f.PrintDecodeMessageFromRequest([]any{"req"})
			}
		} else if bodyField := endpoint.BodyField(); bodyField != nil {
			tgtValue := []any{"req.Get", bodyField.GoName, "()"}
			f.g.P("req.", bodyField.GoName, " = &", bodyField.Message.GoIdent, "{}")
			switch bodyField.Desc.Kind() {
			case protoreflect.MessageKind:
				switch bodyField.Message.Desc.FullName() {
				case "google.api.HttpBody":
					f.PrintDecodeHttpBodyFromRequest(tgtValue)
				default:
					f.PrintDecodeMessageFromRequest(tgtValue)
				}
			}
		}

		if namedPathFields, pathFields := endpoint.NamedPathFields(), endpoint.PathFields(); len(namedPathFields)+len(pathFields) > 0 {
			f.g.P("vars := ", internal.UrlxPackage.Ident("FormFromMap"), "(", internal.VarsIdent, "(r)", ")")
			f.PrintNamedPathField(namedPathFields, endpoint)
			f.PrintPathField(pathFields)
		}

		if queryFields := endpoint.QueryFields(); len(queryFields) > 0 {
			f.g.P("queries := r.URL.Query()")
			f.g.P("var queryErr error")
			f.PrintQueryField(queryFields)
			f.g.P("if queryErr != nil {")
			f.g.P("return nil, queryErr")
			f.g.P("}")
		}

		f.g.P("return req, nil")
		f.g.P("}")
		f.g.P("}")
	}
	f.g.P()
}

func (f *ServerRequestDecoderGenerator) PrintDecodeMessageFromRequest(tgtValue []any) {
	f.g.P(append(append([]any{"if err := ", internal.DecodeMessageFromRequest, "(ctx, r, "}, tgtValue...), ", decoder.unmarshalOptions); err != nil {")...)
	f.g.P("return nil, err")
	f.g.P("}")
}

func (f *ServerRequestDecoderGenerator) PrintDecodeHttpBodyFromRequest(tgtValue []any) {
	f.g.P(append(append([]any{"if err := ", internal.DecodeHttpBodyFromRequest, "(ctx, r, "}, tgtValue...), "); err != nil {")...)
	f.g.P("return nil, err")
	f.g.P("}")
}

func (f *ServerRequestDecoderGenerator) PrintDecodeHttpRequestFromRequest(tgtValue []any) {
	f.g.P(append(append([]any{"if err := ", internal.DecodeHttpRequestFromRequest, "(ctx, r, "}, tgtValue...), "); err != nil {")...)
	f.g.P("return nil, err")
	f.g.P("}")
}

func (f *ServerRequestDecoderGenerator) PrintNamedPathField(namedPathFields []*protogen.Field, endpoint *internal.Endpoint) {
	for i, namedPathField := range namedPathFields {
		fullFieldName := internal.FullFieldName(namedPathFields[:i+1])
		if i < len(namedPathFields)-1 {
			f.g.P("if req.", fullFieldName, " == nil {")
			f.g.P("req.", fullFieldName, " = &", namedPathField.Message.GoIdent, "{}")
			f.g.P("}")
		} else {
			tgtValue := []any{"req.", fullFieldName, " = "}
			srcValue := []any{internal.Sprintf, "(", strconv.Quote(endpoint.NamedPathTemplate())}
			for _, namedPathParameter := range endpoint.NamedPathFieldsParameters() {
				srcValue = append(srcValue, ", vars.Get(", strconv.Quote(namedPathParameter), ")")
			}
			srcValue = append(srcValue, ")")

			switch namedPathField.Desc.Kind() {
			case protoreflect.StringKind:
				f.PrintStringValueAssign(tgtValue, srcValue, namedPathField.Desc.HasPresence())
			case protoreflect.MessageKind:
				f.PrintWrapStringValueAssign(tgtValue, srcValue)
			}
		}
	}
}

func (f *ServerRequestDecoderGenerator) PrintPathField(pathFields []*protogen.Field) {
	if len(pathFields) <= 0 {
		return
	}
	form := "vars"
	errName := "varErr"
	f.g.P("var ", errName, " error")
	for _, field := range pathFields {
		fieldName := string(field.Desc.Name())

		tgtValue := []any{"req.", field.GoName, " = "}
		tgtErrValue := []any{"req.", field.GoName, ", ", errName, " = "}
		srcValue := []any{"vars.Get(", strconv.Quote(fieldName), ")"}

		goType, pointer := internal.FieldGoType(f.g, field)
		if pointer {
			goType = append([]any{"*"}, goType...)
		}

		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
			if pointer {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetBool"), fieldName, form, errName)
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
			if pointer {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt"), fieldName, form, errName)
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
			if pointer {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint"), fieldName, form, errName)
			}
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
			if pointer {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt"), fieldName, form, errName)
			}
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
			if pointer {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint"), fieldName, form, errName)
			}
		case protoreflect.FloatKind: // float32
			if pointer {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat"), fieldName, form, errName)
			}
		case protoreflect.DoubleKind: // float64
			if pointer {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat"), fieldName, form, errName)
			}
		case protoreflect.StringKind: // string
			f.PrintStringValueAssign(tgtValue, srcValue, pointer)
		case protoreflect.EnumKind: // enum int32
			if pointer {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr["+f.g.QualifiedGoIdent(goType[1].(protogen.GoIdent))+"]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt["+f.g.QualifiedGoIdent(goType[0].(protogen.GoIdent))+"]"), fieldName, form, errName)
			}
		case protoreflect.MessageKind:
			switch field.Message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat64Value"), fieldName, form, errName)
			case "google.protobuf.FloatValue":
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat32Value"), fieldName, form, errName)
			case "google.protobuf.Int64Value":
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt64Value"), fieldName, form, errName)
			case "google.protobuf.UInt64Value":
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint64Value"), fieldName, form, errName)
			case "google.protobuf.Int32Value":
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt32Value"), fieldName, form, errName)
			case "google.protobuf.UInt32Value":
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint32Value"), fieldName, form, errName)
			case "google.protobuf.BoolValue":
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolValue"), fieldName, form, errName)
			case "google.protobuf.StringValue":
				f.PrintWrapStringValueAssign(tgtValue, srcValue)
			}
		}
	}
	f.g.P("if ", errName, " != nil {")
	f.g.P("return nil, ", errName)
	f.g.P("}")
}

func (f *ServerRequestDecoderGenerator) PrintQueryField(queryFields []*protogen.Field) {
	for _, field := range queryFields {
		fieldName := string(field.Desc.Name())

		tgtValue := []any{"req.", field.GoName, " = "}
		tgtErrValue := []any{"req.", field.GoName, ", queryErr = "}
		srcValue := []any{"queries.Get(", strconv.Quote(fieldName), ")"}
		if field.Desc.IsList() {
			srcValue = []any{"queries[", strconv.Quote(fieldName), "]"}
		}

		goType, pointer := internal.FieldGoType(f.g, field)
		if pointer {
			goType = append([]any{"*"}, goType...)
		}

		form := "queries"
		errName := "queryErr"

		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
			if field.Desc.IsList() {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolSlice"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolPtr"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetBool"), fieldName, form, errName)
				}
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
			if field.Desc.IsList() {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntSlice[int32]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr[int32]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt[int32]"), fieldName, form, errName)
				}
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
			if field.Desc.IsList() {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintSlice[uint32]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintPtr[uint32]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint[uint32]"), fieldName, form, errName)
				}
			}
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
			if field.Desc.IsList() {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntSlice[int64]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr[int64]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt[int64]"), fieldName, form, errName)
				}
			}
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
			if field.Desc.IsList() {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintSlice[uint64]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintPtr[uint64]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint[uint64]"), fieldName, form, errName)
				}
			}
		case protoreflect.FloatKind: // float32
			if field.Desc.IsList() {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatSlice[float32]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatPtr[float32]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat[float32]"), fieldName, form, errName)
				}
			}
		case protoreflect.DoubleKind: // float64
			if field.Desc.IsList() {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatSlice[float64]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatPtr[float64]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat[float64]"), fieldName, form, errName)
				}
			}
		case protoreflect.StringKind: // string
			if field.Desc.IsList() {
				f.PrintStringListAssign(tgtValue, srcValue)
			} else {
				f.PrintStringValueAssign(tgtValue, srcValue, pointer)
			}
		case protoreflect.EnumKind: // enum int32
			if field.Desc.IsList() {
				f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntSlice["+f.g.QualifiedGoIdent(goType[1].(protogen.GoIdent))+"]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr["+f.g.QualifiedGoIdent(goType[1].(protogen.GoIdent))+"]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt["+f.g.QualifiedGoIdent(goType[0].(protogen.GoIdent))+"]"), fieldName, form, errName)
				}
			}
		case protoreflect.MessageKind:
			switch field.Message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
				if field.Desc.IsList() {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat64ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat64Value"), fieldName, form, errName)
				}
			case "google.protobuf.FloatValue":
				if field.Desc.IsList() {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat32ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat32Value"), fieldName, form, errName)
				}
			case "google.protobuf.Int64Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt64ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt64Value"), fieldName, form, errName)
				}
			case "google.protobuf.UInt64Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint64ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint64Value"), fieldName, form, errName)
				}
			case "google.protobuf.Int32Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt32ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt32Value"), fieldName, form, errName)
				}
			case "google.protobuf.UInt32Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint32ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint32Value"), fieldName, form, errName)
				}
			case "google.protobuf.BoolValue":
				if field.Desc.IsList() {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolValue"), fieldName, form, errName)
				}
			case "google.protobuf.StringValue":
				if field.Desc.IsList() {
					f.PrintWrapStringListAssign(tgtValue, srcValue)
				} else {
					f.PrintWrapStringValueAssign(tgtValue, srcValue)
				}
			}
		}
	}

}

func (f *ServerRequestDecoderGenerator) PrintFieldAssign(tgtValue []any, goType []any, getter protogen.GoIdent, key string, form string, errName string) {
	f.g.P(append(append([]any{}, tgtValue...), append(append([]any{internal.DecodeForm, "["}, goType...), append([]any{"](", errName, ", ", form, ", ", strconv.Quote(key), ", ", getter}, ")")...)...)...)
}

func (f *ServerRequestDecoderGenerator) PrintStringValueAssign(tgtValue []any, srcValue []any, hasPresence bool) {
	if hasPresence {
		f.g.P(append(tgtValue, append(append([]any{internal.ProtoPackage.Ident("String"), "("}, srcValue...), ")")...)...)
	} else {
		f.g.P(append(tgtValue, srcValue...)...)
	}
}

func (f *ServerRequestDecoderGenerator) PrintWrapStringValueAssign(tgtValue []any, srcValue []any) {
	f.g.P(append(tgtValue, append(append([]any{internal.WrapperspbPackage.Ident("String"), "("}, srcValue...), ")")...)...)
}

func (f *ServerRequestDecoderGenerator) PrintStringListAssign(tgtValue []any, srcValue []any) {
	f.g.P(append(tgtValue, srcValue...)...)
}

func (f *ServerRequestDecoderGenerator) PrintWrapStringListAssign(tgtValue []any, srcValue []any) {
	f.g.P(append(tgtValue, append(append([]any{internal.ProtoxPackage.Ident("WrapStringSlice"), "("}, srcValue...), ")")...)...)
}
