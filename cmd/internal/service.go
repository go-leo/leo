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

func (s Service) Name() string {
	return s.Service.GoName
}

func (s Service) UnexportedName() string {
	name := s.Name()
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

func (s Service) GRPCServerName() string {
	return s.Name() + "GRPCServer"
}

func (s Service) UnexportedGRPCServerName() string {
	name := s.GRPCServerName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) GRPCClientName() string {
	return s.Name() + "GRPCClient"
}

func (s Service) UnexportedGRPCClientName() string {
	name := s.GRPCClientName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) HTTPServerName() string {
	return s.Name() + "HTTPServer"
}

func (s Service) UnexportedHTTPServerName() string {
	name := s.HTTPServerName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) HTTPClientName() string {
	return s.Name() + "HTTPClient"
}

func (s Service) UnexportedHTTPClientName() string {
	name := s.HTTPClientName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) UnimplementedServerName() string {
	return "Unimplemented" + s.Service.GoName + "Server"
}

func (s Service) FullName() string {
	return string(s.Service.Desc.FullName())
}

func (s Service) EndpointsName() string {
	return s.Name() + "Endpoints"
}

func (s Service) UnexportedEndpointsName() string {
	name := s.EndpointsName()
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) CQRSName() string {
	return s.Service.GoName + "CQRSService"
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
				return nil, fmt.Errorf("unsupport stream method: %s", fmName)
			}
			endpoints = append(endpoints, &Endpoint{method: method})
		}
		services = append(services, &Service{Service: service, Endpoints: endpoints})
	}
	return services, nil
}

func NewHttpServices(file *protogen.File) ([]*Service, error) {
	var services []*Service
	for _, service := range file.Services {
		var endpoints []*Endpoint
		for _, method := range service.Methods {
			fmName := fmt.Sprintf("/%s/%s", service.Desc.FullName(), method.Desc.Name())
			if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
				continue
				//return nil, fmt.Errorf("unsupport stream method: %s", fmName)
			}
			extHTTP := proto.GetExtension(method.Desc.Options(), annotations.E_Http)
			if extHTTP == nil {
				continue
				//return nil, fmt.Errorf("missing http rule: %s", fmName)
			}
			if extHTTP == annotations.E_Http.InterfaceOf(annotations.E_Http.Zero()) {
				continue
				//return nil, fmt.Errorf("missing http rule: %s", fmName)
			}
			rule := extHTTP.(*annotations.HttpRule)
			if len(rule.AdditionalBindings) > 0 {
				return nil, fmt.Errorf("unsupport additional bindings: %s", fmName)
			}
			endpoints = append(endpoints, &Endpoint{method: method, httpRule: &HttpRule{rule: rule}})
		}
		services = append(services, &Service{Service: service, Endpoints: endpoints})
	}
	return services, nil
}

func NewCQRSServices(file *protogen.File) ([]*Service, error) {
	var services []*Service
	for _, service := range file.Services {
		serviceFullName := service.Desc.FullName()

		commandPkg := proto.GetExtension(service.Desc.Options(), cqrs.E_Command).(*cqrs.Package)
		if commandPkg == nil {
			continue
		}
		queryPkg := proto.GetExtension(service.Desc.Options(), cqrs.E_Query).(*cqrs.Package)
		if queryPkg == nil {
			continue
		}

		commandPkgAbs, commandPkgRel, err := resolvePkgPath(file.Desc.Path(), commandPkg.Relative)
		if err != nil {
			return nil, fmt.Errorf("cqrs: %s, failed to resolve %s package path, %w", serviceFullName, "command", err)
		}
		queryPkgAbs, queryPkgRel, err := resolvePkgPath(file.Desc.Path(), queryPkg.Relative)
		if err != nil {
			return nil, fmt.Errorf("cqrs: %s, failed to resolve %s package path, %w", serviceFullName, "query", err)
		}

		var endpoints []*Endpoint
		for _, method := range service.Methods {
			if method.Desc.IsStreamingClient() {
				continue
			}
			endpoint := method.GoName
			methodFullName := fmt.Sprintf("/%s/%s", serviceFullName, endpoint)
			responsibility := proto.GetExtension(method.Desc.Options(), cqrs.E_Responsibility).(cqrs.Responsibility)
			switch responsibility {
			case cqrs.Responsibility_Unknown:
				return nil, fmt.Errorf("cqrs: %s, cqrs unknown", methodFullName)
			case cqrs.Responsibility_Command:
				endpoints = append(endpoints, &Endpoint{method: method, responsibility: responsibility})
				continue
			case cqrs.Responsibility_Query:
				endpoints = append(endpoints, &Endpoint{method: method, responsibility: responsibility})
				continue
			default:
				return nil, fmt.Errorf("cqrs: %s, %s responsibility unsupported", methodFullName, responsibility)
			}
		}
		services = append(services, &Service{
			Service:   service,
			Command:   NewPackage(commandPkgAbs, commandPkgRel, commandPkg.Package),
			Query:     NewPackage(queryPkgAbs, queryPkgRel, queryPkg.Package),
			Endpoints: endpoints,
		})
	}
	return services, nil
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
