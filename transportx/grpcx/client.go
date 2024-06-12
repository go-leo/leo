package grpcx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-leo/leo/v3/transportx"
	"google.golang.org/grpc"
)

type Client struct {
	client      map[string]*grpctransport.Client
	target      string
	dialOption  []grpc.DialOption
	serviceName string
	method      string
	enc         grpctransport.EncodeRequestFunc
	dec         grpctransport.DecodeResponseFunc
	grpcReply   interface{}
	options     []grpctransport.ClientOption
}

func NewClient(
	target string,
	dialOptions []grpc.DialOption,
	serviceName string,
	method string,
	enc grpctransport.EncodeRequestFunc,
	dec grpctransport.DecodeResponseFunc,
	grpcReply interface{},
	options ...grpctransport.ClientOption,
) *Client {
	return &Client{
		target:      target,
		dialOption:  dialOptions,
		serviceName: serviceName,
		method:      method,
		enc:         enc,
		dec:         dec,
		grpcReply:   grpcReply,
		options:     options,
	}
}

func (c *Client) Endpoint(ctx context.Context) endpoint.Endpoint {
	colors, ok := transportx.ExtractColors(ctx)
	if !ok {

	}

	consul.NewInstancer(ctx, consul.NewClient(ctx), consul.Instances)
	sd.NewEndpointer()

	cc := grpc.NewClient(ctx)

	grpctransport.NewClient(cc, c.serviceName, c.method, c.enc, c.dec, c.grpcReply, c.options...)

	return c.client.Endpoint()
}
