package http

import (
	"fmt"
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
)

type ClientGenerator struct{}

func (f *ClientGenerator) GenerateTransports(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedHttpClientTransportsName(), " struct {")
	g.P("scheme        string")
	g.P("router        *", internal.MuxPackage.Ident("Router"))
	g.P("clientOptions []", internal.HttpTransportPackage.Ident("ClientOption"))
	g.P("middlewares   []", internal.EndpointPackage.Ident("Middleware"))
	g.P("}")
	g.P()

	for _, endpoint := range service.Endpoints {
		g.P("func (t *", service.UnexportedHttpClientTransportsName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", instance string) (", internal.EndpointPackage.Ident("Endpoint"), ", ", internal.IOPackage.Ident("Closer"), ", error) {")
		g.P("opts := []", internal.HttpTransportPackage.Ident("ClientOption"), "{")
		g.P(internal.HttpTransportPackage.Ident("ClientBefore"), "(", internal.HttpxTransportxPackage.Ident("OutgoingMetadataInjector"), "),")
		g.P(internal.HttpTransportPackage.Ident("ClientBefore"), "(", internal.HttpxTransportxPackage.Ident("OutgoingTimeLimiter"), "),")
		g.P(internal.HttpTransportPackage.Ident("ClientBefore"), "(", internal.HttpxTransportxPackage.Ident("OutgoingStain"), "),")
		g.P("}")
		g.P("opts = append(opts, t.clientOptions...)")
		g.P("client := ", internal.HttpTransportPackage.Ident("NewExplicitClient"), "(")
		g.P(endpoint.HttpClientRequestEncoderName(), "(t.router)(t.scheme, instance),")
		g.P(endpoint.HttpClientResponseDecoderName(), ",")
		g.P("opts...,")
		g.P(")")
		g.P("return ", internal.EndpointxPackage.Ident("Chain"), "(client.Endpoint(), t.middlewares...), nil, nil")
		g.P("}")
		g.P()
	}

	g.P("func new", service.HttpClientTransportsName(), "(scheme string, clientOptions []", internal.HttpTransportPackage.Ident("ClientOption"), ", middlewares []", internal.EndpointPackage.Ident("Middleware"), ") ", service.ClientTransportsNameV2(), " {")
	g.P("return &", service.UnexportedHttpClientTransportsName(), "{")
	g.P("scheme:        scheme,")
	g.P("router:        append", service.HttpRoutesName(), "(", internal.MuxPackage.Ident("NewRouter"), "()),")
	g.P("clientOptions: clientOptions,")
	g.P("middlewares:   middlewares,")
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func (f *ClientGenerator) GenerateClient(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P("type ", service.UnexportedHttpClientName(), " struct {")
	g.P("balancers ", service.BalancersName())
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		g.P("func (c *", service.UnexportedHttpClientName(), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
		g.P("ctx = ", internal.EndpointxPackage.Ident("InjectName"), "(ctx, ", strconv.Quote(endpoint.FullName()), ")")
		g.P("ctx = ", internal.TransportxPackage.Ident("InjectName"), "(ctx, ", internal.HttpxTransportxPackage.Ident("HttpClient"), ")")
		g.P("balancer, err := c.balancers.", endpoint.Name(), "(ctx)")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("endpoint, err := balancer.Endpoint()")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.P("rep, err := endpoint(ctx, request)")
		g.P("if err != nil {")
		g.P("return nil, ", internal.StatusxPackage.Ident("From"), "(err)")
		g.P("}")
		g.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
		g.P("}")
		g.P()
	}
	g.P("func New", service.HttpClientName(), "(target string, opts ...", internal.HttpxTransportxPackage.Ident("ClientOption"), ") ", service.ServiceName(), " {")
	g.P("options := ", internal.HttpxTransportxPackage.Ident("NewClientOptions"), "(opts...)")
	g.P("transports := new", service.HttpClientTransportsName(), "(options.Scheme(), options.ClientTransportOptions(), options.Middlewares())")
	g.P("factories := new", service.FactoriesName(), "(transports)")
	g.P("endpointers := new", service.EndpointersName(), "(target, options.InstancerFactory(), factories, options.Logger(), options.EndpointerOptions()...)")
	g.P("balancers := new", service.BalancersName(), "(options.BalancerFactory(), endpointers)")
	g.P("return &", service.UnexportedHttpClientName(), "{balancers: balancers}")
	g.P("}")
	g.P()
	return nil
}

