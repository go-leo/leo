package grpcsonybreaker

import (
	"context"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
)

func UnaryClientInterceptor(Name string, opts ...Option) grpc.UnaryClientInterceptor {
	st := defaultSettings(Name)
	apply(st, opts...)
	cb := gobreaker.NewCircuitBreaker(*st)
	return unaryClientInterceptor(cb)
}

func unaryClientInterceptor(cb *gobreaker.CircuitBreaker) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) (err error) {
		_, err = cb.Execute(func() (interface{}, error) {
			err = invoker(ctx, method, req, reply, cc, opts...)
			if err != nil {
				return nil, err
			}
			return nil, nil
		})
		return err
	}
}
