package generator

import (
	"fmt"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
	"strings"
)

type ClientGenerator struct{}

func (f *ClientGenerator) GenerateClient(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("type ", service.UnexportedHTTPClientName(), " struct {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.UnexportedName(), " ", internal.EndpointPackage.Ident("Endpoint"))
	}
	generatedFile.P("}")
	generatedFile.P()
	for _, endpoint := range service.Endpoints {
		generatedFile.P("func (c *", service.UnexportedHTTPClientName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		generatedFile.P("rep, err := c.", endpoint.UnexportedName(), "(ctx, request)")
		generatedFile.P("if err != nil {")
		generatedFile.P("return nil, err")
		generatedFile.P("}")
		generatedFile.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
		generatedFile.P("}")
		generatedFile.P()
	}
	return nil
}

func (f *ClientGenerator) GenerateNewClient(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("func New", service.HTTPClientName(), "(")
	generatedFile.P("scheme   string,")
	generatedFile.P("instance string,")
	generatedFile.P("mdw []", internal.EndpointPackage.Ident("Middleware"), ",")
	generatedFile.P("opts ...", internal.HttpTransportPackage.Ident("ClientOption"), ",")
	generatedFile.P(") interface {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error)")
	}
	generatedFile.P("} {")
	generatedFile.P("router := ", internal.MuxPackage.Ident("NewRouter"), "()")
	for _, endpoint := range service.Endpoints {
		httpRule := endpoint.HttpRule()
		// 调整路径，来适应 github.com/gorilla/mux 路由规则
		path, _, _, _ := httpRule.RegularizePath(httpRule.Path())
		generatedFile.P("router.NewRoute().")
		generatedFile.P("Name(", strconv.Quote(endpoint.FullName()), ").")
		generatedFile.P("Methods(", strconv.Quote(httpRule.Method()), ").")
		generatedFile.P("Path(", strconv.Quote(path), ")")
	}
	generatedFile.P("return &", service.UnexportedHTTPClientName(), "{")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.UnexportedName(), ":    ", internal.EndpointxPackage.Ident("Chain"), "(")
		generatedFile.P(internal.HttpTransportPackage.Ident("NewExplicitClient"), "(")
		if err := f.PrintEncodeRequestFunc(generatedFile, endpoint); err != nil {
			return err
		}
		if err := f.PrintDecodeResponseFunc(generatedFile); err != nil {
			return err
		}
		generatedFile.P("opts...,")
		generatedFile.P(").Endpoint(),")
		generatedFile.P("mdw...),")
	}
	generatedFile.P("}")
	generatedFile.P("}")
	generatedFile.P()
	return nil
}

