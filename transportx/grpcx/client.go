package grpcx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type Client struct {
	client *grpctransport.Client
}

func NewClient(
	cc *grpc.ClientConn,
	serviceName string,
	method string,
	enc grpctransport.EncodeRequestFunc,
	dec grpctransport.DecodeResponseFunc,
	grpcReply interface{},
	options ...grpctransport.ClientOption,
) *Client {
	return &Client{client: grpctransport.NewClient(cc, serviceName, method, enc, dec, grpcReply, options...)}
}

func (c *Client) Endpoint(ctx context.Context) endpoint.Endpoint {
	return c.client.Endpoint()
}