func (f *ClientGenerator) GenerateCoder(service *internal.Service, g *protogen.GeneratedFile) error {
	for _, endpoint := range service.Endpoints {
		if err := f.PrintEncodeRequestFunc(g, endpoint); err != nil {
			return err
		}
		if err := f.PrintDecodeResponseFunc(g, endpoint); err != nil {
			return err
		}
	}
	return nil
}

func (f *ClientGenerator) PrintEncodeRequestFunc(g *protogen.GeneratedFile, endpoint *internal.Endpoint) error {
	httpRule := endpoint.HttpRule()
	g.P("func ", endpoint.HttpClientRequestEncoderName(), "(router *", internal.MuxPackage.Ident("Router"), ") func(scheme string, instance string) ", internal.HttpTransportPackage.Ident("CreateRequestFunc"), " {")
	g.P("return func(scheme string, instance string) ", internal.HttpTransportPackage.Ident("CreateRequestFunc"), " {")
	g.P("return func(ctx context.Context, obj any) (*", internal.HttpPackage.Ident("Request"), ", error) {")
	g.P("if obj == nil {")
	g.P("return nil, ", internal.StatusxPackage.Ident("ErrInvalidArgument"), ".With(", internal.StatusxPackage.Ident("Message"), "(", strconv.Quote("request is nil"), "))")
	g.P("}")
	g.P("req, ok := obj.(*", endpoint.InputGoIdent(), ")")
	g.P("if !ok {")
	g.P("return nil, ", internal.StatusxPackage.Ident("ErrInvalidArgument"), ".With(", internal.StatusxPackage.Ident("Message"), "(", strconv.Quote("invalid request type, %T"), ", obj))")
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

	g.P("path, err := router.Get(", strconv.Quote(endpoint.FullName()), ").URLPath(pairs...)")
	g.P("if err != nil {")
	g.P("return nil, ", internal.StatusxPackage.Ident("ErrInvalidArgument"), ".With(", internal.StatusxPackage.Ident("Wrap"), "(err))")
	g.P("}")

	g.P("queries := ", internal.UrlPackage.Ident("Values"), "{}")
	if len(queryFields) > 0 {
		f.PrintQueryField(g, queryFields)
	}

	g.P("target := &", internal.UrlPackage.Ident("URL"), "{")
	g.P("Scheme:   scheme,")
	g.P("Host:     instance,")
	g.P("Path:     path.Path,")
	g.P("RawQuery: queries.Encode(),")
	g.P("}")

	g.P("r, err := ", internal.HttpPackage.Ident("NewRequestWithContext"), "(ctx, ", strconv.Quote(httpRule.Method()), ", target.String(), body)")
	g.P("if err != nil {")
	g.P("return nil, ", internal.StatusxPackage.Ident("ErrInvalidArgument"), ".With(", internal.StatusxPackage.Ident("Wrap"), "(err))")
	g.P("}")
	if bodyMessage != nil || bodyField != nil {
		g.P("r.Header.Set(", strconv.Quote(internal.ContentTypeKey), ", contentType)")
	}
	g.P("return r, nil")
	g.P("}")
	g.P("}")
	g.P("}")
	g.P()
	return nil
}

