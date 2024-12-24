package internal

import (
	"fmt"
	"github.com/go-leo/leo/v3/cqrs"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"os"
	"path/filepath"
	"strings"
)

type Service struct {
	Service   *protogen.Service
	Endpoints []*Endpoint

	Command *Package
	Query   *Package
}

func (s Service) ServiceName() string {
	return s.Name() + "Service"
}

func (s Service) FullName() string {
	return string(s.Service.Desc.FullName())
}

func (s Service) Name() string {
	return s.Service.GoName
}

func (s Service) UnexportedName() string {
	name := s.Name()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) ServerEndpointsName() string {
	return s.Name() + "ServerEndpoints"
}

func (s Service) UnexportedServerEndpointsName() string {
	name := s.ServerEndpointsName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) ClientEndpointsName() string {
	return s.Name() + "ClientEndpoints"
}

func (s Service) UnexportedClientEndpointsName() string {
	name := s.ClientEndpointsName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) TransportsName() string {
	return s.Name() + "Transports"
}

func (s Service) UnexportedTransportsName() string {
	name := s.TransportsName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) ClientTransportsName() string {
	return s.Name() + "ClientTransports"
}

func (s Service) UnexportedClientTransportsName() string {
	name := s.ClientTransportsName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) FactoriesName() string {
	return s.Name() + "Factories"
}

func (s Service) UnexportedFactoriesName() string {
	name := s.FactoriesName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) EndpointersName() string {
	return s.Name() + "Endpointers"
}

func (s Service) UnexportedEndpointersName() string {
	name := s.EndpointersName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) BalancersName() string {
	return s.Name() + "Balancers"
}

func (s Service) UnexportedBalancersName() string {
	name := s.BalancersName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) ServerName() string {
	return s.Name() + "Server"
}

func (s Service) UnexportedServerName() string {
	name := s.ServerName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) ClientName() string {
	return s.Name() + "Client"
}

func (s Service) UnexportedClientName() string {
	name := s.ClientName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) GrpcServerName() string {
	return s.Name() + "GrpcServer"
}

func (s Service) UnexportedGrpcServerName() string {
	name := s.GrpcServerName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) GrpcClientName() string {
	return s.Name() + "GrpcClient"
}

func (s Service) UnexportedGrpcClientName() string {
	name := s.GrpcClientName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) GrpcServerTransportsName() string {
	return s.GrpcServerName() + "Transports"
}

func (s Service) UnexportedGrpcServerTransportsName() string {
	name := s.GrpcServerTransportsName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) GrpcClientEndpointsName() string {
	return s.GrpcClientName() + "Endpoints"
}

func (s Service) UnexportedGrpcClientEndpointsName() string {
	name := s.GrpcClientEndpointsName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) GrpcFactoriesName() string {
	return s.GrpcClientName() + "Factories"
}

func (s Service) UnexportedGrpcFactoriesName() string {
	name := s.GrpcFactoriesName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) HttpServerName() string {
	return s.Name() + "HttpServer"
}

func (s Service) UnexportedHttpServerName() string {
	name := s.HttpServerName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) ClientServiceName() string {
	return s.Name() + "ClientService"
}

func (s Service) UnexportedClientServiceName() string {
	name := s.ClientServiceName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) HttpClientName() string {
	return s.Name() + "HttpClient"
}

func (s Service) UnexportedHttpClientName() string {
	name := s.HttpClientName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) HttpServerTransportsName() string {
	return s.HttpServerName() + "Transports"
}

func (s Service) UnexportedHttpServerTransportsName() string {
	name := s.HttpServerTransportsName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) GrpcClientTransportsName() string {
	return s.GrpcClientName() + "Transports"
}

func (s Service) UnexportedGrpcClientTransportsName() string {
	name := s.GrpcClientTransportsName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) HttpClientTransportsName() string {
	return s.HttpClientName() + "Transports"
}

func (s Service) UnexportedHttpClientTransportsName() string {
	name := s.HttpClientTransportsName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) HttpRoutesName() string {
	return s.Name() + "HttpRoutes"
}

func (s Service) UnimplementedServerName() string {
	return "Unimplemented" + s.Service.GoName + "Server"
}

