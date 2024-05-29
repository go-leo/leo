package generator

import (
	"errors"
	"fmt"
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

type ServerGenerator struct{}

func (f *ServerGenerator) GenerateNewServer(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("func New", service.HTTPServerName(), "(")
	generatedFile.P("endpoints interface {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.Name(), "() ", internal.EndpointPackage.Ident("Endpoint"))
	}
	generatedFile.P("},")
	generatedFile.P("opts []", internal.HttpTransportPackage.Ident("ServerOption"), ",")
	generatedFile.P("middlewares ...", internal.EndpointPackage.Ident("Middleware"), ",")
	generatedFile.P(") ", internal.HttpPackage.Ident("Handler"), " {")
	generatedFile.P("router := ", internal.MuxPackage.Ident("NewRouter"), "()")
	for _, endpoint := range service.Endpoints {
		httpRule := endpoint.HttpRule()
		// 调整路径，来适应 github.com/gorilla/mux 路由规则
		path, _, _, _ := httpRule.RegularizePath(httpRule.Path())
		generatedFile.P("router.NewRoute().")
		generatedFile.P("Name(", strconv.Quote(endpoint.FullName()), ").")
		generatedFile.P("Methods(", strconv.Quote(httpRule.Method()), ").")
		generatedFile.P("Path(", strconv.Quote(path), ").")
		generatedFile.P("Handler(", internal.HttpTransportPackage.Ident("NewServer"), "(")
		generatedFile.P(internal.EndpointxPackage.Ident("Chain"), "(endpoints.", endpoint.Name(), "(), middlewares...), ")
		if err := f.PrintDecodeRequestFunc(generatedFile, endpoint); err != nil {
			return err
		}
		if err := f.PrintEncodeResponseFunc(generatedFile, endpoint, httpRule); err != nil {
			return err
		}

		generatedFile.P("},")
		generatedFile.P("opts...,")
		generatedFile.P("))")
	}
	generatedFile.P("return router")
	generatedFile.P("}")
	generatedFile.P()
	return nil
}

func (f *ServerGenerator) PrintDecodeRequestFunc(
	generatedFile *protogen.GeneratedFile, endpoint *internal.Endpoint,
) error {
	generatedFile.P("func(ctx ", internal.ContextPackage.Ident("Context"), ", r *", internal.HttpPackage.Ident("Request"), ") (any, error) {")
	generatedFile.P("req := &", endpoint.InputGoIdent(), "{}")

	bodyMessage, bodyField, namedPathFields, pathFields, queryFields, err := endpoint.ParseParameters()
	if err != nil {
		return err
	}

	if bodyMessage != nil {
		switch bodyMessage.Desc.FullName() {
		case "google.api.HttpBody":
			f.PrintGoogleApiHttpBodyDecodeBlock(generatedFile, []any{"req"}, []any{"r.Body"})
		case "google.rpc.HttpRequest":
			f.PrintGoogleRpcHttpRequestDecodeBlock(generatedFile)
			generatedFile.P("return req, nil")
			generatedFile.P("},")
			return nil
		default:
			f.PrintDecodeBlock(generatedFile, internal.JsonxPackage.Ident("NewDecoder"), []any{"req"}, []any{"r.Body"})
		}
	} else if bodyField != nil {
		if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
			generatedFile.P("req.", bodyField.GoName, " = &", bodyField.Message.GoIdent, "{}")
			tgtValue := []any{"req.", bodyField.GoName}
			srcValue := []any{"r.Body"}
			f.PrintGoogleApiHttpBodyDecodeBlock(generatedFile, tgtValue, srcValue)
		} else {
			tgtValue := []any{"&req.", bodyField.GoName}
			srcValue := []any{"r.Body"}
			f.PrintDecodeBlock(generatedFile, internal.JsonxPackage.Ident("NewDecoder"), tgtValue, srcValue)
		}
	}

	if len(namedPathFields)+len(pathFields) > 0 {
		generatedFile.P("vars := ", internal.UrlxPackage.Ident("FormFromMap"), "(", internal.MuxPackage.Ident("Vars"), "(r)", ")")
		generatedFile.P("var varErr error")
		if err := f.PrintNamedPathField(generatedFile, namedPathFields, endpoint.HttpRule()); err != nil {
			return err
		}
		f.PrintPathField(generatedFile, pathFields)
		generatedFile.P("if varErr != nil {")
		generatedFile.P("return nil, varErr")
		generatedFile.P("}")
	}

	if len(queryFields) > 0 {
		generatedFile.P("queries := r.URL.Query()")
		generatedFile.P("var queryErr error")
		f.PrintQueryField(generatedFile, queryFields)
		generatedFile.P("if queryErr != nil {")
		generatedFile.P("return nil, queryErr")
		generatedFile.P("}")
	}

	generatedFile.P("return req, nil")
	generatedFile.P("},")
	return nil
}