func (f *ClientGenerator) PrintEncodeRequestFunc(generatedFile *protogen.GeneratedFile, endpoint *internal.Endpoint) error {
	httpRule := endpoint.HttpRule()
	generatedFile.P("func(ctx context.Context, obj interface{}) (*", internal.HttpPackage.Ident("Request"), ", error) {")
	generatedFile.P("if obj == nil {")
	generatedFile.P("return nil, ", internal.ErrorsPackage.Ident("New"), "(", strconv.Quote("request object is nil"), ")")
	generatedFile.P("}")
	generatedFile.P("req, ok := obj.(*", endpoint.InputGoIdent(), ")")
	generatedFile.P("if !ok {")
	generatedFile.P("return nil, ", internal.FmtPackage.Ident("Errorf"), "(", strconv.Quote("invalid request object type, %T"), ", obj)")
	generatedFile.P("}")
	generatedFile.P("_ = req")

	bodyMessage, bodyField, namedPathFields, pathFields, queryFields, err := endpoint.ParseParameters()
	if err != nil {
		return err
	}

	generatedFile.P("var body ", internal.IOPackage.Ident("Reader"))
	if bodyMessage != nil {
		switch bodyMessage.Desc.FullName() {
		case "google.api.HttpBody":
			f.PrintReaderBlock(generatedFile, internal.BytesPackage, []any{"body"}, []any{"req.GetData()"})
			generatedFile.P("contentType := req.GetContentType()")
		case "google.rpc.HttpRequest":
			f.PrintGoogleRpcHttpRequest(generatedFile)
			generatedFile.P("return r, nil")
			generatedFile.P("},")
			return nil
		default:
			generatedFile.P("var bodyBuf bytes.Buffer")
			encoder := internal.JsonPackage.Ident("NewEncoder")
			f.PrintEncodeBlock(generatedFile, encoder, []any{"&bodyBuf"}, []any{"req"})
			generatedFile.P("body = &bodyBuf")
			generatedFile.P("contentType := ", strconv.Quote(internal.JsonContentType))
		}
	} else if bodyField != nil {
		if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
			f.PrintReaderBlock(generatedFile, internal.BytesPackage, []any{"body"}, []any{"req.Get", bodyField.GoName, "()", ".GetData()"})
			generatedFile.P("contentType := req.Get", bodyField.GoName, "()", ".GetContentType()")
		} else {
			generatedFile.P("var bodyBuf bytes.Buffer")
			encoder := internal.JsonPackage.Ident("NewEncoder")
			tgtValue := []any{"&bodyBuf"}
			srcValue := []any{"req.Get", bodyField.GoName, "()"}
			f.PrintEncodeBlock(generatedFile, encoder, tgtValue, srcValue)
			generatedFile.P("body = &bodyBuf")
			generatedFile.P("contentType := ", strconv.Quote(internal.JsonContentType))
		}
	}

	generatedFile.P("var pairs []string")
	if len(namedPathFields) > 0 {
		f.PrintNamedPathField(generatedFile, namedPathFields, httpRule)
	}

	if len(pathFields) > 0 {
		for _, pathField := range pathFields {
			field := pathField
			srcValue := []any{"req.", field.GoName}
			isOptional := field.Desc.HasOptionalKeyword()
			if isOptional || pathField.Desc.Kind() == protoreflect.MessageKind {
				generatedFile.P(append(append([]any{"if "}, srcValue...), " == nil {")...)
				generatedFile.P("return nil, ", internal.FmtPackage.Ident("Errorf"), "(", strconv.Quote("%s is nil"), ", ", strconv.Quote(strings.Join(convx.ToStringSlice(srcValue), "")), ")")
				generatedFile.P("}")
			}
		}

		pairs := []any{"pairs = append(pairs"}
		for _, field := range pathFields {
			pairs = append(append(pairs, ",", strconv.Quote(string(field.Desc.Name())), ","), f.PathFieldFormat(field)...)
		}
		pairs = append(pairs, ")")
		generatedFile.P(pairs...)
	}

	generatedFile.P("path, err := router.Get(", strconv.Quote(endpoint.FullName()), ").URLPath(pairs...)")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")

	generatedFile.P("queries := ", internal.UrlPackage.Ident("Values"), "{}")
	if len(queryFields) > 0 {
		for _, field := range queryFields {
			f.PrintQueryField(generatedFile, field)
		}
	}

	generatedFile.P("target := &", internal.UrlPackage.Ident("URL"), "{")
	generatedFile.P("Scheme:   scheme,")
	generatedFile.P("Host:     instance,")
	generatedFile.P("Path:     path.Path,")
	generatedFile.P("RawQuery: queries.Encode(),")
	generatedFile.P("}")

	generatedFile.P("r, err := ", internal.HttpPackage.Ident("NewRequestWithContext"), "(ctx, ", strconv.Quote(httpRule.Method()), ", target.String(), body)")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	if bodyMessage != nil || bodyField != nil {
		generatedFile.P("r.Header.Set(", strconv.Quote(internal.ContentTypeKey), ", contentType)")
	}
	generatedFile.P("return r, nil")
	generatedFile.P("},")
	return nil
}

func (f *ClientGenerator) PrintGoogleRpcHttpRequest(generatedFile *protogen.GeneratedFile) {
	f.PrintReaderBlock(generatedFile, internal.BytesPackage, []any{"body"}, []any{"req.GetBody()"})
	generatedFile.P("r, err := ", internal.HttpPackage.Ident("NewRequestWithContext"), "(ctx, req.GetMethod(), req.GetUri(), body)")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P("for _, header := range req.GetHeaders() {")
	generatedFile.P("r.Header.Add(header.GetKey(), header.GetValue())")
	generatedFile.P("}")

}