func (f *ClientGenerator) PrintDecodeResponseFunc(g *protogen.GeneratedFile, endpoint *internal.Endpoint) error {
	httpRule := endpoint.HttpRule()
	g.P("func ", endpoint.HttpClientResponseDecoderName(), "(ctx context.Context, r *", internal.HttpPackage.Ident("Response"), ") (any, error) {")
	g.P("if ", internal.HttpxTransportxPackage.Ident("IsErrorResponse"), "(r) {")
	g.P("return nil, ", internal.HttpxTransportxPackage.Ident("ErrorDecoder"), "(ctx, r)")
	g.P("}")
	g.P("resp := &", endpoint.Output().GoIdent, "{}")
	bodyParameter := httpRule.ResponseBody()
	switch bodyParameter {
	case "", "*":
		srcValue := []any{"resp"}
		message := endpoint.Output()
		switch message.Desc.FullName() {
		case "google.api.HttpBody":
			f.PrintGoogleApiHttpBodyDecodeBlock(g, srcValue)
		default:
			f.PrintJsonDecodeBlock(g, srcValue)
		}
	default:
		bodyField := internal.FindField(bodyParameter, endpoint.Output())
		if bodyField == nil {
			return fmt.Errorf("%s, failed to find body response field %s", endpoint.FullName(), bodyParameter)
		}
		srcValue := []any{"resp.", bodyField.GoName}
		if bodyField.Desc.Kind() == protoreflect.MessageKind && bodyField.Message.Desc.FullName() == "google.api.HttpBody" {
			g.P(append(append([]any{}, srcValue...), " = &", bodyField.Message.GoIdent, "{}")...)
			f.PrintGoogleApiHttpBodyDecodeBlock(g, srcValue)
		} else {
			_, pointer := internal.FieldGoType(g, bodyField)
			if !pointer {
				srcValue = append([]any{"&"}, srcValue...)
			}
			f.PrintJsonDecodeBlock(g, srcValue)
		}
	}
	g.P("return resp, nil")
	g.P("}")
	g.P()
	return nil
}

func (f *ClientGenerator) PrintReaderBlock(g *protogen.GeneratedFile, readerPkg protogen.GoImportPath, tgtValue []any, srcValue []any) {
	g.P(append(append(append(append([]any{}, tgtValue...), []any{" = ", readerPkg.Ident("NewReader"), "("}...), srcValue...), ")")...)
}

func (f *ClientGenerator) PrintEncodeBlock(g *protogen.GeneratedFile, encoder protogen.GoIdent, tgtValue []any, srcValue []any) {
	g.P(append(append(append(append([]any{"if err := ", encoder, "("}, tgtValue...), []any{").Encode("}...), srcValue...), []any{"); err != nil {"}...)...)
	g.P("return nil, ", internal.StatusxPackage.Ident("ErrInvalidArgument"), ".With(", internal.StatusxPackage.Ident("Wrap"), "(err))")
	g.P("}")
}

func (f *ClientGenerator) PrintNamedPathField(g *protogen.GeneratedFile, namedPathFields []*protogen.Field, httpRule *internal.HttpRule) {
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
	g.P("return nil, ", internal.StatusxPackage.Ident("ErrInvalidArgument"), ".With(", internal.StatusxPackage.Ident("Message"), "(", strconv.Quote("invalid named path parameter, %s"), ", namedPathParameter))")
	g.P("}")

	pairs := []any{"pairs = append(pairs"}
	for i, parameter := range namedPathParameters {
		pairs = append(pairs, ",", strconv.Quote(parameter), ",", fmt.Sprintf("namedPathValues[%d]", i*2+1))
	}
	pairs = append(pairs, ")")
	g.P(pairs...)
}