func (f *ServerGenerator) PrintGoogleApiHttpBodyDecodeBlock(generatedFile *protogen.GeneratedFile, tgtValue []any, srcValue []any) {
	generatedFile.P(append(append([]any{"body, err := ", internal.IOPackage.Ident("ReadAll"), "("}, srcValue...), []any{")"}...)...)
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P(append(append([]any{}, tgtValue...), []any{".Data = body"}...)...)
	generatedFile.P(append(append([]any{}, tgtValue...), []any{".ContentType = r.Header.Get(", strconv.Quote(internal.ContentTypeKey), ")"}...)...)
}

func (f *ServerGenerator) PrintGoogleRpcHttpRequestDecodeBlock(generatedFile *protogen.GeneratedFile) {
	generatedFile.P("req.Method = r.Method")
	generatedFile.P("req.Uri = r.URL.String()")
	generatedFile.P("req.Headers = make([]*", internal.RpcHttpPackage.Ident("HttpHeader"), ", 0, len(r.Header))")
	generatedFile.P("for key, values := range r.Header {")
	generatedFile.P("for _, value := range values {")
	generatedFile.P("req.Headers = append(", "req.Headers, &", internal.RpcHttpPackage.Ident("HttpHeader"), "{Key: key, Value: value})")
	generatedFile.P("}")
	generatedFile.P("}")
	generatedFile.P("body, err := ", internal.IOPackage.Ident("ReadAll"), "(r.Body)")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P("req.Body = body")
}

func (f *ServerGenerator) PrintDecodeBlock(generatedFile *protogen.GeneratedFile, decoder protogen.GoIdent, tgtValue []any, srcValue []any) {
	generatedFile.P(append(append(append(append([]any{"if err := ", decoder, "("}, srcValue...), []any{").Decode("}...), tgtValue...), []any{"); err != nil {"}...)...)
	generatedFile.P("return nil, err")
	generatedFile.P("}")
}

func (f *ServerGenerator) PrintNamedPathField(generatedFile *protogen.GeneratedFile, namedPathFields []*protogen.Field, httpRule *internal.HttpRule) error {
	for i, namedPathField := range namedPathFields {
		fullFieldName := internal.FullFieldName(namedPathFields[:i+1])
		if i < len(namedPathFields)-1 {
			generatedFile.P("if req.", fullFieldName, " == nil {")
			generatedFile.P("req.", fullFieldName, " = &", namedPathField.Message.GoIdent, "{}")
			generatedFile.P("}")
		} else {
			_, _, namedPathTemplate, namedPathParameters := httpRule.RegularizePath(httpRule.Path())
			tgtValue := []any{"req.", fullFieldName, " = "}
			srcValue := []any{internal.FmtPackage.Ident("Sprintf"), "(", strconv.Quote(namedPathTemplate)}
			for _, namedPathParameter := range namedPathParameters {
				srcValue = append(srcValue, ", vars.Get(", strconv.Quote(namedPathParameter), ")")
			}
			srcValue = append(srcValue, ")")

			switch namedPathField.Desc.Kind() {
			case protoreflect.StringKind:
				f.PrintStringValueAssign(generatedFile, tgtValue, srcValue, namedPathField.Desc.HasPresence())
			case protoreflect.MessageKind:
				f.PrintWrapStringValueAssign(generatedFile, tgtValue, srcValue)
			}
		}
	}
	return nil
}

