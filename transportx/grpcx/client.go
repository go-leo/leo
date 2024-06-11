package grpcx

import (
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type SDClient struct {
	client *grpctransport.Client
}

func NewSDClient(client *grpctransport.Client) *SDClient {
	return &SDClient{client: client}
}

func (c *SDClient) Endpoint() endpoint.Endpoint {

	return nil
}
