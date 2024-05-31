package http

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo/generator/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

type ServerGenerator struct{}

func (f *ServerGenerator) GenerateTransports(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.HttpServerTransportsName(), " interface {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.Name(), "() *", internal.HttpTransportPackage.Ident("Server"))
	}
	g.P("}")
	g.P()
	return nil
}

func (f *ServerGenerator) GenerateImplementedTransports(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedHttpServerTransportsName(), " struct {")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), " *", internal.HttpTransportPackage.Ident("Server"))
	}
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (t *", service.UnexportedHttpServerTransportsName(), ") ", endpoint.Name(), "() *", internal.HttpTransportPackage.Ident("Server"), "{")
		g.P("return t.", endpoint.UnexportedName())
		g.P("}")
		g.P()
	}

	g.P("func New", service.HttpServerTransportsName(), "(endpoints ", service.EndpointsName(), ", serverOptions ...", internal.HttpTransportPackage.Ident("ServerOption"), ") ", service.HttpServerTransportsName(), " {")
	g.P("return &", service.UnexportedHttpServerTransportsName(), "{")
	for _, endpoint := range service.Endpoints {
		g.P(endpoint.UnexportedName(), ": ", internal.HttpTransportPackage.Ident("NewServer"), "(")
		g.P("endpoints.", endpoint.Name(), "(), ")
		if err := f.PrintDecodeRequestFunc(g, endpoint); err != nil {
			return err
		}
		if err := f.PrintEncodeResponseFunc(g, endpoint, endpoint.HttpRule()); err != nil {
			return err
		}
		g.P("},")
		g.P("append([]", internal.HttpTransportPackage.Ident("ServerOption"), "{")
		g.P(internal.HttpTransportPackage.Ident("ServerBefore"), "(func(ctx ", internal.ContextPackage.Ident("Context"), ", request *", internal.HttpPackage.Ident("Request"), ") ", internal.ContextPackage.Ident("Context"), " {")
		g.P("return ", internal.EndpointxPackage.Ident("InjectName"), "(ctx, ", strconv.Quote(endpoint.FullName()), ")")
		g.P("}),")
		g.P(internal.HttpTransportPackage.Ident("ServerBefore"), "(func(ctx ", internal.ContextPackage.Ident("Context"), ", request *", internal.HttpPackage.Ident("Request"), ") ", internal.ContextPackage.Ident("Context"), " {")
		g.P("return ", internal.TransportxPackage.Ident("InjectName"), "(ctx, ", internal.TransportxPackage.Ident("HttpServer"), ")")
		g.P("}),")
		g.P("}, serverOptions...)...,")
		g.P("),")
	}
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func (f *ServerGenerator) GenerateServer(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("func New", service.HttpServerHandlerName(), "(endpoints ", service.HttpServerTransportsName(), ") ", internal.HttpPackage.Ident("Handler"), " {")
	g.P("router := ", internal.MuxPackage.Ident("NewRouter"), "()")
	for _, endpoint := range service.Endpoints {
		httpRule := endpoint.HttpRule()
		// 调整路径，来适应 github.com/gorilla/mux 路由规则
		path, _, _, _ := httpRule.RegularizePath(httpRule.Path())
		g.P("router.NewRoute().Name(", strconv.Quote(endpoint.FullName()), ").Methods(", strconv.Quote(httpRule.Method()), ").Path(", strconv.Quote(path), ").Handler(endpoints.", endpoint.Name(), "())")
	}
	g.P("return router")
	g.P("}")
	g.P()
	return nil
}