func (f *ClientGenerator) PrintReaderBlock(generatedFile *protogen.GeneratedFile, readerPkg protogen.GoImportPath, tgtValue []any, srcValue []any) {
	generatedFile.P(append(append(append(append([]any{}, tgtValue...), []any{" = ", readerPkg.Ident("NewReader"), "("}...), srcValue...), ")")...)
}

func (f *ClientGenerator) PrintEncodeBlock(generatedFile *protogen.GeneratedFile, encoder protogen.GoIdent, tgtValue []any, srcValue []any) {
	generatedFile.P(append(append(append(append([]any{"if err := ", encoder, "("}, tgtValue...), []any{").Encode("}...), srcValue...), []any{"); err != nil {"}...)...)
	generatedFile.P("return nil, err")
	generatedFile.P("}")
}

func (f *ClientGenerator) PrintNamedPathField(generatedFile *protogen.GeneratedFile, namedPathFields []*protogen.Field, httpRule *internal.HttpRule) {
	fullFieldGetterName := internal.FullFieldGetterName(namedPathFields)
	_, _, _, namedPathParameters := httpRule.RegularizePath(httpRule.Path())
	lastField := namedPathFields[len(namedPathFields)-1]
	switch lastField.Desc.Kind() {
	case protoreflect.StringKind:
		if lastField.Desc.HasOptionalKeyword() {
			generatedFile.P("namedPathParameter := req.", fullFieldGetterName)
		} else {
			generatedFile.P("namedPathParameter := req.", fullFieldGetterName)
		}
	case protoreflect.MessageKind:
		generatedFile.P("namedPathParameter := req.", fullFieldGetterName, ".GetValue()")
	}

	generatedFile.P("namedPathValues := ", internal.StringsPackage.Ident("Split"), "(namedPathParameter, ", strconv.Quote("/"), ")")
	generatedFile.P("if len(namedPathValues) != ", strconv.Itoa(len(namedPathParameters)*2), " {")
	generatedFile.P("return nil, ", internal.FmtPackage.Ident("Errorf"), "(", strconv.Quote("invalid named path parameter, %s"), ", namedPathParameter)")
	generatedFile.P("}")

	pairs := []any{"pairs = append(pairs"}
	for i, parameter := range namedPathParameters {
		pairs = append(pairs, ",", strconv.Quote(parameter), ",", fmt.Sprintf("namedPathValues[%d]", i*2+1))
	}
	pairs = append(pairs, ")")
	generatedFile.P(pairs...)
}