func (f *ClientGenerator) PrintPathField(g *protogen.GeneratedFile, pathFields []*protogen.Field) {
	pairs := []any{"pairs = append(pairs"}
	for _, field := range pathFields {
		pairs = append(append(pairs, ",", strconv.Quote(string(field.Desc.Name())), ","), f.PathFieldFormat(field)...)
	}
	pairs = append(pairs, ")")
	g.P(pairs...)
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

func (f *ClientGenerator) PrintQueryField(g *protogen.GeneratedFile, queryFields []*protogen.Field) {
	for _, field := range queryFields {
		srcValue := []any{"req.Get", field.GoName, "()"}
		fieldName := string(field.Desc.Name())
		switch field.Desc.Kind() {
		case protoreflect.BoolKind: // bool
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.BoolListFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.BoolValueFormat(srcValue))
			}
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind: // int32
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.IntListFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.IntValueFormat(srcValue))
			}
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind: // uint32
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.UintListFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.UintValueFormat(srcValue))
			}
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind: // int64
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.IntListFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.IntValueFormat(srcValue))
			}
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind: // uint64
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.UintListFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.UintValueFormat(srcValue))
			}
		case protoreflect.FloatKind: // float32
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.FloatListFormat(srcValue, "32"), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.FloatValueFormat(srcValue, "32"))
			}
		case protoreflect.DoubleKind: // float64
			if field.Desc.IsList() {
				f.PrintQuery(g, fieldName, append(f.FloatListFormat(srcValue, "64"), []any{"..."}...))
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
				f.PrintQuery(g, fieldName, append(f.IntListFormat(srcValue), []any{"..."}...))
			} else {
				f.PrintQuery(g, fieldName, f.IntValueFormat(srcValue))
			}
		case protoreflect.MessageKind:
			switch field.Message.Desc.FullName() {
			case "google.protobuf.DoubleValue":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapFloatListFormat(srcValue, "64"), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapFloatValueFormat(srcValue, "64"))
				}
			case "google.protobuf.FloatValue":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapFloatListFormat(srcValue, "32"), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapFloatValueFormat(srcValue, "32"))
				}
			case "google.protobuf.Int64Value":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapIntListFormat(srcValue, "64"), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapIntValueFormat(srcValue))
				}
			case "google.protobuf.UInt64Value":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapUintListFormat(srcValue, "64"), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapUintValueFormat(srcValue))
				}
			case "google.protobuf.Int32Value":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapIntListFormat(srcValue, "32"), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapIntValueFormat(srcValue))
				}
			case "google.protobuf.UInt32Value":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapUintListFormat(srcValue, "32"), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapUintValueFormat(srcValue))
				}
			case "google.protobuf.BoolValue":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapBoolListFormat(srcValue), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapBoolValueFormat(srcValue))
				}
			case "google.protobuf.StringValue":
				if field.Desc.IsList() {
					f.PrintQuery(g, fieldName, append(f.UnwrapStringListFormat(srcValue), []any{"..."}...))
				} else {
					f.PrintQuery(g, fieldName, f.UnwrapStringValueFormat(srcValue))
				}
			}
		}
	}
}

func (f *ClientGenerator) PrintQuery(g *protogen.GeneratedFile, fieldName string, srcValue []any) {
	g.P(append(append([]any{"queries[", strconv.Quote(fieldName), "] = append(queries[", strconv.Quote(fieldName), "], "}, srcValue...), []any{")"}...)...)
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

func (f *ClientGenerator) PrintJsonDecodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{"if err := ", internal.JsonxPackage.Ident("NewDecoder"), "(r.Body).Decode("}, srcValue...), "); err != nil {")...)
	g.P("return nil, err")
	g.P("}")
}

func (f *ClientGenerator) PrintGoogleApiHttpBodyDecodeBlock(g *protogen.GeneratedFile, srcValue []any) {
	g.P(append(append([]any{}, srcValue...), ".ContentType = r.Header.Get(", strconv.Quote(internal.ContentTypeKey), ")")...)
	g.P("body, err := ", internal.IOPackage.Ident("ReadAll"), "(r.Body)")
	g.P("if err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P(append(append([]any{}, srcValue...), ".Data = body")...)

}
