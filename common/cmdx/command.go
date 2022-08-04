package cmdx

import (
	"context"
	"errors"
)

var (
	ErrNotCommand     = errors.New("not implement Command interface")
	ErrNotUndoCommand = errors.New("not implement UndoCommand interface")
	ErrNotRedoCommand = errors.New("not implement RedoCommand interface")
)

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

type DefaultCommand struct {
	executeFunc func(ctx context.Context) (context.Context, error)
	undoFunc    func(ctx context.Context) (context.Context, error)
	redoFunc    func(ctx context.Context) (context.Context, error)
}

func NewDefaultCommand(
	executeFunc func(ctx context.Context) (context.Context, error),
	undoFunc func(ctx context.Context) (context.Context, error),
	redoFunc func(ctx context.Context) (context.Context, error),
) *DefaultCommand {
	return &DefaultCommand{executeFunc: executeFunc, undoFunc: undoFunc, redoFunc: redoFunc}
}

func (cmd *DefaultCommand) Execute(ctx context.Context) (context.Context, error) {
	if cmd.executeFunc == nil {
		return ctx, ErrNotCommand
	}
	return cmd.executeFunc(ctx)
}

func (cmd *DefaultCommand) Undo(ctx context.Context) (context.Context, error) {
	if cmd.undoFunc == nil {
		return ctx, ErrNotUndoCommand
	}
	return cmd.undoFunc(ctx)
}

func (cmd *DefaultCommand) Redo(ctx context.Context) (context.Context, error) {
	if cmd.redoFunc == nil {
		return ctx, ErrNotRedoCommand
	}
	return cmd.redoFunc(ctx)
}