func (f *ServerGenerator) PrintPathField(generatedFile *protogen.GeneratedFile, pathFields []*protogen.Field) {
	for _, field := range pathFields {
		fieldName := string(field.Desc.Name())
		form := "vars"
		errName := "varErr"

		tgtValue := []any{"req.", field.GoName, " = "}
		tgtErrValue := []any{"req.", field.GoName, ", ", errName, " = "}

		srcValue := []any{"vars.Get(", strconv.Quote(fieldName), ")"}

		goType, pointer := internal.FieldGoType(generatedFile, field)
		if pointer {
			goType = append([]any{"*"}, goType...)
		}

		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
			if pointer {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBool"), fieldName, form, errName)
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
			if pointer {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr[int32]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt[int32]"), fieldName, form, errName)
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
			if pointer {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintPtr[uint32]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint[uint32]"), fieldName, form, errName)
			}
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
			if pointer {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr[int64]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt[int64]"), fieldName, form, errName)
			}
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
			if pointer {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintPtr[uint64]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint[uint64]"), fieldName, form, errName)
			}
		case protoreflect.FloatKind: // float32
			if pointer {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatPtr[float32]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat[float32]"), fieldName, form, errName)
			}
		case protoreflect.DoubleKind: // float64
			if pointer {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatPtr[float64]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat[float64]"), fieldName, form, errName)
			}
		case protoreflect.StringKind: // string
			f.PrintStringValueAssign(generatedFile, tgtValue, srcValue, pointer)
		case protoreflect.EnumKind: // enum int32
			if pointer {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr["+generatedFile.QualifiedGoIdent(goType[1].(protogen.GoIdent))+"]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt["+generatedFile.QualifiedGoIdent(goType[0].(protogen.GoIdent))+"]"), fieldName, form, errName)
			}
		case protoreflect.MessageKind:
			switch field.Message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat64Value"), fieldName, form, errName)
			case "google.protobuf.FloatValue":
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat32Value"), fieldName, form, errName)
			case "google.protobuf.Int64Value":
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt64Value"), fieldName, form, errName)
			case "google.protobuf.UInt64Value":
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint64Value"), fieldName, form, errName)
			case "google.protobuf.Int32Value":
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt32Value"), fieldName, form, errName)
			case "google.protobuf.UInt32Value":
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint32Value"), fieldName, form, errName)
			case "google.protobuf.BoolValue":
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolValue"), fieldName, form, errName)
			case "google.protobuf.StringValue":
				f.PrintWrapStringValueAssign(generatedFile, tgtValue, srcValue)
			}
		}
	}
}

