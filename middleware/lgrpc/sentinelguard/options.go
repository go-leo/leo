package sentinelguard

import (
	"context"

	"github.com/alibaba/sentinel-golang/core/base"
	"google.golang.org/grpc"
)

type (
	Option func(*options)

	options struct {
		unaryClientResourceExtract func(context.Context, string, interface{}, *grpc.ClientConn) string
		unaryServerResourceExtract func(context.Context, interface{}, *grpc.UnaryServerInfo) string

		streamClientResourceExtract func(context.Context, *grpc.StreamDesc, *grpc.ClientConn, string) string
		streamServerResourceExtract func(interface{}, grpc.ServerStream, *grpc.StreamServerInfo) string

		unaryClientBlockFallback func(context.Context, string, interface{}, *grpc.ClientConn, *base.BlockError) error
		unaryServerBlockFallback func(context.Context, interface{}, *grpc.UnaryServerInfo, *base.BlockError) (interface{}, error)

		streamClientBlockFallback func(context.Context, *grpc.StreamDesc, *grpc.ClientConn, string, *base.BlockError) (grpc.ClientStream, error)
		streamServerBlockFallback func(interface{}, grpc.ServerStream, *grpc.StreamServerInfo, *base.BlockError) error
	}
)

// UnaryClientResourceExtractor sets the resource extractor of unary client request.
// The second string parameter is the full method name of current invocation.
func UnaryClientResourceExtractor(fn func(context.Context, string, interface{}, *grpc.ClientConn) string) Option {
	return func(opts *options) {
		opts.unaryClientResourceExtract = fn
	}
}

// UnaryServerResourceExtractor sets the resource extractor of unary server request.
func UnaryServerResourceExtractor(fn func(context.Context, interface{}, *grpc.UnaryServerInfo) string) Option {
	return func(opts *options) {
		opts.unaryServerResourceExtract = fn
	}
}

// StreamClientResourceExtractor sets the resource extractor of stream client request.
func StreamClientResourceExtractor(fn func(context.Context, *grpc.StreamDesc, *grpc.ClientConn, string) string) Option {
	return func(opts *options) {
		opts.streamClientResourceExtract = fn
	}
}

// StreamServerResourceExtractor sets the resource extractor of stream server request.
func StreamServerResourceExtractor(fn func(interface{}, grpc.ServerStream, *grpc.StreamServerInfo) string) Option {
	return func(opts *options) {
		opts.streamServerResourceExtract = fn
	}
}

// UnaryClientBlockFallback sets the block fallback handler of unary client request.
// The second string parameter is the full method name of current invocation.
func UnaryClientBlockFallback(fn func(context.Context, string, interface{}, *grpc.ClientConn, *base.BlockError) error) Option {
	return func(opts *options) {
		opts.unaryClientBlockFallback = fn
	}
}

// UnaryServerBlockFallback sets the block fallback handler of unary server request.
func UnaryServerBlockFallback(fn func(context.Context, interface{}, *grpc.UnaryServerInfo, *base.BlockError) (interface{}, error)) Option {
	return func(opts *options) {
		opts.unaryServerBlockFallback = fn
	}
}

// StreamClientBlockFallback sets the block fallback handler of stream client request.
func StreamClientBlockFallback(fn func(context.Context, *grpc.StreamDesc, *grpc.ClientConn, string, *base.BlockError) (grpc.ClientStream, error)) Option {
	return func(opts *options) {
		opts.streamClientBlockFallback = fn
	}
}

// StreamServerBlockFallback sets the block fallback handler of stream server request.
func StreamServerBlockFallback(fn func(interface{}, grpc.ServerStream, *grpc.StreamServerInfo, *base.BlockError) error) Option {
	return func(opts *options) {
		opts.streamServerBlockFallback = fn
	}
}

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}
