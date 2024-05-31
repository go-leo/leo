package generator

import (
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo/generator/core"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo/generator/cqrs"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo/generator/grpc"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo/generator/http"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo/generator/internal"
	"google.golang.org/protobuf/compiler/protogen"
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
	filename := file.GeneratedFilenamePrefix + "_leo.pb.go"
	g := f.Plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-grpc. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	g.P("// =========================== endpoints ===========================")
	g.P()
	coreGen, err := core.NewGenerator(f.Plugin, file)
	if err != nil {
		return err
	}
	if err := coreGen.Generate(g); err != nil {
		return err
	}

	g.P("// =========================== cqrs ===========================")
	g.P()
	cqrsGen, err := cqrs.NewGenerator(f.Plugin, file)
	if err != nil {
		return err
	}
	if err := cqrsGen.Generate(g); err != nil {
		return err
	}

	g.P("// =========================== grpc transports ===========================")
	g.P()
	grpcGen, err := grpc.NewGenerator(f.Plugin, file)
	if err != nil {
		return err
	}
	if err := grpcGen.Generate(g); err != nil {
		return err
	}

	g.P("// =========================== http transports ===========================")
	g.P()
	httpGen, err := http.NewGenerator(f.Plugin, file)
	if err != nil {
		return err
	}
	if err := httpGen.Generate(g); err != nil {
		return err
	}
	return nil
}