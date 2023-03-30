package ratelimiter

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Limiter interface {
	Limit(ctx context.Context) bool
}

var alwaysFalseLimiter = alwaysLimiter{v: false}

type alwaysLimiter struct {
	v bool
}

func (l alwaysLimiter) Limit(ctx context.Context) bool {
	return l.v
}

func UnaryServerInterceptor(limiter Limiter) grpc.UnaryServerInterceptor {
	if limiter == nil {
		limiter = alwaysFalseLimiter
	}
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if limiter.Limit(ctx) {
			return nil, status.Errorf(codes.ResourceExhausted, "rate limit exceeded, method: %s", info.FullMethod)
		}
		return handler(ctx, req)
	}
}

func StreamServerInterceptor(limiter Limiter) grpc.StreamServerInterceptor {
	if limiter == nil {
		limiter = alwaysFalseLimiter
	}
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if limiter.Limit(stream.Context()) {
			return status.Errorf(codes.ResourceExhausted, "rate limit exceeded, method: %s", info.FullMethod)
		}
		return handler(srv, stream)
	}
}
