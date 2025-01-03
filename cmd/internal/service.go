package internal

import (
	"fmt"
	"github.com/go-leo/leo/v3/cqrs"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"os"
	"path/filepath"
	"strings"
)

type Service struct {
	ProtoService *protogen.Service
	Endpoints    []*Endpoint

	Command *Package
	Query   *Package
}

func (s *Service) FullName() string {
	return string(s.ProtoService.Desc.FullName())
}

func (s *Service) Name() string {
	return s.ProtoService.GoName
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

func (s *Service) HttpServerName() string {
	return s.Name() + "HttpServer"
}

func (s *Service) ClientServiceName() string {
	return s.Name() + "ClientService"
}

func (s *Service) HttpClientName() string {
	return s.Name() + "HttpClient"
}

func (s *Service) HttpServerTransportsName() string {
	return s.HttpServerName() + "Transports"
}

func (s *Service) GrpcClientTransportsName() string {
	return s.GrpcClientName() + "Transports"
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

func (s *Service) UnimplementedServerName() string {
	return "Unimplemented" + s.ProtoService.GoName + "Server"
}

func (s *Service) CQRSName() string {
	return s.ProtoService.GoName + "CqrsService"
}

func (s *Service) AssemblerName() string {
	return s.ProtoService.GoName + "Assembler"
}

func (s *Service) BusName() string {
	return s.ProtoService.GoName + "Bus"
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

func (s *Service) SetCommandPackage(file *protogen.File) error {
	pkg := proto.GetExtension(s.ProtoService.Desc.Options(), cqrs.E_Command).(*cqrs.Package)
	if pkg == nil {
		return nil
	}
	pkgAbs, pkgRel, err := s.resolvePkgPath(file.Desc.Path(), pkg.Relative)
	if err != nil {
		return fmt.Errorf("cqrs: %s, failed to resolve %s package path, %w", s.FullName(), "command", err)
	}
	s.Command = NewPackage(pkgAbs, pkgRel, pkg.Package)
	return nil
}

func (s *Service) SetQueryPackage(file *protogen.File) error {
	pkg := proto.GetExtension(s.ProtoService.Desc.Options(), cqrs.E_Query).(*cqrs.Package)
	if pkg == nil {
		return nil
	}
	pkgAbs, pkgRel, err := s.resolvePkgPath(file.Desc.Path(), pkg.Relative)
	if err != nil {
		return fmt.Errorf("cqrs: %s, failed to resolve %s package path, %w", s.FullName(), "query", err)
	}
	s.Query = NewPackage(pkgAbs, pkgRel, pkg.Package)
	return nil
}

func (s *Service) resolvePkgPath(filePath string, rel string) (string, string, error) {
	// 算出query或者command包的绝对路径
	pkgAbs, err := filepath.Abs(filepath.Join(filePath, rel))
	if err != nil {
		return "", "", err
	}
	//
	_, err = os.Stat(pkgAbs)
	if err != nil {
		return "", "", err
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", "", err
	}
	// 算出query或者command包的相对路径
	pkgRel, err := filepath.Rel(wd, pkgAbs)
	if err != nil {
		return "", "", err
	}
	pkgRel = filepath.Clean(pkgRel)
	return pkgAbs, pkgRel, nil
}

func NewServices(file *protogen.File) ([]*Service, error) {
	var services []*Service
	for _, pbService := range file.Services {
		service := &Service{
			ProtoService: pbService,
		}
		if err := service.SetQueryPackage(file); err != nil {
			return nil, err
		}
		if err := service.SetCommandPackage(file); err != nil {
			return nil, err
		}

		var endpoints []*Endpoint
		for _, pbMethod := range pbService.Methods {
			endpoint := &Endpoint{
				protoMethod: pbMethod,
			}
			if endpoint.IsStreaming() {
				return nil, fmt.Errorf("leo: unsupport stream method, %s", endpoint.FullName())
			}
			endpoint.SetHttpRule()
			endpoint.SetResponsibility()
			endpoints = append(endpoints, endpoint)
		}
		service.Endpoints = endpoints
		services = append(services, service)
	}
	return services, nil
}