func (f *ServerGenerator) PrintQueryField(generatedFile *protogen.GeneratedFile, queryFields []*protogen.Field) {
	for _, field := range queryFields {
		fieldName := string(field.Desc.Name())

		tgtValue := []any{"req.", field.GoName, " = "}
		tgtErrValue := []any{"req.", field.GoName, ", queryErr = "}
		srcValue := []any{"queries.Get(", strconv.Quote(fieldName), ")"}
		if field.Desc.IsList() {
			srcValue = []any{"queries[", strconv.Quote(fieldName), "]"}
		}

		goType, pointer := internal.FieldGoType(generatedFile, field)
		if pointer {
			goType = append([]any{"*"}, goType...)
		}

		form := "queries"
		errName := "queryErr"

		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
			if field.Desc.IsList() {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolSlice"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolPtr"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBool"), fieldName, form, errName)
				}
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
			if field.Desc.IsList() {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntSlice[int32]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr[int32]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt[int32]"), fieldName, form, errName)
				}
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
			if field.Desc.IsList() {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintSlice[uint32]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintPtr[uint32]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint[uint32]"), fieldName, form, errName)
				}
			}
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
			if field.Desc.IsList() {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntSlice[int64]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr[int64]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt[int64]"), fieldName, form, errName)
				}
			}
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
			if field.Desc.IsList() {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintSlice[uint64]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintPtr[uint64]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint[uint64]"), fieldName, form, errName)
				}
			}
		case protoreflect.FloatKind: // float32
			if field.Desc.IsList() {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatSlice[float32]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatPtr[float32]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat[float32]"), fieldName, form, errName)
				}
			}
		case protoreflect.DoubleKind: // float64
			if field.Desc.IsList() {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatSlice[float64]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatPtr[float64]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat[float64]"), fieldName, form, errName)
				}
			}
		case protoreflect.StringKind: // string
			if field.Desc.IsList() {
				f.PrintStringListAssign(generatedFile, tgtValue, srcValue)
			} else {
				f.PrintStringValueAssign(generatedFile, tgtValue, srcValue, pointer)
			}
		case protoreflect.EnumKind: // enum int32
			if field.Desc.IsList() {
				f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntSlice["+generatedFile.QualifiedGoIdent(goType[1].(protogen.GoIdent))+"]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr["+generatedFile.QualifiedGoIdent(goType[1].(protogen.GoIdent))+"]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt["+generatedFile.QualifiedGoIdent(goType[0].(protogen.GoIdent))+"]"), fieldName, form, errName)
				}
			}
		case protoreflect.MessageKind:
			switch field.Message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
				if field.Desc.IsList() {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat64ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat64Value"), fieldName, form, errName)
				}
			case "google.protobuf.FloatValue":
				if field.Desc.IsList() {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat32ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat32Value"), fieldName, form, errName)
				}
			case "google.protobuf.Int64Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt64ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt64Value"), fieldName, form, errName)
				}
			case "google.protobuf.UInt64Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint64ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint64Value"), fieldName, form, errName)
				}
			case "google.protobuf.Int32Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt32ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt32Value"), fieldName, form, errName)
				}
			case "google.protobuf.UInt32Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint32ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint32Value"), fieldName, form, errName)
				}
			case "google.protobuf.BoolValue":
				if field.Desc.IsList() {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(generatedFile, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolValue"), fieldName, form, errName)
				}
			case "google.protobuf.StringValue":
				if field.Desc.IsList() {
					f.PrintWrapStringListAssign(generatedFile, tgtValue, srcValue)
				} else {
					f.PrintWrapStringValueAssign(generatedFile, tgtValue, srcValue)
				}
			}
		}
	}

}

func (f *ServerGenerator) PrintFieldAssign(generatedFile *protogen.GeneratedFile, tgtValue []any, goType []any, getter protogen.GoIdent, key string, form string, errName string) {
	generatedFile.P(append(append([]any{}, tgtValue...), append(append([]any{internal.ErrorxPackage.Ident("Break"), "["}, goType...), append(append([]any{"](", errName, ")("}, getter), "(", form, ", ", strconv.Quote(key), "))")...)...)...)
}

func (f *ServerGenerator) PrintStringValueAssign(generatedFile *protogen.GeneratedFile, tgtValue []any, srcValue []any, hasPresence bool) {
	if hasPresence {
		generatedFile.P(append(tgtValue, append(append([]any{internal.ProtoPackage.Ident("String"), "("}, srcValue...), ")")...)...)
	} else {
		generatedFile.P(append(tgtValue, srcValue...)...)
	}
}

func (f *ServerGenerator) PrintWrapStringValueAssign(generatedFile *protogen.GeneratedFile, tgtValue []any, srcValue []any) {
	generatedFile.P(append(tgtValue, append(append([]any{internal.WrapperspbPackage.Ident("String"), "("}, srcValue...), ")")...)...)
}

func (f *ServerGenerator) PrintStringListAssign(generatedFile *protogen.GeneratedFile, tgtValue []any, srcValue []any) {
	generatedFile.P(append(tgtValue, srcValue...)...)
}

func (f *ServerGenerator) PrintWrapStringListAssign(generatedFile *protogen.GeneratedFile, tgtValue []any, srcValue []any) {
	generatedFile.P(append(tgtValue, append(append([]any{internal.ProtoxPackage.Ident("WrapStringSlice"), "("}, srcValue...), ")")...)...)
}