func (f *ClientGenerator) PrintQueryField(generatedFile *protogen.GeneratedFile, field *protogen.Field) {
	srcValue := []any{"req.", field.GoName}
	isOptional := field.Desc.HasOptionalKeyword()
	fieldName := string(field.Desc.Name())
	switch field.Desc.Kind() {
	case protoreflect.BoolKind: // bool
		if field.Desc.IsList() {
			f.PrintQueryList(generatedFile, fieldName, srcValue, f.BoolKindFormat([]any{"item"}, false))
		} else {
			f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.BoolKindFormat(srcValue, isOptional), isOptional)
		}
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
		if field.Desc.IsList() {
			f.PrintQueryList(generatedFile, fieldName, srcValue, f.Int32KindFormat([]any{"item"}, false))
		} else {
			f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.Int32KindFormat(srcValue, isOptional), isOptional)
		}
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
		if field.Desc.IsList() {
			f.PrintQueryList(generatedFile, fieldName, srcValue, f.Uint32KindFormat([]any{"item"}, false))
		} else {
			f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.Uint32KindFormat(srcValue, isOptional), isOptional)
		}
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
		if field.Desc.IsList() {
			f.PrintQueryList(generatedFile, fieldName, srcValue, f.Int64KindFormat([]any{"item"}, false))
		} else {
			f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.Int64KindFormat(srcValue, isOptional), isOptional)
		}
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
		if field.Desc.IsList() {
			f.PrintQueryList(generatedFile, fieldName, srcValue, f.Uint64KindFormat([]any{"item"}, false))
		} else {
			f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.Uint64KindFormat(srcValue, isOptional), isOptional)
		}
	case protoreflect.FloatKind: // float32
		if field.Desc.IsList() {
			f.PrintQueryList(generatedFile, fieldName, srcValue, f.FloatKindFormat([]any{"item"}, false))
		} else {
			f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.FloatKindFormat(srcValue, isOptional), isOptional)
		}
	case protoreflect.DoubleKind: // float64
		if field.Desc.IsList() {
			f.PrintQueryList(generatedFile, fieldName, srcValue, f.DoubleKindFormat([]any{"item"}, false))
		} else {
			f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.DoubleKindFormat(srcValue, isOptional), isOptional)
		}
	case protoreflect.StringKind: // string
		if field.Desc.IsList() {
			f.PrintQueryList(generatedFile, fieldName, srcValue, f.StringKindFormat([]any{"item"}, false))
		} else {
			f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.StringKindFormat(srcValue, isOptional), isOptional)
		}
	case protoreflect.EnumKind: // enum
		if field.Desc.IsList() {
			f.PrintQueryList(generatedFile, fieldName, srcValue, f.Int32KindFormat([]any{"item"}, false))
		} else {
			f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.Int32KindFormat(srcValue, isOptional), isOptional)
		}
	case protoreflect.MessageKind:
		message := field.Message
		switch message.Desc.FullName() {
		case "google.protobuf.DoubleValue":
			if field.Desc.IsList() {
				f.PrintQueryList(generatedFile, fieldName, srcValue, f.WrapDoubleFormat([]any{"item"}))
			} else {
				f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.WrapDoubleFormat(srcValue), true)
			}
		case "google.protobuf.FloatValue":
			if field.Desc.IsList() {
				f.PrintQueryList(generatedFile, fieldName, srcValue, f.WrapFloatFormat([]any{"item"}))
			} else {
				f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.WrapFloatFormat(srcValue), true)
			}
		case "google.protobuf.Int64Value":
			if field.Desc.IsList() {
				f.PrintQueryList(generatedFile, fieldName, srcValue, f.WrapInt64Format([]any{"item"}))
			} else {
				f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.WrapInt64Format(srcValue), true)
			}
		case "google.protobuf.UInt64Value":
			if field.Desc.IsList() {
				f.PrintQueryList(generatedFile, fieldName, srcValue, f.WrapUint64Format([]any{"item"}))
			} else {
				f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.WrapUint64Format(srcValue), true)
			}
		case "google.protobuf.Int32Value":
			if field.Desc.IsList() {
				f.PrintQueryList(generatedFile, fieldName, srcValue, f.WrapInt32Format([]any{"item"}))
			} else {
				f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.WrapInt32Format(srcValue), true)
			}
		case "google.protobuf.UInt32Value":
			if field.Desc.IsList() {
				f.PrintQueryList(generatedFile, fieldName, srcValue, f.WrapUint32Format([]any{"item"}))
			} else {
				f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.WrapUint32Format(srcValue), true)
			}
		case "google.protobuf.BoolValue":
			if field.Desc.IsList() {
				f.PrintQueryList(generatedFile, fieldName, srcValue, f.WrapBoolFormat([]any{"item"}))
			} else {
				f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.WrapBoolFormat(srcValue), true)
			}
		case "google.protobuf.StringValue":
			if field.Desc.IsList() {
				f.PrintQueryList(generatedFile, fieldName, srcValue, f.WrapStringFormat([]any{"item"}))
			} else {
				f.PrintQueryOptional(generatedFile, fieldName, srcValue, f.WrapStringFormat(srcValue), true)
			}
		}
	}
}

func (f *ClientGenerator) PrintQueryOptional(generatedFile *protogen.GeneratedFile, fieldName string, srcValue []any, format []any, isOptional bool) {
	if isOptional {
		generatedFile.P(append(append([]any{"if "}, srcValue...), " == nil {")...)
	}
	generatedFile.P(append(append([]any{"queries.Add(", strconv.Quote(fieldName), ","}, format...), []any{")"}...)...)
	if isOptional {
		generatedFile.P("}")
	}
}

