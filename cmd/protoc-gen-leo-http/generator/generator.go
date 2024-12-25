package generator

import (
	"github.com/go-leo/leo/v3/cmd/internal"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-leo-http/generator/http"
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
	filename := file.GeneratedFilenamePrefix + "_leo.http.pb.go"
	g := f.Plugin.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-leo-http. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	services, err := internal.NewHttpServices(file)
	if err != nil {
		return err
	}

	exportedFunctionGenerator := http.ExportedFunctionGenerator{}
	for _, service := range services {
		if err := exportedFunctionGenerator.GenerateAppendRoutesFunc(service, g); err != nil {
			return err
		}
		if err := exportedFunctionGenerator.GenerateAppendServerFunc(service, g); err != nil {
			return err
		}
		if err := exportedFunctionGenerator.GenerateNewClientFunc(service, g); err != nil {
			return err
		}
	}

	serverTransportsGenerator := http.ServerTransportsGenerator{}
	serverRequestDecoderGenerator := http.ServerRequestDecoderGenerator{}
	serverResponseEncoderGenerator := http.ServerResponseEncoderGenerator{}
	clientTransportsGenerator := http.ClientTransportsGenerator{}
	clientRequestEncoderGenerator := http.ClientRequestEncoderGenerator{}
	clientResponseDecoderGenerator := http.ClientResponseDecoderGenerator{}
	for _, service := range services {
		if err := serverTransportsGenerator.GenerateTransports(service, g); err != nil {
			return err
		}
		if err := serverRequestDecoderGenerator.GenerateServerRequestDecoder(service, g); err != nil {
			return err
		}
		if err := serverResponseEncoderGenerator.GenerateServerResponseEncoder(service, g); err != nil {
			return err
		}
		if err := clientRequestEncoderGenerator.GenerateClientRequestEncoder(service, g); err != nil {
			return err
		}
		if err := clientResponseDecoderGenerator.GenerateClientResponseDecoder(service, g); err != nil {
			return err
		}
		if err := serverTransportsGenerator.GenerateTransportsImplements(service, g); err != nil {
			return err
		}
		if err := serverRequestDecoderGenerator.GenerateServerRequestDecoderImplements(service, g); err != nil {
			return err
		}
		if err := serverResponseEncoderGenerator.GenerateServerResponseEncoderImplements(service, g); err != nil {
			return err
		}
		if err := clientTransportsGenerator.GenerateTransports(service, g); err != nil {
			return err
		}
		if err := clientRequestEncoderGenerator.GenerateClientRequestEncoderImplements(service, g); err != nil {
			return err
		}
		if err := clientResponseDecoderGenerator.GenerateClientResponseDecoderImplements(service, g); err != nil {
			return err
		}
	}
	return nil
}
