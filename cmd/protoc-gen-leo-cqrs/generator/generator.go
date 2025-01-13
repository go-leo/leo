package generator

import (
	"github.com/go-leo/gox/stringx"
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"path"
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
	return f.GenerateFile()
}

func (f *Generator) GenerateFile() error {
	file := f.File
	filename := file.GeneratedFilenamePrefix + "_leo.cqrs.pb.go"
	g := f.Plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-leo-cqrs. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, service := range f.Services {
		if service.Command == nil && service.Query == nil {
			continue
		}
		if err := f.GenerateEndpoints(service); err != nil {
			return err
		}
		if err := f.GenerateAssembler(service, g); err != nil {
			return err
		}
		if err := f.GenerateCQRSService(service, g); err != nil {
			return err
		}
	}
	return nil
}

func (f *Generator) GenerateEndpoints(service *internal.Service) error {
	for _, endpoint := range service.Endpoints {
		if err := f.GenerateEndpoint(service, endpoint); err != nil {
			return err
		}
	}
	return nil
}

func (f *Generator) GenerateEndpoint(service *internal.Service, endpoint *internal.Endpoint) error {
	switch {
	case endpoint.IsCommand():
		return f.GenerateCommand(service, endpoint)
	case endpoint.IsQuery():
		return f.GenerateQuery(service, endpoint)
	}
	return nil
}

func (f *Generator) GenerateCommand(service *internal.Service, endpoint *internal.Endpoint) error {
	filename := path.Join(service.Command.RelPath(), stringx.SnackCase(endpoint.Name())+".go")
	_, err := os.Stat(filename)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}
	g := f.Plugin.NewGeneratedFile(filename, f.File.GoImportPath)
	g.P("package ", service.Command.Name())
	g.P("type ", endpoint.ArgsName(), " struct {")
	g.P("}")
	g.P()
	g.P("type ", endpoint.Name(), " ", internal.CqrsPackage.Ident("CommandHandler"), "[*", endpoint.ArgsName(), "]")
	g.P()
	g.P("func New", endpoint.Name(), "() ", endpoint.Name(), " {")
	g.P("return &", endpoint.Unexported(endpoint.Name()), "{}")
	g.P("}")
	g.P()
	g.P("type ", endpoint.Unexported(endpoint.Name()), " struct {")
	g.P("}")
	g.P()
	g.P("func (h *", endpoint.Unexported(endpoint.Name()), ") Handle(ctx ", internal.ContextPackage.Ident("Context"), ", args *", endpoint.ArgsName(), ") (", internal.MetadataxPackage.Ident("Metadata"), ", error) {")
	g.P(internal.Comments("TODO implement me"))
	g.P("panic(", strconv.Quote("implement me"), ")")
	g.P("}")
	return nil
}

func (f *Generator) GenerateQuery(service *internal.Service, endpoint *internal.Endpoint) error {
	filename := path.Join(service.Query.RelPath(), stringx.SnackCase(endpoint.Name())+".go")
	_, err := os.Stat(filename)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}
	g := f.Plugin.NewGeneratedFile(filename, f.File.GoImportPath)
	g.P("package ", service.Query.Name())
	g.P("type ", endpoint.ArgsName(), " struct {")
	g.P("}")
	g.P()
	g.P("type ", endpoint.ResName(), " struct {")
	g.P("}")
	g.P()
	g.P("type ", endpoint.Name(), " ", internal.CqrsPackage.Ident("QueryHandler"), "[*", endpoint.ArgsName(), ", *", endpoint.ResName(), "]")
	g.P()
	g.P("func New", endpoint.Name(), "() ", endpoint.Name(), " {")
	g.P("return &", endpoint.Unexported(endpoint.Name()), "{}")
	g.P("}")
	g.P()
	g.P("type ", endpoint.Unexported(endpoint.Name()), " struct {")
	g.P("}")
	g.P()
	g.P("func (h *", endpoint.Unexported(endpoint.Name()), ") Handle(ctx ", internal.ContextPackage.Ident("Context"), ", args *", endpoint.ArgsName(), ") (*", endpoint.ResName(), ", error) {")
	g.P(internal.Comments("TODO implement me"))
	g.P("panic(", strconv.Quote("implement me"), ")")
	g.P("}")
	return nil
}

