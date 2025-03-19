package internal

import (
	"fmt"
	"github.com/go-leo/gox/slicex"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

type Service struct {
	file         *protogen.File
	protoService *protogen.Service
	Endpoints    []*Endpoint
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

func (s *Service) BusName() string {
	return s.protoService.GoName + "Bus"
}

func (s *Service) CommandEndpoints() []*Endpoint {
	return slicex.Filter(s.Endpoints, func(_ int, endpoint *Endpoint) bool {
		return endpoint.IsCommand()
	})
}

func (s *Service) QueryEndpoints() []*Endpoint {
	return slicex.Filter(s.Endpoints, func(_ int, endpoint *Endpoint) bool {
		return endpoint.IsQuery()
	})
}

func NewServices(file *protogen.File) ([]*Service, error) {
	var services []*Service
	for _, pbService := range file.Services {
		service := &Service{
			file:         file,
			protoService: pbService,
		}
		var endpoints []*Endpoint
		for _, pbMethod := range pbService.Methods {
			endpoint := &Endpoint{
				service:     service,
				protoMethod: pbMethod,
			}
			if endpoint.IsStreaming() {
				return nil, fmt.Errorf("leo: unsupport stream httpMethod, %s", endpoint.FullName())
			}
			if err := endpoint.ParseHttpRule(); err != nil {
				return nil, err
			}
			endpoints = append(endpoints, endpoint)
		}
		// skip empty endpoints
		if len(endpoints) == 0 {
			continue
		}
		service.Endpoints = endpoints
		services = append(services, service)
	}
	return services, nil
}
