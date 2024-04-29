package generator

import (
	"fmt"
	"github.com/go-leo/gox/slicex"
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
	"strings"
)

type Generator struct {
	Plugin   *protogen.Plugin
	File     *protogen.File
	Services []*internal.Service
}

func NewGenerator(plugin *protogen.Plugin, file *protogen.File) (*Generator, error) {
	services, err := internal.NewServices(file)
	if err != nil {
		return nil, err
	}
	return &Generator{Plugin: plugin, File: file, Services: services}, nil
}

func (f *Generator) Generate() error {
	return f.GenerateFile()
}

func (f *Generator) GenerateFile() error {
	file := f.File
	filename := file.GeneratedFilenamePrefix + "_http.leo.pb.go"
	g := f.Plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-grpc. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	for _, service := range f.Services {
		if err := f.GenerateNewServer(service, g); err != nil {
			return err
		}
	}
	//
	//for _, service := range f.Services {
	//	if err := f.GenerateClient(service, g); err != nil {
	//		return err
	//	}
	//}
	//
	//for _, service := range f.Services {
	//	if err := f.GenerateNewClient(service, g); err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (f *Generator) GenerateNewServer(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
	generatedFile.P("func New", service.HTTPServerName(), "(")
	generatedFile.P("endpoints interface {")
	for _, endpoint := range service.Endpoints {
		generatedFile.P(endpoint.Name(), "() ", internal.EndpointPackage.Ident("Endpoint"))
	}
	generatedFile.P("},")
	generatedFile.P("mdw []", internal.EndpointPackage.Ident("Middleware"), ",")
	generatedFile.P("opts ...", internal.HttpTransportPackage.Ident("ServerOption"), ",")
	generatedFile.P(") ", internal.HttpPackage.Ident("Handler"), " {")
	generatedFile.P("r := ", internal.MuxPackage.Ident("NewRouter"), "()")
	for _, endpoint := range service.Endpoints {
		for _, httpRule := range endpoint.HttpRules() {
			method := httpRule.Method()
			path := httpRule.Path()
			// 调整路径，来适应 github.com/gorilla/mux 路由规则
			path, names, template, namedPathParameters := httpRule.RegularizePath(path)
			_ = template
			generatedFile.P("r.Methods(", strconv.Quote(method), ").")
			generatedFile.P("Path(", strconv.Quote(path), ").")
			generatedFile.P("Handler(", internal.HttpTransportPackage.Ident("NewServer"), "(")
			generatedFile.P(internal.EndpointxPackage.Ident("Chain"), "(endpoints.", endpoint.Name(), "(), mdw...), ")
			generatedFile.P("func(ctx ", internal.ContextPackage.Ident("Context"), ", r *", internal.HttpPackage.Ident("Request"), ") (any, error) {")
			generatedFile.P("var req *", endpoint.InputGoIdent())

			// body arguments
			{
				switch httpRule.Body() {
				case "":
				case "*":
				default:
				}
			}
			// path arguments
			{
				pathParameters := httpRule.PathParameters(path)
				if len(pathParameters) > 0 {
					generatedFile.P("vars := ", internal.MuxPackage.Ident("Vars"), "(r)")
					// 命名路径参数设值
					if len(namedPathParameters) > 0 {
						errNotFoundField := fmt.Errorf("%s, failed to find field %s", endpoint.FullName(), strings.Join(names, "."))
						errInvalidType := fmt.Errorf("%s, %s field type invalid", endpoint.FullName(), strings.Join(names, "."))

						inMessage := endpoint.Input()
						var fields []*protogen.Field
						for i := 0; i < len(names)-1; i++ {
							name := names[i]
							field := internal.FindField(name, inMessage)
							if field == nil {
								return errNotFoundField
							}
							if field.Desc.Kind() != protoreflect.MessageKind {
								return errInvalidType
							}
							fields = append(fields, field)
							inMessage = field.Message
							var fieldNames []string
							for _, p := range fields {
								fieldNames = append(fieldNames, p.GoName)
							}
							fullFieldName := strings.Join(fieldNames, ".")
							generatedFile.P("if req.", fullFieldName, " == nil {")
							generatedFile.P("req.", fullFieldName, " = &", field.Message.GoIdent, "{}")
							generatedFile.P("}")
						}
						field := internal.FindField(names[len(names)-1], inMessage)
						if field == nil {
							generatedFile.P("// 142")
							return errNotFoundField
						}
						var fieldNames []string
						for _, p := range fields {
							fieldNames = append(fieldNames, p.GoName)
						}
						fieldNames = append(fieldNames, field.GoName)
						fullFieldName := strings.Join(fieldNames, ".")

						fmtSeg := []any{internal.FmtPackage.Ident("Sprintf"), "(", strconv.Quote(template)}
						for _, namedPathParameter := range namedPathParameters {
							fmtSeg = append(fmtSeg, ", vars[", strconv.Quote(namedPathParameter), "]")
						}
						switch field.Desc.Kind() {
						case protoreflect.StringKind:
						case protoreflect.BytesKind:
							fmtSeg = slicex.Prepend(append(fmtSeg, ")"), "[]byte(")
						default:
							return errInvalidType
						}
						line := []any{"req.", fullFieldName, " = "}
						line = append(line, fmtSeg...)
						line = append(line, ")")
						generatedFile.P(line...)
					}
					// 普通路径参数设值
					// 去掉命名路径参数
					pathParameters := slicex.Difference(pathParameters, namedPathParameters)
					for _, pathParameter := range pathParameters {
						field := internal.FindField(pathParameter, endpoint.Input())
						generatedFile.P("req.", field.GoName, " = vars[", strconv.Quote(pathParameter), "]")
					}
				}
			}

			// query arguments
			{

			}

			generatedFile.P("return nil, nil")
			generatedFile.P("},")
			generatedFile.P("func(ctx ", internal.ContextPackage.Ident("Context"), ", w ", internal.HttpPackage.Ident("ResponseWriter"), ", resp any) error {")
			generatedFile.P("return nil")
			generatedFile.P("},")
			generatedFile.P("opts...,")
			generatedFile.P("))")

		}
	}
	generatedFile.P("return r")
	generatedFile.P("}")
	generatedFile.P()
	return nil
}

