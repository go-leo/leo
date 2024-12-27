package cqrs

import (
	"context"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/metadatax"
)

// CommandHandler is a command handler that to update data.
type CommandHandler[Command any, Result any] interface {
	endpointx.Handler[Command, Result]
	CommandName() string
}

// QueryHandler is a query handler that to handlers to read data.
type QueryHandler[Query any, Result any] interface {
	endpointx.Handler[Query, Result]
	QueryName() string
}

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
