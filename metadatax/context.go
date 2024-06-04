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
	if !ok {
		return nil, false
	}
	return Join(md), true
}

type outgoingKey struct{}

// NewOutgoingContext creates a new context with outgoing Metadata attached.
func NewOutgoingContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, outgoingKey{}, rawMD{md: md})
}

// FromOutgoingContext returns the outgoing metadata in ctx if it exists.
func FromOutgoingContext(ctx context.Context) (Metadata, bool) {
	md, ok := ctx.Value(outgoingKey{}).(rawMD)
	if !ok {
		return nil, false
	}
	out := Join(append([]Metadata{md.md}, md.added...)...)
	return out, ok
}

// AppendToOutgoingContext appends the Metadata to the outgoing context.
func AppendToOutgoingContext(ctx context.Context, mds ...Metadata) context.Context {
	oldMD, _ := ctx.Value(outgoingKey{}).(rawMD)
	added := append(append(make([]Metadata, 0, len(oldMD.added)+len(mds)), oldMD.added...), mds...)
	return context.WithValue(ctx, outgoingKey{}, rawMD{md: oldMD.md, added: added})
}

type rawMD struct {
	md    Metadata
	added []Metadata
}