//
//func (f *Generator) GenerateClient(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
//	generatedFile.P("type ", service.UnexportedGRPCClientName(), " struct {")
//	for _, endpoint := range service.Endpoints {
//		generatedFile.P(endpoint.UnexportedName(), " ", endpointPackage.Ident("Endpoint"))
//	}
//	generatedFile.P("}")
//	generatedFile.P()
//	for _, endpoint := range service.Endpoints {
//		generatedFile.P("func (c *", service.UnexportedGRPCClientName(), ") ", endpoint.Name(), "(ctx ", contextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
//		generatedFile.P("rep, err := c.", endpoint.UnexportedName(), "(ctx, request)")
//		generatedFile.P("if err != nil {")
//		generatedFile.P("return nil, err")
//		generatedFile.P("}")
//		generatedFile.P("return rep.(*", endpoint.OutputGoIdent(), "), nil")
//		generatedFile.P("}")
//		generatedFile.P()
//	}
//	return nil
//}
//
//func (f *Generator) GenerateNewClient(service *internal.Service, generatedFile *protogen.GeneratedFile) error {
//	generatedFile.P("func New", service.GRPCClientName(), "(")
//	generatedFile.P("conn *", grpcPackage.Ident("ClientConn"), ",")
//	generatedFile.P("mdw []", endpointPackage.Ident("Middleware"), ",")
//	generatedFile.P("opts ...", grpcTransportPackage.Ident("ClientOption"), ",")
//	generatedFile.P(") interface {")
//	for _, endpoint := range service.Endpoints {
//		generatedFile.P(endpoint.Name(), "(ctx ", contextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error)")
//	}
//	generatedFile.P("} {")
//	generatedFile.P("return &", service.UnexportedGRPCClientName(), "{")
//	for _, endpoint := range service.Endpoints {
//		_ = endpoint
//		generatedFile.P(endpoint.UnexportedName(), ":    ", endpointxPackage.Ident("Chain"), "(", grpcTransportPackage.Ident("NewClient"), "(conn, ", strconv.Quote(service.FullName()), ",", strconv.Quote(endpoint.Name()), ", func(_ ", contextPackage.Ident("Context"), ", v any) (any, error) { return v, nil }", ", ", "func(_ ", contextPackage.Ident("Context"), ", v any) (any, error) { return v, nil }, ", endpoint.OutputGoIdent(), "{}, opts...).Endpoint(), mdw...),")
//	}
//	generatedFile.P("}")
//	generatedFile.P("}")
//	generatedFile.P()
//	return nil
//}
