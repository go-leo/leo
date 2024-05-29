package http

import (
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo/generator/internal"
	"google.golang.org/protobuf/compiler/protogen"
)

type Generator struct {
	Plugin   *protogen.Plugin
	File     *protogen.File
	Services []*internal.Service
}

func NewGenerator(plugin *protogen.Plugin, file *protogen.File) (*Generator, error) {
	services, err := internal.NewHttpServices(file)
	if err != nil {
		return nil, err
	}
	return &Generator{Plugin: plugin, File: file, Services: services}, nil
}

func (f *Generator) Generate(g *protogen.GeneratedFile) error {
	server := ServerGenerator{}
	client := ClientGenerator{}

	for _, service := range f.Services {
		if err := server.GenerateTransports(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := client.GenerateTransports(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := server.GenerateImplementedTransports(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := client.GenerateImplementedTransports(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := server.GenerateServer(service, g); err != nil {
			return err
		}
	}

	for _, service := range f.Services {
		if err := client.GenerateClient(service, g); err != nil {
			return err
		}
	}
	return nil
}
