package requestid

import (
	"context"
	"regexp"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/stringx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := &options{}
	o.apply(opts...)
	o.init()
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		var requestID string
		// 1. from incoming
		requestID, _ = fromIncoming(ctx, o)
		if stringx.IsNotBlank(requestID) {
			return serverHandler(ctx, req, requestID, o, handler)
		}

		// 2. from TraceContext, TraceContext is a propagator that supports the W3C Trace Context format
		// (https://www.w3.org/TR/trace-context/)
		requestID, _ = fromTrace(ctx)
		if stringx.IsNotBlank(requestID) {
			return serverHandler(ctx, req, requestID, o, handler)
		}

		// 3. generate
		requestID = o.generator()
		return serverHandler(ctx, req, requestID, o, handler)
	}
}

func serverHandler(ctx context.Context, req interface{}, requestID string, o *options, handler grpc.UnaryHandler) (interface{}, error) {
	ctx = NewContext(ctx, requestID)
	_ = grpc.SetTrailer(ctx, metadata.Pairs("X-Request-ID", requestID))
	o.handler(ctx, requestID)
	return handler(ctx, req)
}

func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := &options{}
	o.apply(opts...)
	o.init()
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx := ss.Context()
		var requestID string

		// 1. from incoming
		requestID, _ = fromIncoming(ctx, o)
		if stringx.IsNotBlank(requestID) {
			return serverStreamer(srv, ctx, requestID, o, handler)
		}

		// 2. from TraceContext, TraceContext is a propagator that supports the W3C Trace Context format
		// (https://www.w3.org/TR/trace-context/)
		requestID, _ = fromTrace(ctx)
		if stringx.IsNotBlank(requestID) {
			return serverStreamer(srv, ctx, requestID, o, handler)
		}

		// 3. generate
		requestID = o.generator()
		return serverStreamer(srv, ctx, requestID, o, handler)
	}
}

func serverStreamer(srv interface{}, ctx context.Context, requestID string, o *options, handler grpc.StreamHandler) error {
	ctx = NewContext(ctx, requestID)
	_ = grpc.SetTrailer(ctx, metadata.Pairs("X-Request-ID", requestID))
	o.handler(ctx, requestID)
	return handler(srv, &serverStream{ctx: ctx})
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *serverStream) Context() context.Context {
	return w.ctx
}

func fromIncoming(ctx context.Context, o *options) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		vals := md.Get(o.headerKey)
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return "", false
}

var traceCtxRegExp = regexp.MustCompile("^(?P<version>[0-9a-f]{2})-(?P<traceID>[a-f0-9]{32})-(?P<spanID>[a-f0-9]{16})-(?P<traceFlags>[a-f0-9]{2})(?:-.*)?$")

const traceparentHeader = "traceparent"

func fromTrace(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false

	}
	vals := md.Get(traceparentHeader)
	if len(vals) <= 0 {
		return "", false
	}

	h := vals[0]
	if h == "" {
		return "", false
	}
	matches := traceCtxRegExp.FindStringSubmatch(h)
	if len(matches) == 0 {
		return "", false
	}
	if len(matches[2]) != 32 {
		return "", false
	}
	return matches[2][:32], true
}
