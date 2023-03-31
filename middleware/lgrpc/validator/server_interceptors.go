package validator

import (
	"context"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := validate(req); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapper := &serverStreamWrapper{stream}
		return handler(srv, wrapper)
	}
}

type serverStreamWrapper struct {
	grpc.ServerStream
}

func (s *serverStreamWrapper) RecvMsg(m interface{}) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	return validate(m)
}
