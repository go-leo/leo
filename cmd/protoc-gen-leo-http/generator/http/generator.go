package http

import (
	"github.com/go-leo/leo/v3/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"strconv"
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

func (f *Generator) GenerateRoutes(g *protogen.GeneratedFile) error {
	for _, service := range f.Services {
		g.P("func append", service.HttpRoutesName(), "(router *", internal.MuxPackage.Ident("Router"), ") *", internal.MuxPackage.Ident("Router"), "{")
		for _, endpoint := range service.Endpoints {
			httpRule := endpoint.HttpRule()
			// 调整路径，来适应 github.com/gorilla/mux 路由规则
			path, _, _, _ := httpRule.RegularizePath(httpRule.Path())
			g.P("router.NewRoute().Name(", strconv.Quote(endpoint.FullName()), ").Methods(", strconv.Quote(httpRule.Method()), ").Path(", strconv.Quote(path), ")")
		}
		g.P("return router")
		g.P("}")
	}
	return nil
}

func (f *Generator) GenerateServer(g *protogen.GeneratedFile) error {
	server := ServerGenerator{}
	for _, service := range f.Services {
		if err := server.GenerateServer(service, g); err != nil {
			return err
		}
	}
	return nil
}

func (f *Generator) GenerateClient(g *protogen.GeneratedFile) error {
	client := ClientGenerator{}
	for _, service := range f.Services {
		if err := client.GenerateTransports(service, g); err != nil {
			return err
		}
		if err := client.GenerateTransportsV2(service, g); err != nil {
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

func (f *Generator) GenerateTransport(g *protogen.GeneratedFile) error {
	server := ServerGenerator{}
	client := ClientGenerator{}
	for _, service := range f.Services {
		for _, endpoint := range service.Endpoints {
			if err := server.GenerateServerTransport(service, g, endpoint); err != nil {
				return err
			}
			if err := client.GenerateClientTransport(service, g, endpoint); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *Generator) GenerateCoder(g *protogen.GeneratedFile) error {
	server := ServerGenerator{}
	client := ClientGenerator{}
	for _, service := range f.Services {
		for _, endpoint := range service.Endpoints {
			if err := server.PrintDecodeRequestFunc(g, endpoint); err != nil {
				return err
			}
			if err := client.PrintEncodeRequestFunc(g, endpoint); err != nil {
				return err
			}
			if err := server.PrintEncodeResponseFunc(g, endpoint); err != nil {
				return err
			}
			if err := client.PrintDecodeResponseFunc(g, endpoint); err != nil {
				return err
			}
		}
	}
	return nil
}