func (f *ServerGenerator) PrintDecodeRequestFunc(
	g *protogen.GeneratedFile, endpoint *internal.Endpoint,
) error {
	g.P("func(ctx ", internal.ContextPackage.Ident("Context"), ", r *", internal.HttpPackage.Ident("Request"), ") (any, error) {")
	g.P("req := &", endpoint.InputGoIdent(), "{}")

	bodyMessage, bodyField, namedPathFields, pathFields, queryFields, err := endpoint.ParseParameters()
	if err != nil {
		return err
	}

	if bodyMessage != nil {
		switch bodyMessage.Desc.FullName() {
		case "google.api.HttpBody":
			f.PrintGoogleApiHttpBodyDecodeBlock(g, []any{"req"}, []any{"r.Body"})
		default:
			f.PrintDecodeBlock(g, internal.JsonxPackage.Ident("NewDecoder"), []any{"req"}, []any{"r.Body"})
		}
	} else if bodyField != nil {
		if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
			g.P("req.", bodyField.GoName, " = &", bodyField.Message.GoIdent, "{}")
			tgtValue := []any{"req.", bodyField.GoName}
			srcValue := []any{"r.Body"}
			f.PrintGoogleApiHttpBodyDecodeBlock(g, tgtValue, srcValue)
		} else {
			tgtValue := []any{"&req.", bodyField.GoName}
			srcValue := []any{"r.Body"}
			f.PrintDecodeBlock(g, internal.JsonxPackage.Ident("NewDecoder"), tgtValue, srcValue)
		}
	}

	if len(namedPathFields)+len(pathFields) > 0 {
		g.P("vars := ", internal.UrlxPackage.Ident("FormFromMap"), "(", internal.MuxPackage.Ident("Vars"), "(r)", ")")
		g.P("var varErr error")
		if err := f.PrintNamedPathField(g, namedPathFields, endpoint.HttpRule()); err != nil {
			return err
		}
		f.PrintPathField(g, pathFields)
		g.P("if varErr != nil {")
		g.P("return nil, varErr")
		g.P("}")
	}

	if len(queryFields) > 0 {
		g.P("queries := r.URL.Query()")
		g.P("var queryErr error")
		f.PrintQueryField(g, queryFields)
		g.P("if queryErr != nil {")
		g.P("return nil, queryErr")
		g.P("}")
	}

	g.P("return req, nil")
	g.P("},")
	return nil
}

func (f *ServerGenerator) PrintGoogleApiHttpBodyDecodeBlock(g *protogen.GeneratedFile, tgtValue []any, srcValue []any) {
	g.P(append(append([]any{"body, err := ", internal.IOPackage.Ident("ReadAll"), "("}, srcValue...), []any{")"}...)...)
	g.P("if err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P(append(append([]any{}, tgtValue...), []any{".Data = body"}...)...)
	g.P(append(append([]any{}, tgtValue...), []any{".ContentType = r.Header.Get(", strconv.Quote(internal.ContentTypeKey), ")"}...)...)
}

func (f *ServerGenerator) PrintDecodeBlock(g *protogen.GeneratedFile, decoder protogen.GoIdent, tgtValue []any, srcValue []any) {
	g.P(append(append(append(append([]any{"if err := ", decoder, "("}, srcValue...), []any{").Decode("}...), tgtValue...), []any{"); err != nil {"}...)...)
	g.P("return nil, err")
	g.P("}")
}

func (f *ServerGenerator) PrintNamedPathField(g *protogen.GeneratedFile, namedPathFields []*protogen.Field, httpRule *internal.HttpRule) error {
	for i, namedPathField := range namedPathFields {
		fullFieldName := internal.FullFieldName(namedPathFields[:i+1])
		if i < len(namedPathFields)-1 {
			g.P("if req.", fullFieldName, " == nil {")
			g.P("req.", fullFieldName, " = &", namedPathField.Message.GoIdent, "{}")
			g.P("}")
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
				f.PrintStringValueAssign(g, tgtValue, srcValue, namedPathField.Desc.HasPresence())
			case protoreflect.MessageKind:
				f.PrintWrapStringValueAssign(g, tgtValue, srcValue)
			}
		}
	}
	return nil
}

