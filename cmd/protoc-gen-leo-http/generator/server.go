package generator

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
	"strings"
	"sync"
)

func errNotFoundField(endpoint *internal.Endpoint, names []string) error {
	return fmt.Errorf("%s, failed to find field %s", endpoint.FullName(), strings.Join(names, "."))
}

type ServerGenerator struct{}

func (f *ServerGenerator) GenerateNewServer(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("func New", service.HTTPServerName(), "(")
	generatedFile.P("endpoints interface {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.Name(), "() ", internal.EndpointPackage.Ident("Endpoint"))
	}
	generatedFile.P("},")
	generatedFile.P("mdw []", internal.EndpointPackage.Ident("Middleware"), ",")
	generatedFile.P("opts ...", internal.HttpTransportPackage.Ident("ServerOption"), ",")
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
		generatedFile.P(internal.EndpointxPackage.Ident("Chain"), "(endpoints.", endpoint.Name(), "(), mdw...), ")
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
			f.PrintGoogleApiHttpBodyBlock(generatedFile, []any{"req"}, []any{"r.Body"})
		case "google.rpc.HttpRequest":
			f.PrintGoogleRpcHttpRequest(generatedFile)
			generatedFile.P("return req, nil")
			generatedFile.P("},")
			return nil
		default:
			f.PrintDecodeBlock(generatedFile, internal.JsonPackage.Ident("NewDecoder"), []any{"req"}, []any{"r.Body"})
		}
	} else if bodyField != nil {
		if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
			generatedFile.P("req.", bodyField.GoName, " = &", bodyField.Message.GoIdent, "{}")
			tgtValue := []any{"req.", bodyField.GoName}
			srcValue := []any{"r.Body"}
			f.PrintGoogleApiHttpBodyBlock(generatedFile, tgtValue, srcValue)
		} else {
			tgtValue := []any{"req.", bodyField.GoName}
			srcValue := []any{"r.Body"}
			f.PrintDecodeBlock(generatedFile, internal.JsonPackage.Ident("NewDecoder"), tgtValue, srcValue)
		}
	}

	var pathOnce sync.Once
	if len(namedPathFields) > 0 {
		pathOnce.Do(func() {
			generatedFile.P("vars := ", internal.MuxPackage.Ident("Vars"), "(r)")
		})
		if err := f.PrintNamedPathField(generatedFile, namedPathFields, endpoint.HttpRule()); err != nil {
			return err
		}
	}

	for _, pathField := range pathFields {
		pathOnce.Do(func() {
			generatedFile.P("vars := ", internal.MuxPackage.Ident("Vars"), "(r)")
		})
		left := []any{"req.", pathField.GoName, " = "}
		right := []any{"vars[", strconv.Quote(string(pathField.Desc.Name())), "]"}
		if err := f.printAssign(generatedFile, pathField, left, right, false); err != nil {
			return err
		}
	}

	var queryOnce sync.Once
	for _, field := range queryFields {
		queryOnce.Do(func() {
			generatedFile.P("queries := r.URL.Query()")
		})
		fieldName := string(field.Desc.Name())
		if field.Message != nil && field.Message.Desc.FullName() == "google.protobuf.FieldMask" {
			if bodyField != nil {
				generatedFile.P("mask, err := ", internal.FieldmaskpbPackage.Ident("New"), "(req.", bodyField.GoName, ", queries[", strconv.Quote(fieldName), "]...)")
			} else if bodyMessage != nil {
				generatedFile.P("mask, err := ", internal.FieldmaskpbPackage.Ident("New"), "(req", ", queries[", strconv.Quote(fieldName), "]...)")
			}
			generatedFile.P("if err != nil {")
			generatedFile.P("return nil, err")
			generatedFile.P("}")
			generatedFile.P("req.UpdateMask = mask")
			continue
		}
		left := []any{"req.", field.GoName, " = "}
		right := []any{"queries.Get(", strconv.Quote(fieldName), ")"}
		if field.Desc.IsList() {
			right = []any{"queries[", strconv.Quote(fieldName), "]"}
		}
		if err := f.printAssign(generatedFile, field, left, right, field.Desc.IsList()); err != nil {
			return err
		}
	}

	generatedFile.P("return req, nil")
	generatedFile.P("},")
	return nil
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
				srcValue = append(srcValue, ", vars[", strconv.Quote(namedPathParameter), "]")
			}
			srcValue = append(srcValue, ")")
			if err := f.printAssign(generatedFile, namedPathField, tgtValue, srcValue, false); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *ServerGenerator) PrintDecodeBlock(generatedFile *protogen.GeneratedFile, decoder protogen.GoIdent, tgtValue []any, srcValue []any) {
	generatedFile.P(append(append(append(append([]any{"if err := ", decoder, "("}, srcValue...), []any{").Decode("}...), tgtValue...), []any{"); err != nil {"}...)...)
	generatedFile.P("return nil, err")
	generatedFile.P("}")
}

