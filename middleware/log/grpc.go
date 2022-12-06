package log

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"github.com/go-leo/leo/v2/log"
)

func GRPCClientMiddleware(loggerFactory func(ctx context.Context) log.Logger, opts ...Option) grpc.UnaryClientInterceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	skipMap := make(map[string]struct{}, len(o.Skips))
	for _, skip := range o.Skips {
		skipMap[skip] = struct{}{}
	}
	return func(
		ctx context.Context,
		method string,
		req any,
		reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		if _, ok := skipMap[method]; ok {
			return invoker(ctx, method, req, reply, cc, opts...)
		}
		if loggerFactory == nil {
			return invoker(ctx, method, req, reply, cc, opts...)
		}
		logger := loggerFactory(ctx)
		if logger == nil {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		startTime := time.Now()

		err := invoker(ctx, method, req, reply, cc, opts...)

		builder := NewFieldBuilder().
			System("grpc.client").
			StartTime(startTime).
			Deadline(ctx).
			Method(method).
			Status(status.Code(err).String()).
			Latency(time.Since(startTime))
		peerInfo, ok := peer.FromContext(ctx)
		if ok {
			builder.PeerAddress(peerInfo.Addr.String())
		}

		if err != nil {
			builder.Error(err.Error())
		}

		switch {
		case o.Payload:
			builder.Request(req).Response(reply)
		case o.PayloadWhenError || err != nil:
			builder.Request(req).Response(reply)
		}

		if err != nil {
			logger.ErrorF(builder.Build()...)
		} else {
			logger.InfoF(builder.Build()...)
		}
		return err
	}
}

func GRPCServerMiddleware(loggerFactory func(ctx context.Context) log.Logger, opts ...Option) grpc.UnaryServerInterceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	skipMap := make(map[string]struct{}, len(o.Skips))
	for _, skip := range o.Skips {
		skipMap[skip] = struct{}{}
	}
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if _, ok := skipMap[info.FullMethod]; ok {
			return handler(ctx, req)
		}
		if loggerFactory == nil {
			return handler(ctx, req)
		}
		logger := loggerFactory(ctx)
		if logger == nil {
			return handler(ctx, req)
		}

		startTime := time.Now()

		resp, err := handler(ctx, req)

		builder := NewFieldBuilder().
			System("grpc.server").
			StartTime(startTime).
			Deadline(ctx).
			Method(info.FullMethod).
			Status(status.Code(err).String()).
			Latency(time.Since(startTime))
		peerInfo, ok := peer.FromContext(ctx)
		if ok {
			builder.PeerAddress(peerInfo.Addr.String())
		}
		if err != nil {
			builder.Error(err.Error())
		}

		switch {
		case o.Payload:
			builder.Request(req).Response(resp)
		case o.PayloadWhenError || err != nil:
			builder.Request(req).Response(resp)
		}

		if err != nil {
			logger.ErrorF(builder.Build()...)
		} else {
			logger.InfoF(builder.Build()...)
		}
		return resp, err
	}
}