func (f *ClientGenerator) PrintQueryList(generatedFile *protogen.GeneratedFile, fieldName string, srcValue []any, format []any) {
	generatedFile.P(append(append([]any{"for _, item := range "}, srcValue...), "{")...)
	generatedFile.P(append(append([]any{"queries.Add(", strconv.Quote(fieldName), ", "}, format...), []any{")"}...)...)
	generatedFile.P("}")
}

func (f *ClientGenerator) PathFieldFormat(field *protogen.Field) []any {
	srcValue := []any{"req.", field.GoName}
	switch field.Desc.Kind() {
	case protoreflect.BoolKind: // bool
		return f.BoolKindFormat(srcValue, field.Desc.HasOptionalKeyword())
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
		return f.Int32KindFormat(srcValue, field.Desc.HasOptionalKeyword())
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
		return f.Uint32KindFormat(srcValue, field.Desc.HasOptionalKeyword())
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
		return f.Int64KindFormat(srcValue, field.Desc.HasOptionalKeyword())
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
		return f.Uint64KindFormat(srcValue, field.Desc.HasOptionalKeyword())
	case protoreflect.FloatKind: // float32
		return f.FloatKindFormat(srcValue, field.Desc.HasOptionalKeyword())
	case protoreflect.DoubleKind: // float64
		return f.DoubleKindFormat(srcValue, field.Desc.HasOptionalKeyword())
	case protoreflect.StringKind: // string
		return f.StringKindFormat(srcValue, field.Desc.HasOptionalKeyword())
	case protoreflect.EnumKind: // enum
		return f.Int32KindFormat(srcValue, field.Desc.HasOptionalKeyword())
	case protoreflect.MessageKind:
		switch field.Message.Desc.FullName() {
		case "google.protobuf.DoubleValue":
			return f.WrapDoubleFormat(srcValue)
		case "google.protobuf.FloatValue":
			return f.WrapFloatFormat(srcValue)
		case "google.protobuf.Int64Value":
			return f.WrapInt64Format(srcValue)
		case "google.protobuf.UInt64Value":
			return f.WrapUint64Format(srcValue)
		case "google.protobuf.Int32Value":
			return f.WrapInt32Format(srcValue)
		case "google.protobuf.UInt32Value":
			return f.WrapUint32Format(srcValue)
		case "google.protobuf.BoolValue":
			return f.WrapBoolFormat(srcValue)
		case "google.protobuf.StringValue":
			return f.WrapStringFormat(srcValue)
		}
	}
	return nil
}

func (f *ClientGenerator) BytesKindFormat(srcValue []any, isOptional bool) []any {
	return append(append([]any{}, srcValue...), []any{}...)
}

func (f *ClientGenerator) StringKindFormat(srcValue []any, isOptional bool) []any {
	format := []any{}
	if isOptional {
		format = append(format, "*")
	}
	return append(append(format, srcValue...), []any{}...)
}

func (f *ClientGenerator) DoubleKindFormat(srcValue []any, isOptional bool) []any {
	format := []any{internal.StrconvPackage.Ident("FormatFloat"), "("}
	if isOptional {
		format = append(format, "*")
	}
	return append(append(format, srcValue...), []any{", 'f', -1, 64)"}...)
}

func (f *ClientGenerator) FloatKindFormat(srcValue []any, isOptional bool) []any {
	format := []any{internal.StrconvPackage.Ident("FormatFloat"), "(float64("}
	if isOptional {
		format = append(format, "*")
	}
	return append(append(format, srcValue...), []any{"), 'f', -1, 32)"}...)
}

func (f *ClientGenerator) Uint64KindFormat(srcValue []any, isOptional bool) []any {
	format := []any{internal.StrconvPackage.Ident("FormatUint"), "("}
	if isOptional {
		format = append(format, "*")
	}
	return append(append(format, srcValue...), []any{", 10)"}...)
}

