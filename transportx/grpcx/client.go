package grpcx

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-leo/leo/v3/sdx"
	"google.golang.org/grpc"
	"io"
)

func ClientFactory(
	serviceName string,
	method string,
	enc grpctransport.EncodeRequestFunc,
	dec grpctransport.DecodeResponseFunc,
	grpcReply interface{},
	options ...grpctransport.ClientOption,
) sdx.Factory {
	return func(ctx context.Context, args any) sd.Factory {
		if args == nil {
			args = make([]grpc.DialOption, 0)
		}
		return func(target string) (endpoint.Endpoint, io.Closer, error) {
			dialOptions, ok := args.([]grpc.DialOption)
			if !ok {
				return nil, nil, errors.New("invalid grpc factory args")
			}
			conn, err := grpc.NewClient(target, dialOptions...)
			if err != nil {
				return nil, nil, err
			}
			client := grpctransport.NewClient(conn, serviceName, method, enc, dec, grpcReply, options...)
			return sdx.WithTarget(target, client.Endpoint()), conn, nil
		}
	}

}
