package metadatax

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func GrpcOutgoingInjector(ctx context.Context, grpcMD *metadata.MD) context.Context {
	md, ok := FromOutgoingContext(ctx)
	if !ok {
		return ctx
	}
	for _, key := range md.Keys() {
		grpcMD.Set(key, md.Values(key)...)
	}
	return ctx
}

func GrpcIncomingInjector(ctx context.Context, md metadata.MD) context.Context {
	return NewIncomingContext(ctx, FromGrpcMetadata(md))
}

// AsGrpcMetadata Convert Metadata to metadata.MD
func AsGrpcMetadata(md Metadata) metadata.MD {
	res := metadata.MD{}
	for _, key := range md.Keys() {
		res.Set(key, md.Values(key)...)
	}
	return res
}

// FromGrpcMetadata Convert metadata.MD to Metadata
//
// the key is converted to lowercase.
func FromGrpcMetadata(md metadata.MD) Metadata {
	res := New()
	for key, values := range md {
		res.Set(key, values...)
	}
	return res
}