func (f *ServerGenerator) PrintPathField(g *protogen.GeneratedFile, pathFields []*protogen.Field) {
	for _, field := range pathFields {
		fieldName := string(field.Desc.Name())
		form := "vars"
		errName := "varErr"

		tgtValue := []any{"req.", field.GoName, " = "}
		tgtErrValue := []any{"req.", field.GoName, ", ", errName, " = "}

		srcValue := []any{"vars.Get(", strconv.Quote(fieldName), ")"}

		goType, pointer := internal.FieldGoType(g, field)
		if pointer {
			goType = append([]any{"*"}, goType...)
		}

		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolPtr"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBool"), fieldName, form, errName)
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr[int32]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt[int32]"), fieldName, form, errName)
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintPtr[uint32]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint[uint32]"), fieldName, form, errName)
			}
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr[int64]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt[int64]"), fieldName, form, errName)
			}
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintPtr[uint64]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint[uint64]"), fieldName, form, errName)
			}
		case protoreflect.FloatKind: // float32
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatPtr[float32]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat[float32]"), fieldName, form, errName)
			}
		case protoreflect.DoubleKind: // float64
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatPtr[float64]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat[float64]"), fieldName, form, errName)
			}
		case protoreflect.StringKind: // string
			f.PrintStringValueAssign(g, tgtValue, srcValue, pointer)
		case protoreflect.EnumKind: // enum int32
			if pointer {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr["+g.QualifiedGoIdent(goType[1].(protogen.GoIdent))+"]"), fieldName, form, errName)
			} else {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt["+g.QualifiedGoIdent(goType[0].(protogen.GoIdent))+"]"), fieldName, form, errName)
			}
		case protoreflect.MessageKind:
			switch field.Message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat64Value"), fieldName, form, errName)
			case "google.protobuf.FloatValue":
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat32Value"), fieldName, form, errName)
			case "google.protobuf.Int64Value":
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt64Value"), fieldName, form, errName)
			case "google.protobuf.UInt64Value":
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint64Value"), fieldName, form, errName)
			case "google.protobuf.Int32Value":
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt32Value"), fieldName, form, errName)
			case "google.protobuf.UInt32Value":
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint32Value"), fieldName, form, errName)
			case "google.protobuf.BoolValue":
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolValue"), fieldName, form, errName)
			case "google.protobuf.StringValue":
				f.PrintWrapStringValueAssign(g, tgtValue, srcValue)
			}
		}
	}
}

func (f *ServerGenerator) PrintQueryField(g *protogen.GeneratedFile, queryFields []*protogen.Field) {
	for _, field := range queryFields {
		fieldName := string(field.Desc.Name())

		tgtValue := []any{"req.", field.GoName, " = "}
		tgtErrValue := []any{"req.", field.GoName, ", queryErr = "}
		srcValue := []any{"queries.Get(", strconv.Quote(fieldName), ")"}
		if field.Desc.IsList() {
			srcValue = []any{"queries[", strconv.Quote(fieldName), "]"}
		}

		goType, pointer := internal.FieldGoType(g, field)
		if pointer {
			goType = append([]any{"*"}, goType...)
		}

		form := "queries"
		errName := "queryErr"

		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolSlice"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolPtr"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBool"), fieldName, form, errName)
				}
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntSlice[int32]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr[int32]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt[int32]"), fieldName, form, errName)
				}
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintSlice[uint32]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintPtr[uint32]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint[uint32]"), fieldName, form, errName)
				}
			}
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntSlice[int64]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr[int64]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt[int64]"), fieldName, form, errName)
				}
			}
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintSlice[uint64]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUintPtr[uint64]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint[uint64]"), fieldName, form, errName)
				}
			}
		case protoreflect.FloatKind: // float32
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatSlice[float32]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatPtr[float32]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat[float32]"), fieldName, form, errName)
				}
			}
		case protoreflect.DoubleKind: // float64
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatSlice[float64]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloatPtr[float64]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat[float64]"), fieldName, form, errName)
				}
			}
		case protoreflect.StringKind: // string
			if field.Desc.IsList() {
				f.PrintStringListAssign(g, tgtValue, srcValue)
			} else {
				f.PrintStringValueAssign(g, tgtValue, srcValue, pointer)
			}
		case protoreflect.EnumKind: // enum int32
			if field.Desc.IsList() {
				f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntSlice["+g.QualifiedGoIdent(goType[1].(protogen.GoIdent))+"]"), fieldName, form, errName)
			} else {
				if pointer {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetIntPtr["+g.QualifiedGoIdent(goType[1].(protogen.GoIdent))+"]"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt["+g.QualifiedGoIdent(goType[0].(protogen.GoIdent))+"]"), fieldName, form, errName)
				}
			}
		case protoreflect.MessageKind:
			switch field.Message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat64ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat64Value"), fieldName, form, errName)
				}
			case "google.protobuf.FloatValue":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat32ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetFloat32Value"), fieldName, form, errName)
				}
			case "google.protobuf.Int64Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt64ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt64Value"), fieldName, form, errName)
				}
			case "google.protobuf.UInt64Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint64ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint64Value"), fieldName, form, errName)
				}
			case "google.protobuf.Int32Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt32ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetInt32Value"), fieldName, form, errName)
				}
			case "google.protobuf.UInt32Value":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint32ValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetUint32Value"), fieldName, form, errName)
				}
			case "google.protobuf.BoolValue":
				if field.Desc.IsList() {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolValueSlice"), fieldName, form, errName)
				} else {
					f.PrintFieldAssign(g, tgtErrValue, goType, internal.UrlxPackage.Ident("GetBoolValue"), fieldName, form, errName)
				}
			case "google.protobuf.StringValue":
				if field.Desc.IsList() {
					f.PrintWrapStringListAssign(g, tgtValue, srcValue)
				} else {
					f.PrintWrapStringValueAssign(g, tgtValue, srcValue)
				}
			}
		}
	}

}

