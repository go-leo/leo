package cqrs

import (
	"context"
	"github.com/go-leo/leo/v3/metadatax"
)

// ========================== Command ==========================

// CommandHandler is a command handler that to update data.
type CommandHandler[Args any] interface {
	Handle(ctx context.Context, args Args) (metadatax.Metadata, error)
}

// The CommandHandlerFunc type is an adapter to allow the use of ordinary functions as CommandHandler.
type CommandHandlerFunc[Args any] func(ctx context.Context, args Args) (metadatax.Metadata, error)

// Handle calls f(ctx).
func (f CommandHandlerFunc[Args]) Handle(ctx context.Context, args Args) (metadatax.Metadata, error) {
	return f(ctx, args)
}

// NoopCommand is an CommandHandler that does nothing and returns a nil error.
type NoopCommand[Args any] struct{}

func (NoopCommand[Args]) Handle(context.Context, Args) (metadatax.Metadata, error) { return nil, nil }

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
	Exec(ctx context.Context, args any) (metadatax.Metadata, error)

	// Query executes a query.
	Query(ctx context.Context, args any) (any, error)

	// Close bus gracefully.
	Close(ctx context.Context) error
}
