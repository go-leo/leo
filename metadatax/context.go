package metadatax

import (
	"context"
)

type incomingKey struct{}

// NewIncomingContext creates a new context with incoming md attached.
func NewIncomingContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, incomingKey{}, md)
}

// FromIncomingContext returns the incoming metadata in ctx if it exists.
func FromIncomingContext(ctx context.Context) (Metadata, bool) {
	md, ok := ctx.Value(incomingKey{}).(Metadata)
	return md, ok
}

type outgoingKey struct{}

// NewOutgoingContext creates a new context with outgoing md attached.
func NewOutgoingContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, outgoingKey{}, md)
}

// FromOutgoingContext returns the outgoing metadata in ctx if it exists.
func FromOutgoingContext(ctx context.Context) (Metadata, bool) {
	md, ok := ctx.Value(outgoingKey{}).(Metadata)
	return md, ok
}
