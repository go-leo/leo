package cqrs

import (
	"context"
)

// CommandHandler is a command handler that to update data.
type CommandHandler[C any] interface {
	// Handle handles a command.
	Handle(ctx context.Context, command C) error
}

// QueryHandler is a query handler that to handlers to read data.
type QueryHandler[Q any, R any] interface {
	// Handle handles a query.
	Handle(ctx context.Context, query Q) (R, error)
}

// Bus is a bus, register CommandHandler and QueryHandler, execute Command and query Query
type Bus interface {
	// RegisterCommand register CommandHandler.
	RegisterCommand(handler any) error

	// RegisterQuery register QueryHandler.
	RegisterQuery(handler any) error

	// Exec executes a command.
	Exec(ctx context.Context, command any) error

	// AsyncExec executes a command asynchronously.
	AsyncExec(ctx context.Context, command any, errC chan<- error) error

	// Query executes a query.
	Query(ctx context.Context, query any) (any, error)

	// AsyncQuery executes a query asynchronously.
	AsyncQuery(ctx context.Context, query any, resultC chan<- any, errC chan<- error) error

	// Close bus gracefully.
	Close(ctx context.Context) error
}
