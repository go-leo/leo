package cmdx

import "context"

// A Command encapsulates a unit of processing work to be performed.
type Command interface {
	// Execute a unit of processing work to be performed
	Execute(ctx context.Context) (context.Context, error)
}

// The CommandFunc type is an adapter to allow the use of ordinary functions as Command.
// If f is a function with the appropriate signature, CommandFunc(f) is a Command that calls f.
type CommandFunc func(ctx context.Context) (context.Context, error)

// Execute calls f(ctx).
func (f CommandFunc) Execute(ctx context.Context) (context.Context, error) {
	return f(ctx)
}

// UndoCommand undo a Command
type UndoCommand interface {
	Undo(ctx context.Context) (context.Context, error)
}

// RedoCommand redo a Command
type RedoCommand interface {
	Redo(ctx context.Context) (context.Context, error)
}