func (f *ServerGenerator) PrintGoogleApiHttpBodyBlock(generatedFile *protogen.GeneratedFile, tgtValue []any, srcValue []any) {
	generatedFile.P(append(append([]any{"body, err := ", internal.IOPackage.Ident("ReadAll"), "("}, srcValue...), []any{")"}...)...)
	generatedFile.P("if err != nil {")
	generatedFile.P("return nil, err")
	generatedFile.P("}")
	generatedFile.P(append(append([]any{}, tgtValue...), []any{".Data = body"}...)...)
	generatedFile.P(append(append([]any{}, tgtValue...), []any{".ContentType = r.Header.Get(", strconv.Quote(internal.ContentTypeKey), ")"}...)...)
}

func (f *ServerGenerator) PrintGoogleRpcHttpRequest(generatedFile *protogen.GeneratedFile) {
	generatedFile.P("req.Method = r.Method")
	generatedFile.P("req.Uri = r.RequestURI")
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

func (f *ServerGenerator) printAssign(generatedFile *protogen.GeneratedFile, field *protogen.Field, tgtValue []any, srcValue []any, isList bool) error {
	isOptional := field.Desc.HasOptionalKeyword()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		// bool
		if isList {
			srcValue = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseBoolSlice"), "("}, srcValue...)
		} else {
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseBool"), "("}, srcValue...)
		}
		srcValue = append(srcValue, "); err != nil {")
		generatedFile.P(srcValue...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if isOptional {
			generatedFile.P(append(tgtValue, internal.ProtoPackage.Ident("Bool"), "(v)")...)
		} else {
			generatedFile.P(append(tgtValue, "v")...)
		}
		generatedFile.P("}")
	case protoreflect.EnumKind:
		generatedFile.P("// enum")

	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		// int32
		if isList {
			srcValue = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseIntSlice[int32]"), "("}, srcValue...)
		} else {
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseInt"), "("}, srcValue...)
		}
		srcValue = append(srcValue, ", 10, 32); err != nil {")
		generatedFile.P(srcValue...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if isOptional {
			generatedFile.P(append(tgtValue, internal.ProtoPackage.Ident("Int32"), "(int32(v))")...)
		} else if isList {
			generatedFile.P(append(tgtValue, "v")...)
		} else {
			generatedFile.P(append(tgtValue, "int32(v)")...)
		}
		generatedFile.P("}")
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		// uint32
		if isList {
			srcValue = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseUintSlice[uint32]"), "("}, srcValue...)
		} else {
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseUint"), "("}, srcValue...)
		}
		srcValue = append(srcValue, ", 10, 32); err != nil {")
		generatedFile.P(srcValue...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if isOptional {
			generatedFile.P(append(tgtValue, internal.ProtoPackage.Ident("Uint32"), "(uint32(v))")...)
		} else if isList {
			generatedFile.P(append(tgtValue, "v")...)
		} else {
			generatedFile.P(append(tgtValue, "uint32(v)")...)
		}
		generatedFile.P("}")
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		// int64
		if isList {
			srcValue = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseIntSlice[int64]"), "("}, srcValue...)
		} else {
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseInt"), "("}, srcValue...)
		}
		srcValue = append(srcValue, ", 10, 64); err != nil {")
		generatedFile.P(srcValue...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if isOptional {
			generatedFile.P(append(tgtValue, internal.ProtoPackage.Ident("Int64"), "(v)")...)
		} else {
			generatedFile.P(append(tgtValue, "v")...)
		}
		generatedFile.P("}")
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		// uint64
		if isList {
			srcValue = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseUintSlice[uint64]"), "("}, srcValue...)
		} else {
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseUint"), "("}, srcValue...)
		}
		srcValue = append(srcValue, ", 10, 64); err != nil {")
		generatedFile.P(srcValue...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if isOptional {
			generatedFile.P(append(tgtValue, internal.ProtoPackage.Ident("Uint64"), "(v)")...)
		} else {
			generatedFile.P(append(tgtValue, "v")...)
		}
		generatedFile.P("}")
	case protoreflect.FloatKind:
		// float32
		if isList {
			srcValue = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseFloatSlice[float32]"), "("}, srcValue...)
		} else {
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseFloat"), "("}, srcValue...)
		}
		srcValue = append(srcValue, ", 32); err != nil {")
		generatedFile.P(srcValue...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if isOptional {
			generatedFile.P(append(tgtValue, internal.ProtoPackage.Ident("Float32"), "(float32(v))")...)
		} else if isList {
			generatedFile.P(append(tgtValue, "v")...)
		} else {
			generatedFile.P(append(tgtValue, "float32(v)")...)
		}
		generatedFile.P("}")
	case protoreflect.DoubleKind:
		// float64
		if isList {
			srcValue = append([]any{"if v, err := ", internal.ConvxPackage.Ident("ParseFloatSlice[float64]"), "("}, srcValue...)
		} else {
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseFloat"), "("}, srcValue...)
		}
		srcValue = append(srcValue, ", 32); err != nil {")
		generatedFile.P(srcValue...)
		generatedFile.P("return nil, err")
		generatedFile.P("} else {")
		if isOptional {
			generatedFile.P(append(tgtValue, internal.ProtoPackage.Ident("Float64"), "(v)")...)
		} else {
			generatedFile.P(append(tgtValue, "v")...)
		}
		generatedFile.P("}")
	case protoreflect.StringKind:
		// string
		if isOptional {
			a := []any{internal.ProtoPackage.Ident("String"), "("}
			srcValue = append(a, srcValue...)
			srcValue = append(srcValue, ")")
			generatedFile.P(append(tgtValue, srcValue...)...)
		} else {
			generatedFile.P(append(tgtValue, srcValue...)...)
		}
	case protoreflect.BytesKind:
		// []byte
		if isList {
			srcValue = append([]any{internal.ConvxPackage.Ident("ParseBytesSlice"), "("}, srcValue...)
			srcValue = append(srcValue, ")")
			generatedFile.P(append(tgtValue, srcValue...)...)
		} else {
			srcValue = append([]any{"[]byte("}, srcValue...)
			srcValue = append(srcValue, ")")
			generatedFile.P(append(tgtValue, srcValue...)...)
		}
	case protoreflect.MessageKind:
		switch field.Message.Desc.FullName() {
		case "google.protobuf.DoubleValue":
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseFloat"), "("}, srcValue...)
			srcValue = append(srcValue, ", 64); err != nil {")
			generatedFile.P(srcValue...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(tgtValue, internal.WrapperspbPackage.Ident("Double"), "(v)")...)
			generatedFile.P("}")
		case "google.protobuf.FloatValue":
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseFloat"), "("}, srcValue...)
			srcValue = append(srcValue, ", 32); err != nil {")
			generatedFile.P(srcValue...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(tgtValue, internal.WrapperspbPackage.Ident("Float"), "(float32(v))")...)
			generatedFile.P("}")
		case "google.protobuf.Int64Value":
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseInt"), "("}, srcValue...)
			srcValue = append(srcValue, ", 10, 64); err != nil {")
			generatedFile.P(srcValue...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(tgtValue, internal.WrapperspbPackage.Ident("Int64"), "(v)")...)
			generatedFile.P("}")
		case "google.protobuf.UInt64Value":
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseUint"), "("}, srcValue...)
			srcValue = append(srcValue, ", 10, 64); err != nil {")
			generatedFile.P(srcValue...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(tgtValue, internal.WrapperspbPackage.Ident("UInt64"), "(v)")...)
			generatedFile.P("}")
		case "google.protobuf.Int32Value":
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseInt"), "("}, srcValue...)
			srcValue = append(srcValue, ", 10, 32); err != nil {")
			generatedFile.P(srcValue...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(tgtValue, internal.WrapperspbPackage.Ident("Int32"), "(int32(v))")...)
			generatedFile.P("}")
		case "google.protobuf.UInt32Value":
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseUint"), "("}, srcValue...)
			srcValue = append(srcValue, ", 10, 32); err != nil {")
			generatedFile.P(srcValue...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(tgtValue, internal.WrapperspbPackage.Ident("UInt32"), "(uint32(v))")...)
			generatedFile.P("}")
		case "google.protobuf.BoolValue":
			srcValue = append([]any{"if v, err := ", internal.StrconvPackage.Ident("ParseBool"), "("}, srcValue...)
			srcValue = append(srcValue, "); err != nil {")
			generatedFile.P(srcValue...)
			generatedFile.P("return nil, err")
			generatedFile.P("} else {")
			generatedFile.P(append(tgtValue, internal.WrapperspbPackage.Ident("Bool"), "(v)")...)
			generatedFile.P("}")
		case "google.protobuf.StringValue":
			a := []any{internal.WrapperspbPackage.Ident("String"), "("}
			srcValue = append(a, srcValue...)
			srcValue = append(srcValue, ")")
			generatedFile.P(append(tgtValue, srcValue...)...)
		case "google.protobuf.BytesValue":
			a := []any{internal.WrapperspbPackage.Ident("Bytes"), "([]byte("}
			srcValue = append(a, srcValue...)
			srcValue = append(srcValue, "))")
			generatedFile.P(append(tgtValue, srcValue...)...)
		default:
			generatedFile.P("if err := ", internal.ProtoJsonPackage.Ident("Unmarshal"), "(body, req.", field.GoName, "); err != nil {")
			generatedFile.P("return nil, err")
			generatedFile.P("}")
		}
	case protoreflect.GroupKind:
		generatedFile.P("// group")

	default:
		return fmt.Errorf("unsupported field type: %+v", internal.FullMessageTypeName(field.Desc.Message()))
	}
	return nil
}

func (f *ServerGenerator) PrintEncodeResponseFunc(generatedFile *protogen.GeneratedFile, endpoint *internal.Endpoint, httpRule *internal.HttpRule) error {
	generatedFile.P("func(ctx ", internal.ContextPackage.Ident("Context"), ", w ", internal.HttpPackage.Ident("ResponseWriter"), ", obj any) error {")
	generatedFile.P("resp := obj.(*", endpoint.Output().GoIdent, ")")
	generatedFile.P("_ = resp")
	bodyParameter := httpRule.ResponseBody()
	switch bodyParameter {
	case "":
		if err := f.PrintResponse(generatedFile, endpoint.Output(), "resp"); err != nil {
			return err
		}
	default:
		field := internal.FindField(bodyParameter, endpoint.Output())
		if field == nil {
			return errNotFoundField(endpoint, []string{bodyParameter})
		}
		if err := f.PrintResponse(generatedFile, field.Message, "resp."+field.GoName); err != nil {
			return err
		}
	}
	generatedFile.P("return nil")
	return nil
}

func (f *ServerGenerator) PrintResponse(generatedFile *protogen.GeneratedFile, message *protogen.Message, prefix string) error {
	switch message.Desc.FullName() {
	case "google.api.HttpBody":
		generatedFile.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
		generatedFile.P("w.Header().Set(", strconv.Quote("Content-Type"), ", ", prefix, ".GetContentType())")
		generatedFile.P()
		generatedFile.P("if ", "_, err := w.Write(", prefix, ".GetData())", "; err != nil {")
		generatedFile.P("return err")
		generatedFile.P("}")
	case "google.rpc.HttpResponse":
		generatedFile.P("w.WriteHeader(int(", prefix, ".GetStatus()))")
		generatedFile.P("for _, header := range ", prefix, ".GetHeaders() {")
		generatedFile.P("w.Header().Add(header.Key, header.Value)")
		generatedFile.P("}")
		generatedFile.P("if ", "_, err := w.Write(", prefix, ".GetBody())", "; err != nil {")
		generatedFile.P("return err")
		generatedFile.P("}")
	default:
		generatedFile.P("w.WriteHeader(", internal.HttpPackage.Ident("StatusOK"), ")")
		generatedFile.P("data, err := ", internal.ProtoJsonPackage.Ident("Marshal"), "(", prefix, ")")
		generatedFile.P("if err != nil {")
		generatedFile.P("return err")
		generatedFile.P("}")
		generatedFile.P("if _, err := w.Write(data); err != nil {")
		generatedFile.P("return err")
		generatedFile.P("}")
	}
	return nil
}
