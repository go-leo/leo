package cqrs

import (
	"github.com/go-leo/gox/stringx"
	"github.com/go-leo/leo/v3/cmd/gen/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"path/filepath"
	"strconv"
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
	// 没有定义服务，则不生成代码
	if len(f.Services) <= 0 {
		return nil
	}
	filename := f.File.GeneratedFilenamePrefix + "_leo.cqrs.pb.go"
	g := f.Plugin.NewGeneratedFile(filename, f.File.GoImportPath)
	g.P("// Code generated by protoc-gen-go-leo. DO NOT EDIT.")
	g.P()
	g.P("package ", f.File.GoPackageName)
	g.P()

	for _, service := range f.Services {
		f.PrintNewCQRSService(g, service)
	}

	for _, service := range f.Services {
		f.PrintCQRSService(g, service)
		commandEndpoints := service.CommandEndpoints()
		queryEndpoints := service.QueryEndpoints()
		f.PrintCommand(g, commandEndpoints)
		f.PrintQuery(g, queryEndpoints)
		f.PrintUnimplementedCommand(g, commandEndpoints)
		f.PrintUnimplementedQuery(g, queryEndpoints)
	}

	for _, service := range f.Services {
		for _, endpoint := range service.Endpoints {
			dir := filepath.Dir(f.File.GeneratedFilenamePrefix)
			path := "cq"
			filename := stringx.JSONSnakeCase(endpoint.Unexported(endpoint.Name())) + "_leo.query.pb.go"
			g.P("// ", filepath.Join(dir, path, filename))
		}
	}

	f.GenerateCQ()
	return nil
}

func (f *Generator) PrintCommandHandler(g *protogen.GeneratedFile, commandEndpoints []*internal.Endpoint) {
	if len(commandEndpoints) <= 0 {
		return
	}
	g.P("type (")
	for _, endpoint := range commandEndpoints {
		g.P(endpoint.HandlerName(), " ", internal.CommandHandler, "[", endpoint.CommandName(), "]")
	}
	g.P(")")
	g.P()
}

func (f *Generator) PrintQueryHandler(g *protogen.GeneratedFile, queryEndpoints []*internal.Endpoint) {
	if len(queryEndpoints) <= 0 {
		return
	}
	g.P("type (")
	for _, endpoint := range queryEndpoints {
		g.P(endpoint.HandlerName(), " ", internal.QueryHandler, "[", endpoint.QueryName(), ", ", endpoint.ResultName(), "]")
	}
	g.P(")")
	g.P()
}

func (f *Generator) PrintNewCQRSService(g *protogen.GeneratedFile, service *internal.Service) {
	g.P("func New", service.CQRSName(), "[")
	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			g.P(endpoint.CommandTypeName(), " ", endpoint.CommandName(), ", ")
		case endpoint.IsQuery():
			g.P(endpoint.QueryTypeName(), " ", endpoint.QueryName(), ", ", endpoint.ResultTypeName(), " ", endpoint.ResultName(), ", ")
		}
	}
	g.P("](")
	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			g.P(endpoint.Unexported(endpoint.HandlerName()), " ", internal.CommandHandler, "[", endpoint.CommandTypeName(), "],")
		case endpoint.IsQuery():
			g.P(endpoint.Unexported(endpoint.HandlerName()), " ", internal.QueryHandler, "[", endpoint.QueryTypeName(), ", ", endpoint.ResultTypeName(), "],")
		}
	}
	g.P(") (", service.ServiceName(), ", error) {")
	g.P("var bus ", internal.CqrsPackage.Ident("SampleBus"))
	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			g.P("if err := bus.RegisterCommand(", endpoint.Unexported(endpoint.HandlerName()), "); err != nil {")
			g.P("return nil, err")
			g.P("}")
		case endpoint.IsQuery():
			g.P("if err := bus.RegisterQuery(", endpoint.Unexported(endpoint.HandlerName()), "); err != nil {")
			g.P("return nil, err")
			g.P("}")
		}
	}
	g.P("return &", service.Unexported(service.CQRSName()), "[")
	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			g.P(endpoint.CommandTypeName(), ", ")
		case endpoint.IsQuery():
			g.P(endpoint.QueryTypeName(), ", ", endpoint.ResultTypeName(), ", ")
		}
	}
	g.P("]{bus: &bus}, nil")
	g.P("}")
	g.P()
}