func (f *ServerGenerator) PrintEncodeResponseFunc(generatedFile *protogen.GeneratedFile, endpoint *internal.Endpoint, httpRule *internal.HttpRule) error {
	generatedFile.P("func(ctx ", internal.ContextPackage.Ident("Context"), ", w ", internal.HttpPackage.Ident("ResponseWriter"), ", obj any) error {")
	generatedFile.P("resp := obj.(*", endpoint.Output().GoIdent, ")")
	bodyParameter := httpRule.ResponseBody()
	switch bodyParameter {
	case "", "*":
		srcValue := []any{"resp"}
		message := endpoint.Output()
		switch message.Desc.FullName() {
		case "google.api.HttpBody":
			f.PrintGoogleApiHttpBodyEncodeBlock(generatedFile, srcValue)
		case "google.rpc.HttpResponse":
			f.PrintGoogleRpcHttpResponseEncodeBlock(generatedFile, srcValue)
		default:
			f.PrintJsonEncodeBlock(generatedFile, srcValue)
		}
	default:
		bodyField := internal.FindField(bodyParameter, endpoint.Output())
		if bodyField == nil {
			return fmt.Errorf("%s, failed to find body response field %s", endpoint.FullName(), bodyParameter)
		}
		if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.rpc.HttpResponse" {
			return errors.New("google.rpc.HttpResponse can only be used as output to a method")
		}
		srcValue := []any{"resp.Get", bodyField.GoName, "()"}
		if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
			f.PrintGoogleApiHttpBodyEncodeBlock(generatedFile, srcValue)
		} else {
			f.PrintJsonEncodeBlock(generatedFile, srcValue)
		}
	}
	generatedFile.P("return nil")
	return nil
}

func (f *ServerGenerator) PrintGoogleApiHttpBodyEncodeBlock(generatedFile *protogen.GeneratedFile, srcValue []any) {
	generatedFile.P(append(append([]any{"w.Header().Set(", strconv.Quote("Content-Type"), ", "}, srcValue...), ".GetContentType())")...)
	generatedFile.P(append(append([]any{"for _, src := range "}, srcValue...), ".GetExtensions() {")...)
	generatedFile.P("dst, err := ", internal.AnypbPackage.Ident("UnmarshalNew"), "(src, ", internal.ProtoPackage.Ident("UnmarshalOptions"), "{})")
	generatedFile.P("if err != nil {")
	generatedFile.P("return err")
	generatedFile.P("}")
	generatedFile.P("metadata, ok := dst.(*", internal.StructpbPackage.Ident("Struct"), ")")
	generatedFile.P("if !ok {")
	generatedFile.P("continue")
	generatedFile.P("}")
	generatedFile.P("for key, value := range metadata.GetFields() {")
	generatedFile.P("w.Header().Add(key, string(", internal.ErrorxPackage.Ident("Ignore"), "(", internal.JsonxPackage.Ident("Marshal"), "(value))))")
	generatedFile.P("}")
	generatedFile.P("}")
	generatedFile.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
	generatedFile.P(append(append([]any{"if ", "_, err := w.Write("}, srcValue...), ".GetData())", "; err != nil {")...)
	generatedFile.P("return err")
	generatedFile.P("}")
}

func (f *ServerGenerator) PrintGoogleRpcHttpResponseEncodeBlock(generatedFile *protogen.GeneratedFile, srcValue []any) {
	generatedFile.P(append(append([]any{"for _, header := range "}, srcValue...), ".GetHeaders() {")...)
	generatedFile.P("w.Header().Add(header.Key, header.Value)")
	generatedFile.P("}")
	generatedFile.P(append(append([]any{"w.WriteHeader(int("}, srcValue...), ".GetStatus()))")...)
	generatedFile.P(append(append([]any{"if ", "_, err := w.Write("}, srcValue...), ".GetBody())", "; err != nil {")...)
	generatedFile.P("return err")
	generatedFile.P("}")
}

func (f *ServerGenerator) PrintJsonEncodeBlock(generatedFile *protogen.GeneratedFile, srcValue []any) {
	generatedFile.P("w.Header().Set(", strconv.Quote("Content-Type"), ", ", strconv.Quote(internal.JsonContentType), ")")
	generatedFile.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
	generatedFile.P(append(append([]any{"if err := ", internal.JsonxPackage.Ident("NewEncoder"), "(w).Encode("}, srcValue...), "); err != nil {")...)
	generatedFile.P("return err")
	generatedFile.P("}")
}
