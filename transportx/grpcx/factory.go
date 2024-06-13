package grpcx

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"io"
)

func Factory(
	serviceName string,
	method string,
	enc grpctransport.EncodeRequestFunc,
	dec grpctransport.DecodeResponseFunc,
	grpcReply interface{},
	dialOption []grpc.DialOption,
	options ...grpctransport.ClientOption,
) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.NewClient(instance, dialOption...)
		if err != nil {
			return nil, nil, err
		}
		client := grpctransport.NewClient(conn, serviceName, method, enc, dec, grpcReply, options...)
		return client.Endpoint(), conn, nil
	}
}
