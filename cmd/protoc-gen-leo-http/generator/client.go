package generator

import (
	"errors"
	"fmt"
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
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
	generatedFile.P("opts []", internal.HttpTransportPackage.Ident("ClientOption"), ",")
	generatedFile.P("mdw ...", internal.EndpointPackage.Ident("Middleware"), ",")
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
		httpRule := endpoint.HttpRule()
		generatedFile.P(endpoint.UnexportedName(), ":    ", internal.EndpointxPackage.Ident("Chain"), "(")
		generatedFile.P(internal.HttpTransportPackage.Ident("NewExplicitClient"), "(")
		if err := f.PrintEncodeRequestFunc(generatedFile, endpoint); err != nil {
			return err
		}
		if err := f.PrintDecodeResponseFunc(generatedFile, endpoint, httpRule); err != nil {
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
			encoder := internal.JsonxPackage.Ident("NewEncoder")
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
			encoder := internal.JsonxPackage.Ident("NewEncoder")
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
		f.PrintPathField(generatedFile, pathFields)
	}

	generatedFile.P("path, err := router.Get(", strconv.Quote(endpoint.FullName()), ").URLPath(pairs...)")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")

	generatedFile.P("queries := ", internal.UrlPackage.Ident("Values"), "{}")
	if len(queryFields) > 0 {
		f.PrintQueryField(generatedFile, queryFields)
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
		generatedFile.P("namedPathParameter := req.", fullFieldGetterName)
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

func (f *ClientGenerator) PrintPathField(generatedFile *protogen.GeneratedFile, pathFields []*protogen.Field) {
	pairs := []any{"pairs = append(pairs"}
	for _, field := range pathFields {
		pairs = append(append(pairs, ",", strconv.Quote(string(field.Desc.Name())), ","), f.PathFieldFormat(field)...)
	}
	pairs = append(pairs, ")")
	generatedFile.P(pairs...)
}

func (f *ClientGenerator) PathFieldFormat(field *protogen.Field) []any {
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

func (f *ClientGenerator) PrintQueryField(generatedFile *protogen.GeneratedFile, queryFields []*protogen.Field) {
	for _, field := range queryFields {
		srcValue := []any{"req.Get", field.GoName, "()"}
		fieldName := string(field.Desc.Name())
		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
			if field.Desc.IsList() {
				f.PrintQuery(generatedFile, fieldName, append(f.BoolListFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(generatedFile, fieldName, f.BoolValueFormat(srcValue))
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
			if field.Desc.IsList() {
				f.PrintQuery(generatedFile, fieldName, append(f.IntListFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(generatedFile, fieldName, f.IntValueFormat(srcValue))
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
			if field.Desc.IsList() {
				f.PrintQuery(generatedFile, fieldName, append(f.UintListFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(generatedFile, fieldName, f.UintValueFormat(srcValue))
			}
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
			if field.Desc.IsList() {
				f.PrintQuery(generatedFile, fieldName, append(f.IntListFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(generatedFile, fieldName, f.IntValueFormat(srcValue))
			}
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
			if field.Desc.IsList() {
				f.PrintQuery(generatedFile, fieldName, append(f.UintListFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(generatedFile, fieldName, f.UintValueFormat(srcValue))
			}
		case protoreflect.FloatKind: // float32
			if field.Desc.IsList() {
				f.PrintQuery(generatedFile, fieldName, append(f.FloatListFormat(srcValue, "32"), []any{"..."}...))
			} else {
				f.PrintQuery(generatedFile, fieldName, f.FloatValueFormat(srcValue, "32"))
			}
		case protoreflect.DoubleKind: // float64
			if field.Desc.IsList() {
				f.PrintQuery(generatedFile, fieldName, append(f.FloatListFormat(srcValue, "64"), []any{"..."}...))
			} else {
				f.PrintQuery(generatedFile, fieldName, f.FloatValueFormat(srcValue, "64"))
			}
		case protoreflect.StringKind: // string
			if field.Desc.IsList() {
				f.PrintQuery(generatedFile, fieldName, append(f.StringKindFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(generatedFile, fieldName, f.StringKindFormat(srcValue))
			}
		case protoreflect.EnumKind: // enum int32
			if field.Desc.IsList() {
				f.PrintQuery(generatedFile, fieldName, append(f.IntListFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(generatedFile, fieldName, f.IntValueFormat(srcValue))
			}
		case protoreflect.MessageKind:
			switch field.Message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
				if field.Desc.IsList() {
					f.PrintQuery(generatedFile, fieldName, append(f.UnwrapFloatListFormat(srcValue, "64"), []any{"..."}...))
				} else {
					f.PrintQuery(generatedFile, fieldName, f.UnwrapFloatValueFormat(srcValue, "64"))
				}
			case "google.protobuf.FloatValue":
				if field.Desc.IsList() {
					f.PrintQuery(generatedFile, fieldName, append(f.UnwrapFloatListFormat(srcValue, "32"), []any{"..."}...))
				} else {
					f.PrintQuery(generatedFile, fieldName, f.UnwrapFloatValueFormat(srcValue, "32"))
				}
			case "google.protobuf.Int64Value":
				if field.Desc.IsList() {
					f.PrintQuery(generatedFile, fieldName, append(f.UnwrapIntListFormat(srcValue, "64"), []any{"..."}...))
				} else {
					f.PrintQuery(generatedFile, fieldName, f.UnwrapIntValueFormat(srcValue))
				}
			case "google.protobuf.UInt64Value":
				if field.Desc.IsList() {
					f.PrintQuery(generatedFile, fieldName, append(f.UnwrapUintListFormat(srcValue, "64"), []any{"..."}...))
				} else {
					f.PrintQuery(generatedFile, fieldName, f.UnwrapUintValueFormat(srcValue))
				}
			case "google.protobuf.Int32Value":
				if field.Desc.IsList() {
					f.PrintQuery(generatedFile, fieldName, append(f.UnwrapIntListFormat(srcValue, "32"), []any{"..."}...))
				} else {
					f.PrintQuery(generatedFile, fieldName, f.UnwrapIntValueFormat(srcValue))
				}
			case "google.protobuf.UInt32Value":
				if field.Desc.IsList() {
					f.PrintQuery(generatedFile, fieldName, append(f.UnwrapUintListFormat(srcValue, "32"), []any{"..."}...))
				} else {
					f.PrintQuery(generatedFile, fieldName, f.UnwrapUintValueFormat(srcValue))
				}
			case "google.protobuf.BoolValue":
				if field.Desc.IsList() {
					f.PrintQuery(generatedFile, fieldName, append(f.UnwrapBoolListFormat(srcValue), []any{"..."}...))
				} else {
					f.PrintQuery(generatedFile, fieldName, f.UnwrapBoolValueFormat(srcValue))
				}
			case "google.protobuf.StringValue":
				if field.Desc.IsList() {
					f.PrintQuery(generatedFile, fieldName, append(f.UnwrapStringListFormat(srcValue), []any{"..."}...))
				} else {
					f.PrintQuery(generatedFile, fieldName, f.UnwrapStringValueFormat(srcValue))
				}
			}
		}
	}
}

func (f *ClientGenerator) PrintQuery(generatedFile *protogen.GeneratedFile, fieldName string, srcValue []any) {
	generatedFile.P(append(append([]any{"queries[", strconv.Quote(fieldName), "] = append(queries[", strconv.Quote(fieldName), "], "}, srcValue...), []any{")"}...)...)
}

func (f *ClientGenerator) StringKindFormat(srcValue []any) []any {
	return append(append([]any{}, srcValue...), []any{}...)
}

func (f *ClientGenerator) FloatValueFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatFloat"), "("}, srcValue...), []any{", 'f', -1, ", bitSize, ")"}...)
}

func (f *ClientGenerator) FloatListFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatFloatSlice"), "("}, srcValue...), []any{", 'f', -1, ", bitSize, ")"}...)
}

func (f *ClientGenerator) UintValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatUint"), "("}, srcValue...), []any{", 10)"}...)
}

func (f *ClientGenerator) UintListFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatUintSlice"), "("}, srcValue...), []any{", 10)"}...)
}

func (f *ClientGenerator) IntValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatInt"), "("}, srcValue...), []any{", 10)"}...)
}

func (f *ClientGenerator) IntListFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatIntSlice"), "("}, srcValue...), []any{", 10)"}...)
}

func (f *ClientGenerator) BoolValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatBool"), "("}, srcValue...), []any{")"}...)
}

func (f *ClientGenerator) BoolListFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatBoolSlice"), "("}, srcValue...), []any{")"}...)
}

func (f *ClientGenerator) UnwrapStringValueFormat(srcValue []any) []any {
	return append(append([]any{}, srcValue...), []any{".GetValue()"}...)
}

func (f *ClientGenerator) UnwrapStringListFormat(srcValue []any) []any {
	return append(append([]any{internal.ProtoxPackage.Ident("UnwrapStringSlice"), "("}, srcValue...), []any{")"}...)
}

func (f *ClientGenerator) UnwrapBoolValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatBool"), "("}, srcValue...), []any{".GetValue()", ")"}...)
}

func (f *ClientGenerator) UnwrapBoolListFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatBoolSlice"), "(", internal.ProtoxPackage.Ident("UnwrapBoolSlice"), "("}, srcValue...), []any{"))"}...)
}

func (f *ClientGenerator) UnwrapIntValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatInt"), "("}, srcValue...), []any{".GetValue()", ", 10)"}...)
}

func (f *ClientGenerator) UnwrapIntListFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatIntSlice"), "(", internal.ProtoxPackage.Ident("UnwrapInt" + bitSize + "Slice"), "("}, srcValue...), []any{"), 10)"}...)
}

func (f *ClientGenerator) UnwrapUintValueFormat(srcValue []any) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatUint"), "("}, srcValue...), []any{".GetValue()", ", 10)"}...)
}

func (f *ClientGenerator) UnwrapUintListFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatUintSlice"), "(", internal.ProtoxPackage.Ident("UnwrapUint" + bitSize + "Slice"), "("}, srcValue...), []any{"), 10)"}...)
}

func (f *ClientGenerator) UnwrapFloatValueFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatFloat"), "("}, srcValue...), []any{".GetValue()", ", 'f', -1, ", bitSize, ")"}...)
}

func (f *ClientGenerator) UnwrapFloatListFormat(srcValue []any, bitSize string) []any {
	return append(append([]any{internal.StrconvxPackage.Ident("FormatFloatSlice"), "(", internal.ProtoxPackage.Ident("UnwrapFloat" + bitSize + "Slice"), "("}, srcValue...), []any{"), 'f', -1, ", bitSize, ")"}...)
}

func (f *ClientGenerator) PrintDecodeResponseFunc(generatedFile *protogen.GeneratedFile, endpoint *internal.Endpoint, httpRule *internal.HttpRule) error {
	generatedFile.P("func(ctx context.Context, r *", internal.HttpPackage.Ident("Response"), ") (interface{}, error) {")
	generatedFile.P("resp := &", endpoint.Output().GoIdent, "{}")
	bodyParameter := httpRule.ResponseBody()
	switch bodyParameter {
	case "", "*":
		srcValue := []any{"resp"}
		message := endpoint.Output()
		switch message.Desc.FullName() {
		case "google.api.HttpBody":
			f.PrintGoogleApiHttpBodyDecodeBlock(generatedFile, srcValue)
		case "google.rpc.HttpResponse":
			f.PrintGoogleRpcHttpResponseDecodeBlock(generatedFile, srcValue)
		default:
			f.PrintJsonDecodeBlock(generatedFile, srcValue)
		}
	default:
		bodyField := internal.FindField(bodyParameter, endpoint.Output())
		if bodyField == nil {
			return fmt.Errorf("%s, failed to find body response field %s", endpoint.FullName(), bodyParameter)
		}
		if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.rpc.HttpResponse" {
			return errors.New("google.rpc.HttpResponse can only be used as output to a method")
		}
		srcValue := []any{"resp.", bodyField.GoName}
		if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
			generatedFile.P(append(append([]any{}, srcValue...), " = &", bodyField.Message.GoIdent, "{}")...)
			f.PrintGoogleApiHttpBodyDecodeBlock(generatedFile, srcValue)
		} else {
			_, pointer := internal.FieldGoType(generatedFile, bodyField)
			if !pointer {
				srcValue = append([]any{"&"}, srcValue...)
			}
			f.PrintJsonDecodeBlock(generatedFile, srcValue)
		}
	}
	generatedFile.P("return resp, nil")
	generatedFile.P("},")
	return nil
}

