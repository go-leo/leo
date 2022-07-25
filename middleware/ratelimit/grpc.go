package ratelimit

import (
	"context"

	"google.golang.org/grpc"
)

func AllowerGRPCServerMiddleware(limit Allower) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if !limit.Allow() {
			return nil, ErrLimited
		}
		return handler(ctx, req)
	}
}

func WaiterGRPCServerMiddleware(limit Waiter) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if err := limit.Wait(ctx); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