func (f *ClientGenerator) Int64KindFormat(srcValue []any, isOptional bool) []any {
	format := []any{internal.StrconvPackage.Ident("FormatInt"), "("}
	if isOptional {
		format = append(format, "*")
	}
	return append(append(format, srcValue...), []any{", 10)"}...)
}

func (f *ClientGenerator) Uint32KindFormat(srcValue []any, isOptional bool) []any {
	format := []any{internal.StrconvPackage.Ident("FormatUint"), "(uint64("}
	if isOptional {
		format = append(format, "*")
	}
	return append(append(format, srcValue...), []any{"), 10)"}...)
}

func (f *ClientGenerator) Int32KindFormat(srcValue []any, isOptional bool) []any {
	format := []any{internal.StrconvPackage.Ident("FormatInt"), "(int64("}
	if isOptional {
		format = append(format, "*")
	}
	return append(append(format, srcValue...), []any{"), 10)"}...)
}

func (f *ClientGenerator) BoolKindFormat(srcValue []any, isOptional bool) []any {
	format := []any{internal.StrconvPackage.Ident("FormatBool"), "("}
	if isOptional {
		format = append(format, "*")
	}
	return append(append(format, srcValue...), []any{")"}...)
}

func (f *ClientGenerator) HttpRequestFormat(srcValue []any) []any {
	return append(append([]any{}, srcValue...), []any{".Body"}...)
}

func (f *ClientGenerator) HttpBodyFormat(srcValue []any) []any {
	return append(append([]any{}, srcValue...), []any{".Data"}...)
}

func (f *ClientGenerator) WrapBytesFormat(srcValue []any) []any {
	return append(append([]any{}, srcValue...), []any{".Value"}...)
}

func (f *ClientGenerator) WrapStringFormat(srcValue []any) []any {
	return append(append([]any{}, srcValue...), []any{".Value"}...)
}

func (f *ClientGenerator) WrapBoolFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvPackage.Ident("FormatBool"), "("}, srcValue...), []any{".Value", ")"}...)
}

func (f *ClientGenerator) WrapUint32Format(srcValue []any) []any {
	return append(append([]any{internal.StrconvPackage.Ident("FormatUint"), "(uint64("}, srcValue...), []any{".Value", "), 10)"}...)
}

func (f *ClientGenerator) WrapInt32Format(srcValue []any) []any {
	return append(append([]any{internal.StrconvPackage.Ident("FormatInt"), "(int64("}, srcValue...), []any{".Value", "), 10)"}...)
}

func (f *ClientGenerator) WrapUint64Format(srcValue []any) []any {
	return append(append([]any{internal.StrconvPackage.Ident("FormatUint"), "("}, srcValue...), []any{".Value", ", 10)"}...)
}

func (f *ClientGenerator) WrapInt64Format(srcValue []any) []any {
	return append(append([]any{internal.StrconvPackage.Ident("FormatInt"), "("}, srcValue...), []any{".Value", ", 10)"}...)
}

func (f *ClientGenerator) WrapFloatFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvPackage.Ident("FormatFloat"), "(float64("}, srcValue...), []any{".Value", "), 'f', -1, 32)"}...)
}

func (f *ClientGenerator) WrapDoubleFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvPackage.Ident("FormatFloat"), "("}, srcValue...), []any{".Value", ", 'f', -1, 64)"}...)
}

func (f *ClientGenerator) TimestampFormat(srcValue []any) []any {
	//srcValue = append([]any{},srcValue...)
	return append(append([]any{"string(", internal.ErrorxPackage.Ident("Ignore"), "(", internal.JsonPackage.Ident("Marshal"), "("}, srcValue...), []any{")))"}...)
}

func (f *ClientGenerator) DurationFormat(srcValue []any) []any {
	return append(append([]any{"string(", internal.ErrorxPackage.Ident("Ignore"), "(", internal.JsonPackage.Ident("Marshal"), "("}, srcValue...), []any{")))"}...)
}

func (f *ClientGenerator) PrintDecodeResponseFunc(generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("func(ctx context.Context, r *", internal.HttpPackage.Ident("Response"), ") (interface{}, error) {")
	generatedFile.P("return nil, nil")
	generatedFile.P("},")

	return nil
}
