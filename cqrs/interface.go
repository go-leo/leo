package cqrs

import (
	"context"
)

// Metadata is a mapping from metadata keys to values.
type Metadata interface {
	// Add adds the key, value pair.
	Add(key, value string)

	// Set sets the value of a given key with a slice of values.
	Set(key string, value ...string)

	// Append adds the values to key, not overwriting what was already stored at that key.
	Append(key string, value ...string)

	// Get gets the first value associated with the given key.
	Get(key string) string

	// Values returns all values associated with the given key.
	Values(key string) []string

	// Keys returns the keys of the Metadata.
	Keys() []string

	// Delete removes the values for a given key.
	Delete(key string)

	// Len returns the number of items in Metadata.
	Len() int

	// Clone returns a copy of Metadata or nil if Metadata is nil.
	Clone() Metadata
}

// ========================== Command ==========================

// CommandHandler is a command handler that to update data.
type CommandHandler[Args any] interface {
	Handle(ctx context.Context, args Args) (Metadata, error)
}

// The CommandHandlerFunc type is an adapter to allow the use of ordinary functions as CommandHandler.
type CommandHandlerFunc[Args any] func(ctx context.Context, args Args) (Metadata, error)

// Handle calls f(ctx).
func (f CommandHandlerFunc[Args]) Handle(ctx context.Context, args Args) (Metadata, error) {
	return f(ctx, args)
}

// NoopCommand is an CommandHandler that does nothing and returns a nil error.
type NoopCommand[Args any] struct{}

func (NoopCommand[Args]) Handle(context.Context, Args) (Metadata, error) { return nil, nil }

// ========================== Query ==========================

// QueryHandler is a query handler that to handlers to read data.
type QueryHandler[Args any, Result any] interface {
	Handle(ctx context.Context, args Args) (Result, error)
}

// The QueryHandlerFunc type is an adapter to allow the use of ordinary functions as QueryHandler.
type QueryHandlerFunc[Args any, Result any] func(ctx context.Context, args Args) (Result, error)

// Handle calls f(ctx).
func (f QueryHandlerFunc[Args, Result]) Handle(ctx context.Context, args Args) (Result, error) {
	return f(ctx, args)
}

// NoopQuery is an QueryHandler that does nothing and returns a nil error.
type NoopQuery[Args any, Result any] struct{}

func (NoopQuery[Args, Result]) Handle(context.Context, Args) (Result, error) {
	return *(new(Result)), nil
}

// ========================== Bus ==========================

// Bus is a bus, register CommandHandler and QueryHandler, execute Command and query Query
type Bus interface {

	// RegisterCommand register CommandHandler.
	RegisterCommand(handler any) error

	// RegisterQuery register QueryHandler.
	RegisterQuery(handler any) error

	// Exec executes a command.
	Exec(ctx context.Context, args any) (Metadata, error)

	// Query executes a query.
	Query(ctx context.Context, args any) (any, error)

	// Close bus gracefully.
	Close(ctx context.Context) error
}
