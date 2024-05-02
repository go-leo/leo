package internal

import (
	"fmt"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"strings"
)

type Service struct {
	Service   *protogen.Service
	Endpoints []*Endpoint
}

func (s Service) Name() string {
	return s.Service.GoName
}

func (s Service) UnexportedName() string {
	name := s.Service.GoName
	return strings.ToLower(name[:1]) + name[1:]
}

func (s Service) ServerName() interface{} {
	return s.Name() + "Server"
}

func (s Service) GRPCServerName() interface{} {
	return s.Name() + "GRPCServer"
}

func (s Service) GRPCClientName() interface{} {
	return s.Name() + "GRPCClient"
}

func (s Service) HTTPServerName() interface{} {
	return s.Name() + "HTTPServer"
}

func (s Service) HTTPClientName() interface{} {
	return s.Name() + "HTTPClient"
}

func (s Service) UnexportedGRPCServerName() interface{} {
	return "gRPC" + s.Name() + "Server"
}

func (s Service) UnexportedGRPCClientName() interface{} {
	return "gRPC" + s.Name() + "Client"
}

func (s Service) UnexportedHTTPServerName() interface{} {
	return "http" + s.Name() + "Server"
}

func (s Service) UnexportedHTTPClientName() interface{} {
	return "http" + s.Name() + "Client"
}

func (s Service) UnexportedGRPCEndpointsName() interface{} {
	return "gRPC" + s.Name() + "Endpoints"
}

func (s Service) UnimplementedServerName() string {
	return "Unimplemented" + s.Service.GoName + "Server"
}

func (s Service) FullName() string {
	return string(s.Service.Desc.FullName())
}

func (s Service) EndpointsName() interface{} {
	return s.UnexportedName() + "Endpoints"
}

func (s Service) UnexportedEndpointsName() interface{} {
	return s.UnexportedName() + "Endpoints"
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
			extHTTP := proto.GetExtension(method.Desc.Options(), annotations.E_Http)
			if extHTTP == nil {
				return nil, fmt.Errorf("missing http rule: %s", fmName)
			}
			if extHTTP == annotations.E_Http.InterfaceOf(annotations.E_Http.Zero()) {
				return nil, fmt.Errorf("missing http rule: %s", fmName)
			}
			rule := extHTTP.(*annotations.HttpRule)
			if len(rule.AdditionalBindings) > 0 {
				return nil, fmt.Errorf("unsupport additional bindings: %s", fmName)
			}
			endpoints = append(endpoints, NewEndpoint(method, rule))
		}
		services = append(services, &Service{Service: service, Endpoints: endpoints})
	}
	return services, nil
}
