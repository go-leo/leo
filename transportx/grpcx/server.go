package grpcx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type ServerTransport struct {
	server *grpctransport.Server
}

func NewServerTransport(server *grpctransport.Server) *ServerTransport {
	return &ServerTransport{server: server}
}

func (t *ServerTransport) Endpoint() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		_, resp, err := t.server.ServeGRPC(ctx, request)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