func (f *Generator) GenerateAssembler(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P(internal.Comments(service.AssemblerName() + " responsible for completing the transformation between domain model objects and DTOs"))
	g.P("type ", service.AssemblerName(), " interface {")
	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			g.P(internal.Comments("From" + endpoint.RequestName() + " convert request to command arguments"))
			g.P("From", endpoint.RequestName(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", protogen.GoImportPath(service.Command.FullName()).Ident(endpoint.ArgsName()), ", ", internal.ContextPackage.Ident("Context"), ", error)")
		case endpoint.IsQuery():
			g.P(internal.Comments("From" + endpoint.RequestName() + " convert request to query arguments"))
			g.P("From", endpoint.RequestName(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", protogen.GoImportPath(service.Query.FullName()).Ident(endpoint.ArgsName()), ", ", internal.ContextPackage.Ident("Context"), ", error)")
			g.P(internal.Comments("To" + endpoint.ResponseName() + " convert query result to response"))
			g.P("To", endpoint.ResponseName(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ", res *", protogen.GoImportPath(service.Query.FullName()).Ident(endpoint.ResName()), ") (*", endpoint.OutputGoIdent(), ", error)")
		}
	}
	g.P("}")
	return nil
}

func (f *Generator) GenerateCQRSService(service *internal.Service, g *protogen.GeneratedFile) error {
	g.P(internal.Comments(service.Unexported(service.CQRSName()) + " implement the " + service.ServiceName() + " with CQRS pattern"))
	g.P("type ", service.Unexported(service.CQRSName()), " struct {")
	g.P("bus       ", internal.CqrsPackage.Ident("Bus"))
	g.P("assembler ", service.AssemblerName())
	g.P("}")
	g.P()

	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			g.P("func (svc *", service.Unexported(service.CQRSName()), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
			g.P("command, ctx, err := svc.assembler.From", endpoint.Name(), "Request(ctx, request)")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("if err := svc.bus.Exec(ctx, command);err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("return new(", endpoint.OutputGoIdent(), "), nil")
			g.P("}")
			g.P()
		case endpoint.IsQuery():
			g.P("func (svc *", service.Unexported(service.CQRSName()), ") ", endpoint.Name(), "(ctx ", internal.ContextPackage.Ident("Context"), ", request *", endpoint.InputGoIdent(), ") (*", endpoint.OutputGoIdent(), ", error){")
			g.P("args, ctx, err := svc.assembler.From", endpoint.Name(), "Request(ctx, request)")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("res, err := svc.bus.Query(ctx, args)")
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			g.P("return svc.assembler.To", endpoint.Name(), "Response(ctx, request, res.(*", protogen.GoImportPath(service.Query.FullName()).Ident(endpoint.Name()+"Res"), "))")
			g.P("}")
			g.P()
		}
	}

	g.P("func New", service.CQRSName(), "(")
	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			importPath := protogen.GoImportPath(service.Command.FullName())
			g.P(endpoint.Unexported(endpoint.Name()), " ", importPath.Ident(endpoint.Name()), ",")
		case endpoint.IsQuery():
			importPath := protogen.GoImportPath(service.Query.FullName())
			g.P(endpoint.Unexported(endpoint.Name()), " ", importPath.Ident(endpoint.Name()), ",")
		}
	}
	g.P("assembler ", service.AssemblerName(), ",")
	g.P(") (", service.ServiceName(), ", error) {")
	g.P("var bus ", internal.CqrsPackage.Ident("SampleBus"))
	for _, endpoint := range service.Endpoints {
		switch {
		case endpoint.IsCommand():
			g.P("if err := bus.RegisterCommand(", endpoint.Unexported(endpoint.Name()), "); err != nil {")
			g.P("return nil, err")
			g.P("}")
		case endpoint.IsQuery():
			g.P("if err := bus.RegisterQuery(", endpoint.Unexported(endpoint.Name()), "); err != nil {")
			g.P("return nil, err")
			g.P("}")
		}
	}
	g.P("return &", service.Unexported(service.CQRSName()), "{")
	g.P("bus: &bus, ")
	g.P("assembler: assembler,")
	g.P("}, nil")
	g.P("}")
	g.P()
	return nil
}
