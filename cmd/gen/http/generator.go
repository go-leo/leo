package http

import (
	"github.com/go-leo/leo/v3/cmd/gen/internal"
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
	if len(f.Services) <= 0 {
		return nil
	}
	file := f.File
	filename := file.GeneratedFilenamePrefix + "_leo.http.pb.go"
	g := f.Plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-leo. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	services, err := internal.NewServices(file)
	if err != nil {
		return err
	}

	for _, service := range services {
		functionGenerator := FunctionGenerator{
			service: service,
			g:       g,
		}
		functionGenerator.GenerateAppendRoutesFunc()
		functionGenerator.GenerateAppendServerFunc()
		functionGenerator.GenerateNewClientFunc()

		serverTransportsGenerator := ServerTransportsGenerator{
			service: service,
			g:       g,
		}
		serverTransportsGenerator.GenerateTransports()

		serverRequestDecoderGenerator := ServerRequestDecoderGenerator{
			service: service,
			g:       g,
		}
		serverRequestDecoderGenerator.GenerateServerRequestDecoder()

		responseEncoderGenerator := ResponseEncoderGenerator{
			service: service,
			g:       g,
		}
		responseEncoderGenerator.GenerateResponseEncoder()

		requestEncoderGenerator := RequestEncoderGenerator{
			service: service,
			g:       g,
		}
		requestEncoderGenerator.GenerateRequestEncoder()

		responseDecoderGenerator := ResponseDecoderGenerator{
			service: service,
			g:       g,
		}
		responseDecoderGenerator.GenerateClientResponseDecoder()

		serverTransportsGenerator.GenerateTransportsImplements()

		serverRequestDecoderGenerator.GenerateServerRequestDecoderImplements()

		if err := responseEncoderGenerator.GenerateServerResponseEncoderImplements(); err != nil {
			return err
		}

		clientTransportsGenerator := ClientTransportsGenerator{
			service: service,
			g:       g,
		}
		clientTransportsGenerator.GenerateTransports()

		requestEncoderGenerator.GenerateClientRequestEncoderImplements()

		if err := responseDecoderGenerator.GenerateClientResponseDecoderImplements(); err != nil {
			return err
		}
	}

	return nil
}