func (f *Generator) PrintCQRSService(g *protogen.GeneratedFile, service *internal.Service) {
	g.P("type ", service.Unexported(service.CQRSName()), "[")
	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			g.P(endpoint.CommandTypeName(), " ", endpoint.CommandName(), ", ")
		case endpoint.IsQuery():
			g.P(endpoint.QueryTypeName(), " ", endpoint.QueryName(), ", ", endpoint.ResultTypeName(), " ", endpoint.ResultName(), ", ")
		}
	}
	g.P("] struct {")
	g.P("bus ", internal.Bus)
	g.P("}")
	g.P()
	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			f.PrintCommandMethod(g, service, endpoint)
		case endpoint.IsQuery():
			f.PrintQueryMethod(g, service, endpoint)
		}
	}
}

func (f *Generator) PrintCommandMethod(g *protogen.GeneratedFile, service *internal.Service, endpoint *internal.Endpoint) {
	g.P("func (svc *", service.Unexported(service.CQRSName()), "[")
	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			g.P(endpoint.CommandTypeName(), ", ")
		case endpoint.IsQuery():
			g.P(endpoint.QueryTypeName(), ", ", endpoint.ResultTypeName(), ", ")
		}
	}
	g.P("]) ", endpoint.Name(), "(ctx ", internal.Context, ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
	g.P("var command ", endpoint.CommandTypeName())
	g.P("cmd, ctx, err := command.From(ctx, request)")
	g.P("if err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P("if err := svc.bus.Exec(ctx, cmd); err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P("return new(", endpoint.OutputGoIdent(), "), nil")
	g.P("}")
	g.P()
}

func (f *Generator) PrintQueryMethod(g *protogen.GeneratedFile, service *internal.Service, endpoint *internal.Endpoint) {
	g.P("func (svc *", service.Unexported(service.CQRSName()), "[")
	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			g.P(endpoint.CommandTypeName(), ", ")
		case endpoint.IsQuery():
			g.P(endpoint.QueryTypeName(), ", ", endpoint.ResultTypeName(), ", ")
		}
	}
	g.P("]) ", endpoint.Name(), "(ctx ", internal.Context, ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
	g.P("var query ", endpoint.QueryTypeName())
	g.P("q, ctx, err := query.From(ctx, request)")
	g.P("if err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P("r, err := svc.bus.Query(ctx, q)")
	g.P("if err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P("return r.(", endpoint.ResultTypeName(), ").To(ctx)")
	g.P("}")
	g.P()
}

func (f *Generator) PrintCommand(g *protogen.GeneratedFile, commandEndpoints []*internal.Endpoint) {
	if len(commandEndpoints) <= 0 {
		return
	}
	g.P("type (")
	for _, endpoint := range commandEndpoints {
		g.P(endpoint.CommandName(), " interface {")
		g.P(endpoint.IsCommandMethod(), "()")
		g.P("From(", internal.Context, ", *", endpoint.InputGoIdent(), ") (", endpoint.CommandName(), ", ", internal.Context, ", error)")
		g.P("}")
		g.P(endpoint.UnimplementedCommandName(), " struct{}")
		g.P()
	}
	g.P(")")
	g.P()
}

func (f *Generator) PrintQuery(g *protogen.GeneratedFile, queryEndpoints []*internal.Endpoint) {
	if len(queryEndpoints) <= 0 {
		return
	}
	g.P("type (")
	for _, endpoint := range queryEndpoints {
		g.P(endpoint.QueryName(), " interface {")
		g.P(endpoint.IsQueryMethod(), "()")
		g.P("From(", internal.Context, ", *", endpoint.InputGoIdent(), ") (", endpoint.QueryName(), ", ", internal.Context, ", error)")
		g.P("}")
		g.P(endpoint.ResultName(), " interface {")
		g.P(endpoint.IsResultMethod(), "()")
		g.P("To(", internal.Context, ") (*", endpoint.OutputGoIdent(), ", error)")
		g.P("}")
		g.P(endpoint.UnimplementedQueryName(), " struct{}")
		g.P(endpoint.UnimplementedResultName(), " struct{}")
		g.P()
	}
	g.P(")")
	g.P()
}

func (f *Generator) PrintUnimplementedCommand(g *protogen.GeneratedFile, commandEndpoints []*internal.Endpoint) {
	if len(commandEndpoints) <= 0 {
		return
	}
	for _, endpoint := range commandEndpoints {
		g.P("func (", endpoint.UnimplementedCommandName(), ") ", endpoint.IsCommandMethod(), "(){}")
		g.P("func (", endpoint.UnimplementedCommandName(), ") ", "From(", internal.Context, ", *", endpoint.InputGoIdent(), ") (", endpoint.CommandName(), ", ", internal.Context, ", error) {")
		g.P("panic(", strconv.Quote("implement me"), ")")
		g.P("return nil, nil, nil")
		g.P("}")
	}
}

func (f *Generator) PrintUnimplementedQuery(g *protogen.GeneratedFile, queryEndpoints []*internal.Endpoint) {
	if len(queryEndpoints) <= 0 {
		return
	}
	for _, endpoint := range queryEndpoints {
		g.P("func (", endpoint.UnimplementedQueryName(), ") ", endpoint.IsQueryMethod(), "(){}")
		g.P("func (", endpoint.UnimplementedQueryName(), ") From(", internal.Context, ", *", endpoint.InputGoIdent(), ") (", endpoint.QueryName(), ", ", internal.Context, ", error)", " {")
		g.P("panic(", strconv.Quote("implement me"), ")")
		g.P("return nil, nil, nil")
		g.P("}")
		g.P("func (", endpoint.UnimplementedResultName(), ") ", endpoint.IsResultMethod(), "(){}")
		g.P("func (", endpoint.UnimplementedResultName(), ") To(", internal.Context, ") (*", endpoint.OutputGoIdent(), ", error) {")
		g.P("panic(", strconv.Quote("implement me"), ")")
		g.P("return nil, nil")
		g.P("}")
		g.P()
	}
}

func (f *Generator) GenerateCQ() {
	for _, service := range f.Services {
		for _, endpoint := range service.Endpoints {
			switch {
			case endpoint.IsCommand():
				f.GenerateCommand(endpoint)
			case endpoint.IsQuery():
				f.GenerateQuery(endpoint)
			}
		}
	}
}

func (f *Generator) GenerateCommand(endpoint *internal.Endpoint) {
	filename := filepath.Join(filepath.Dir(f.File.GeneratedFilenamePrefix), "cq", stringx.JSONSnakeCase(endpoint.Unexported(endpoint.Name()))+"_leo.command.pb.go")
	g := f.Plugin.NewGeneratedFile(filename, f.File.GoImportPath+"/cq")
	g.P("// Code generated by protoc-gen-go-leo. DO NOT EDIT.")
	g.P("// If you want edit it, can move this file to another directory.")
	g.P()
	g.P("package cq")
	g.P()
	g.P("var _ ", endpoint.HandlerName(), " = (*", endpoint.Unexported(endpoint.HandlerName()), ")(nil)")
	g.P()
	g.P("type ", endpoint.HandlerName(), " ", internal.CqrsPackage.Ident("CommandHandler"), "[", endpoint.CommandName(), "]")
	g.P()
	g.P("type ", endpoint.CommandName(), " struct {")
	g.P(f.File.GoImportPath.Ident(endpoint.UnimplementedCommandName()))
	g.P("}")
	g.P()
	g.P("func (", endpoint.CommandName(), ") From(ctx ", internal.Context, ", req *", endpoint.InputGoIdent(), ") (", f.File.GoImportPath.Ident(endpoint.CommandName()), ", ", internal.Context, ", error) {")
	g.P("panic(", strconv.Quote("implement me"), ")")
	g.P("return ", endpoint.CommandName(), "{}, ctx, nil")
	g.P("}")
	g.P()
	g.P("func New", endpoint.HandlerName(), "() ", endpoint.HandlerName(), " {")
	g.P("return &", endpoint.Unexported(endpoint.HandlerName()), "{}")
	g.P("}")
	g.P()
	g.P("type ", endpoint.Unexported(endpoint.HandlerName()), " struct {")
	g.P("}")
	g.P()
	g.P("func (h *", endpoint.Unexported(endpoint.HandlerName()), ") Handle(ctx ", internal.Context, ", cmd ", endpoint.CommandName(), ") error {")
	g.P(internal.Comments("TODO implement me"))
	g.P("panic(", strconv.Quote("implement me"), ")")
	g.P("}")
}

func (f *Generator) GenerateQuery(endpoint *internal.Endpoint) {
	filename := filepath.Join(filepath.Dir(f.File.GeneratedFilenamePrefix), "cq", stringx.JSONSnakeCase(endpoint.Unexported(endpoint.Name()))+"_leo.query.pb.go")
	g := f.Plugin.NewGeneratedFile(filename, f.File.GoImportPath+"/cq")
	g.P("// Code generated by protoc-gen-go-leo. DO NOT EDIT.")
	g.P("// If you want edit it, can move this file to another directory.")
	g.P()
	g.P("package cq")
	g.P()
	g.P("var _ ", endpoint.HandlerName(), " = (*", endpoint.Unexported(endpoint.HandlerName()), ")(nil)")
	g.P()
	g.P("type ", endpoint.HandlerName(), " ", internal.CqrsPackage.Ident("QueryHandler"), "[", endpoint.QueryName(), ", ", endpoint.ResultName(), "]")
	g.P()
	g.P("type ", endpoint.QueryName(), " struct {")
	g.P(f.File.GoImportPath.Ident(endpoint.UnimplementedQueryName()))
	g.P("}")
	g.P()
	g.P("func (", endpoint.QueryName(), ") From(ctx ", internal.Context, ", req*", endpoint.InputGoIdent(), ") (", f.File.GoImportPath.Ident(endpoint.QueryName()), ", ", internal.Context, ", error)", " {")
	g.P("panic(", strconv.Quote("implement me"), ")")
	g.P("return ", endpoint.QueryName(), "{}, ctx, nil")
	g.P("}")
	g.P()
	g.P("type ", endpoint.ResultName(), " struct {")
	g.P(f.File.GoImportPath.Ident(endpoint.UnimplementedResultName()))
	g.P("}")
	g.P()
	g.P("func (r ", endpoint.ResultName(), ") To(ctx ", internal.Context, ") (*", endpoint.OutputGoIdent(), ", error) {")
	g.P("panic(", strconv.Quote("implement me"), ")")
	g.P("return &", endpoint.OutputGoIdent(), "{}, nil")
	g.P("}")
	g.P()
	g.P("func New", endpoint.HandlerName(), "() ", endpoint.HandlerName(), " {")
	g.P("return &", endpoint.Unexported(endpoint.HandlerName()), "{}")
	g.P("}")
	g.P()
	g.P("type ", endpoint.Unexported(endpoint.HandlerName()), " struct {")
	g.P("}")
	g.P()
	g.P("func (h *", endpoint.Unexported(endpoint.HandlerName()), ") Handle(ctx ", internal.Context, ", q ", endpoint.QueryName(), ") (", endpoint.ResultName(), ", error) {")
	g.P(internal.Comments("TODO implement me"))
	g.P("panic(", strconv.Quote("implement me"), ")")
	g.P("}")
	return
}