func (f *ClientGenerator) PrintJsonDecodeBlock(generatedFile *protogen.GeneratedFile, srcValue []any) {
	generatedFile.P(append(append([]any{"if err := ", internal.JsonxPackage.Ident("NewDecoder"), "(r.Body).Decode("}, srcValue...), "); err != nil {")...)
	generatedFile.P("return nil, err")
	generatedFile.P("}")
}

func (f *ClientGenerator) PrintGoogleRpcHttpResponseDecodeBlock(generatedFile *protogen.GeneratedFile, srcValue []any) {
	generatedFile.P(append(append([]any{}, srcValue...), ".Status = int32(r.StatusCode)")...)
	generatedFile.P(append(append([]any{}, srcValue...), ".Reason = r.Status")...)
	generatedFile.P("for key, values := range r.Header {")
	generatedFile.P("for _, value := range values {")
	generatedFile.P(append(append(append(append([]any{}, srcValue...), ".Headers = append("), srcValue...), ".Headers, &", internal.RpcHttpPackage.Ident("HttpHeader"), "{Key: key, Value: value})")...)
	generatedFile.P("}")
	generatedFile.P("}")
	generatedFile.P("body, err := ", internal.IOPackage.Ident("ReadAll"), "(r.Body)")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P(append(append([]any{}, srcValue...), ".Body = body")...)
}

func (f *ClientGenerator) PrintGoogleApiHttpBodyDecodeBlock(generatedFile *protogen.GeneratedFile, srcValue []any) {
	generatedFile.P(append(append([]any{}, srcValue...), ".ContentType = ", strconv.Quote(internal.JsonContentType))...)
	generatedFile.P("body, err := ", internal.IOPackage.Ident("ReadAll"), "(r.Body)")
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P(append(append([]any{}, srcValue...), ".Data = body")...)

}