func (s Service) CQRSName() string {
	return s.Service.GoName + "CqrsService"
}

func (s Service) UnexportedCQRSName() string {
	name := s.CQRSName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) AssemblerName() string {
	return s.Service.GoName + "Assembler"
}

func (s Service) BusName() string {
	return s.Service.GoName + "Bus"
}

func NewServices(file *protogen.File) ([]*Service, error) {
	var services []*Service
	for _, service := range file.Services {
		var endpoints []*Endpoint
		for _, method := range service.Methods {
			fmName := fmt.Sprintf("/%s/%s", service.Desc.FullName(), method.Desc.Name())
			if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
				return nil, fmt.Errorf("leo: %s, unsupport stream method", fmName)
			}
			endpoints = append(endpoints, &Endpoint{method: method})
		}
		services = append(services, &Service{Service: service, Endpoints: endpoints})
	}
	return services, nil
}

func NewHttpServices(file *protogen.File) ([]*Service, error) {
	services, err := NewServices(file)
	if err != nil {
		return nil, err
	}
	var httpServices []*Service
	for _, service := range services {
		var endpoints []*Endpoint
		for _, endpoint := range service.Endpoints {
			method := endpoint.method
			extHTTP := proto.GetExtension(method.Desc.Options(), annotations.E_Http)
			if extHTTP == nil || extHTTP == annotations.E_Http.InterfaceOf(annotations.E_Http.Zero()) {
				extHTTP = &annotations.HttpRule{
					Pattern: &annotations.HttpRule_Post{Post: endpoint.FullName()},
					Body:    "*",
				}
			}
			endpoint.httpRule = &HttpRule{rule: extHTTP.(*annotations.HttpRule)}
			endpoints = append(endpoints, endpoint)
		}
		httpServices = append(httpServices, &Service{Service: service.Service, Endpoints: endpoints})
	}
	return httpServices, nil
}

func NewCQRSServices(file *protogen.File) ([]*Service, error) {
	services, err := NewServices(file)
	if err != nil {
		return nil, err
	}

	var cqrsServices []*Service
	for _, service := range services {
		commandPkg := proto.GetExtension(service.Service.Desc.Options(), cqrs.E_Command).(*cqrs.Package)
		if commandPkg == nil {
			continue
		}
		queryPkg := proto.GetExtension(service.Service.Desc.Options(), cqrs.E_Query).(*cqrs.Package)
		if queryPkg == nil {
			continue
		}

		commandPkgAbs, commandPkgRel, err := resolvePkgPath(file.Desc.Path(), commandPkg.Relative)
		if err != nil {
			return nil, fmt.Errorf("cqrs: %s, failed to resolve %s package path, %w", service.FullName(), "command", err)
		}
		queryPkgAbs, queryPkgRel, err := resolvePkgPath(file.Desc.Path(), queryPkg.Relative)
		if err != nil {
			return nil, fmt.Errorf("cqrs: %s, failed to resolve %s package path, %w", service.FullName(), "query", err)
		}

		var endpoints []*Endpoint
		for _, endpoint := range service.Endpoints {
			method := endpoint.method
			responsibility := proto.GetExtension(method.Desc.Options(), cqrs.E_Responsibility).(cqrs.Responsibility)
			switch responsibility {
			case cqrs.Responsibility_Unknown:
				return nil, fmt.Errorf("cqrs: %s, cqrs unknown", endpoint.FullName())
			case cqrs.Responsibility_Command:
				endpoints = append(endpoints, &Endpoint{method: method, responsibility: responsibility})
				continue
			case cqrs.Responsibility_Query:
				endpoints = append(endpoints, &Endpoint{method: method, responsibility: responsibility})
				continue
			default:
				return nil, fmt.Errorf("cqrs: %s, %s responsibility unsupported", endpoint.FullName(), responsibility)
			}
		}
		cqrsServices = append(cqrsServices, &Service{
			Service:   service.Service,
			Command:   NewPackage(commandPkgAbs, commandPkgRel, commandPkg.Package),
			Query:     NewPackage(queryPkgAbs, queryPkgRel, queryPkg.Package),
			Endpoints: endpoints,
		})
	}
	return cqrsServices, nil
}

func resolvePkgPath(filePath string, rel string) (string, string, error) {
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
