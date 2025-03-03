package internal

import (
	"errors"
	"fmt"
	"golang.org/x/mod/modfile"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"path/filepath"
	"strings"
)

type Service struct {
	file           *protogen.File
	protoService   *protogen.Service
	Endpoints      []*Endpoint
	CommandPath    string
	CommandPackage string
	QueryPath      string
	QueryPackage   string
}

func (s *Service) FullName() string {
	return string(s.protoService.Desc.FullName())
}

func (s *Service) Name() string {
	return s.protoService.GoName
}

func (s *Service) Unexported(name string) string {
	return strings.ToLower(name[:1]) + name[1:]
}

func (s *Service) ServiceName() string {
	return s.Name() + "Service"
}

func (s *Service) ServerEndpointsName() string {
	return s.Name() + "ServerEndpoints"
}

func (s *Service) ClientEndpointsName() string {
	return s.Name() + "ClientEndpoints"
}

func (s *Service) TransportsName() string {
	return s.Name() + "Transports"
}

func (s *Service) ClientTransportsName() string {
	return s.Name() + "ClientTransports"
}

func (s *Service) FactoriesName() string {
	return s.Name() + "Factories"
}

func (s *Service) EndpointersName() string {
	return s.Name() + "Endpointers"
}

func (s *Service) BalancersName() string {
	return s.Name() + "Balancers"
}

func (s *Service) ServerName() string {
	return s.Name() + "Server"
}

func (s *Service) ClientName() string {
	return s.Name() + "Client"
}

func (s *Service) ClientServiceName() string {
	return s.Name() + "ClientService"
}

// ------------------ Grpc ------------------

func (s *Service) GrpcServerName() string {
	return s.Name() + "GrpcServer"
}

func (s *Service) GrpcClientName() string {
	return s.Name() + "GrpcClient"
}

func (s *Service) GrpcServerTransportsName() string {
	return s.GrpcServerName() + "Transports"
}

func (s *Service) GrpcClientEndpointsName() string {
	return s.GrpcClientName() + "Endpoints"
}

func (s *Service) GrpcFactoriesName() string {
	return s.GrpcClientName() + "Factories"
}

func (s *Service) GrpcClientTransportsName() string {
	return s.GrpcClientName() + "Transports"
}

func (s *Service) UnimplementedServerName() string {
	return "Unimplemented" + s.protoService.GoName + "Server"
}

// ------------------ Http ------------------

func (s *Service) HttpServerName() string {
	return s.Name() + "HttpServer"
}

func (s *Service) HttpClientName() string {
	return s.Name() + "HttpClient"
}

func (s *Service) HttpServerTransportsName() string {
	return s.HttpServerName() + "Transports"
}

func (s *Service) HttpClientTransportsName() string {
	return s.HttpClientName() + "Transports"
}

func (s *Service) HttpRoutesName() string {
	return s.Name() + "HttpRoutes"
}

func (s *Service) HttpServerRoutesName() string {
	return s.Name() + "HttpServerRoutes"
}

func (s *Service) HttpServerRequestDecoderName() string {
	return s.HttpServerName() + "RequestDecoder"
}

func (s *Service) HttpServerResponseEncoderName() string {
	return s.HttpServerName() + "ResponseEncoder"
}

func (s *Service) HttpClientRequestEncoderName() string {
	return s.HttpClientName() + "RequestEncoder"
}

func (s *Service) HttpClientResponseDecoderName() string {
	return s.HttpClientName() + "ResponseDecoder"
}

// ------------------ CQRS ------------------

func (s *Service) CQRSName() string {
	return s.protoService.GoName + "CqrsService"
}

func (s *Service) AssemblerName() string {
	return s.protoService.GoName + "Assembler"
}

func (s *Service) BusName() string {
	return s.protoService.GoName + "Bus"
}

func (s *Service) ParseCqrs(file *protogen.File) error {
	// 查找go.mod
	filePath := s.file.Desc.Path()
	pkgAbs, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	dir := filepath.Dir(pkgAbs)
	var found bool
	for {
		stat, _ := os.Stat(filepath.Join(dir, "go.mod"))
		if stat != nil {
			found = true
			break
		}
		par := filepath.Dir(dir)
		if par == dir {
			break
		}
		dir = par
	}
	if !found {
		return errors.New("failed to found go.mod")
	}
	data, err := os.ReadFile(filepath.Join(dir, "go.mod"))
	if err != nil {
		return err
	}
	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return err
	}
	_ = modFile
	//modulePath := modFile.Module.Mod.Path
	//
	//s.CommandPath = filepath.Join(wd, "command")
	//s.CommandPackage = path.Join(file.GoImportPath.String(), "command")
	//s.QueryPath = filepath.Join(wd, "query")
	//s.QueryPackage = path.Join(file.GoImportPath.String(), "query")
	return nil
}

func NewServices(file *protogen.File) ([]*Service, error) {
	var services []*Service
	for _, pbService := range file.Services {
		service := &Service{
			file:         file,
			protoService: pbService,
		}
		if err := service.ParseCqrs(file); err != nil {
			return nil, err
		}
		var endpoints []*Endpoint
		for _, pbMethod := range pbService.Methods {
			endpoint := &Endpoint{
				protoMethod: pbMethod,
			}
			if endpoint.IsStreaming() {
				return nil, fmt.Errorf("leo: unsupport stream httpMethod, %s", endpoint.FullName())
			}
			if err := endpoint.ParseHttpRule(); err != nil {
				return nil, err
			}
			endpoint.ParseCqrs()
			endpoints = append(endpoints, endpoint)
		}
		service.Endpoints = endpoints
		services = append(services, service)
	}
	return services, nil
}