func (f *ServerGenerator) PrintFieldAssign(g *protogen.GeneratedFile, tgtValue []any, goType []any, getter protogen.GoIdent, key string, form string, errName string) {
	g.P(append(append([]any{}, tgtValue...), append(append([]any{internal.ErrorxPackage.Ident("Break"), "["}, goType...), append(append([]any{"](", errName, ")("}, getter), "(", form, ", ", strconv.Quote(key), "))")...)...)...)
}

func (f *ServerGenerator) PrintStringValueAssign(g *protogen.GeneratedFile, tgtValue []any, srcValue []any, hasPresence bool) {
	if hasPresence {
		g.P(append(tgtValue, append(append([]any{internal.ProtoPackage.Ident("String"), "("}, srcValue...), ")")...)...)
	} else {
		g.P(append(tgtValue, srcValue...)...)
	}
}

func (f *ServerGenerator) PrintWrapStringValueAssign(g *protogen.GeneratedFile, tgtValue []any, srcValue []any) {
	g.P(append(tgtValue, append(append([]any{internal.WrapperspbPackage.Ident("String"), "("}, srcValue...), ")")...)...)
}

func (f *ServerGenerator) PrintStringListAssign(g *protogen.GeneratedFile, tgtValue []any, srcValue []any) {
	g.P(append(tgtValue, srcValue...)...)
}

func (f *ServerGenerator) PrintWrapStringListAssign(g *protogen.GeneratedFile, tgtValue []any, srcValue []any) {
	g.P(append(tgtValue, append(append([]any{internal.ProtoxPackage.Ident("WrapStringSlice"), "("}, srcValue...), ")")...)...)
}

func (f *ServerGenerator) PrintEncodeResponseFunc(g *protogen.GeneratedFile, endpoint *internal.Endpoint, httpRule *internal.HttpRule) error {
	g.P("func(ctx ", internal.ContextPackage.Ident("Context"), ", w ", internal.HttpPackage.Ident("ResponseWriter"), ", obj any) error {")
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
	return nil
}

func (f *ServerGenerator) PrintGoogleApiHttpBodyEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
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

func (f *ServerGenerator) PrintJsonEncodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P("w.Header().Set(", strconv.Quote("Content-Type"), ", ", strconv.Quote(internal.JsonContentType), ")")
	g.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
	g.P(append(append([]any{"if err := ", internal.JsonxPackage.Ident("NewEncoder"), "(w).Encode("}, srcValue...), "); err != nil {")...)
	g.P("return err")
	g.P("}")
}