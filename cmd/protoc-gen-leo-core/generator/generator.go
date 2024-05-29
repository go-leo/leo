package generator

import (
	"github.com/go-leo/leo/v3/cmd/internal"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo-core/generator/core"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo-core/generator/grpc"
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
	filename := file.GeneratedFilenamePrefix + "_leo.core.pb.go"
	g := f.Plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-grpc. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	coreGen, err := core.NewGenerator(f.Plugin, file)
	if err != nil {
		return err
	}
	if err := coreGen.Generate(g); err != nil {
		return err
	}

	grpcGen, err := grpc.NewGenerator(f.Plugin, file)
	if err != nil {
		return err
	}
	if err := grpcGen.Generate(g); err != nil {
		return err
	}
	return nil
}
