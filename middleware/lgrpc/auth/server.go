package auth

import (
	"context"
	"strings"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/slicex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Authorizer interface {
	Authorize(ctx context.Context, fullMethodName string) (context.Context, error)
}

type AuthorizerFunc func(ctx context.Context, fullMethodName string) (context.Context, error)

func (f AuthorizerFunc) Authorize(ctx context.Context, fullMethodName string) (context.Context, error) {
	return f(ctx, fullMethodName)
}

func DefaultAuthorizer(valid func(ctx context.Context, scheme, credentials string) bool) Authorizer {
	return AuthorizerFunc(func(ctx context.Context, fullMethodName string) (context.Context, error) {
		authorizations := metadata.ValueFromIncomingContext(ctx, "authorization")
		if slicex.IsEmpty(authorizations) {
			return ctx, status.Errorf(codes.InvalidArgument, "missing metadata")
		}
		auth := strings.TrimSpace(authorizations[0])
		authSeg := strings.SplitN(auth, " ", 2)
		if len(authSeg) != 2 {
			return ctx, status.Errorf(codes.InvalidArgument, "authorization is invalid")
		}
		if !valid(ctx, authSeg[0], authSeg[1]) {
			return ctx, status.Errorf(codes.Unauthenticated, "invalid token")
		}
		return ctx, nil
	})
}

func UnaryServerInterceptor(authorizer AuthorizerFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if authorizerSrv, ok := info.Server.(Authorizer); ok {
			newCtx, err := authorizerSrv.Authorize(ctx, info.FullMethod)
			if err != nil {
				return nil, err
			}
			return handler(newCtx, req)
		}
		if authorizer == nil {
			return handler(ctx, req)
		}
		newCtx, err := authorizer.Authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}

func StreamServerInterceptor(authorizer AuthorizerFunc) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if authorizerSrv, ok := srv.(Authorizer); ok {
			newCtx, err := authorizerSrv.Authorize(ss.Context(), info.FullMethod)
			if err != nil {
				return err
			}
			return handler(srv, &serverStream{ctx: newCtx})
		}
		if authorizer == nil {
			return handler(srv, ss)
		}
		newCtx, err := authorizer.Authorize(ss.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		return handler(srv, &serverStream{ctx: newCtx})
	}
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *serverStream) Context() context.Context {
	return w.ctx
}
