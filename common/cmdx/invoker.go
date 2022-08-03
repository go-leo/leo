package cmdx

import (
	"container/list"
	"context"
	"sync"
)

type Invoker struct {
	sync.Mutex
	undoStack *list.List
	redoStack *list.List
}

func (invoker *Invoker) Call(ctx context.Context, cmd Command) (context.Context, error) {
	ctx, err := cmd.Execute(ctx)
	if err != nil {
		return ctx, err
	}
	invoker.Lock()
	defer invoker.Unlock()
	invoker.pushUndoCommand(cmd)
	return ctx, nil
}

func (invoker *Invoker) Undo(ctx context.Context) (context.Context, error) {
	invoker.Lock()
	defer invoker.Unlock()
	if invoker.isUndoCommandsEmpty() {
		return ctx, nil
	}
	cmd, ok := invoker.popUndoCommand()
	if !ok {
		return ctx, nil
	}
	undo, ok := cmd.(UndoCommand)
	if !ok {
		return ctx, nil
	}
	ctx, err := undo.Undo(ctx)
	if err != nil {
		return ctx, err
	}
	invoker.pushRedoCommand(cmd)
	return ctx, nil
}

func (invoker *Invoker) Redo(ctx context.Context) (context.Context, error) {
	invoker.Lock()
	defer invoker.Unlock()
	if invoker.isRedoCommandsEmpty() {
		return ctx, nil
	}
	cmd, ok := invoker.popRedoCommand()
	if !ok {
		return ctx, nil
	}
	redo, ok := cmd.(RedoCommand)
	if !ok {
		return ctx, nil
	}
	ctx, err := redo.Redo(ctx)
	if err != nil {
		return ctx, err
	}
	invoker.pushUndoCommand(cmd)
	return ctx, nil
}

func (invoker *Invoker) undoCommands() *list.List {
	if invoker.undoStack == nil {
		invoker.undoStack = list.New()
	}
	return invoker.undoStack
}

func (invoker *Invoker) isUndoCommandsEmpty() bool {
	return invoker.undoCommands().Len() <= 0
}

func (invoker *Invoker) pushUndoCommand(cmd Command) {
	_ = invoker.undoCommands().PushBack(cmd)
}

func (invoker *Invoker) popUndoCommand() (Command, bool) {
	element := invoker.undoCommands().Back()
	if element == nil {
		return nil, false
	}
	cmd, ok := invoker.undoCommands().Remove(element).(Command)
	return cmd, ok
}

func (invoker *Invoker) redoCommands() *list.List {
	if invoker.redoStack == nil {
		invoker.redoStack = list.New()
	}
	return invoker.redoStack
}

func (invoker *Invoker) isRedoCommandsEmpty() bool {
	return invoker.redoCommands().Len() <= 0
}

func (invoker *Invoker) pushRedoCommand(cmd Command) {
	invoker.redoCommands().PushBack(cmd)
}

func (invoker *Invoker) popRedoCommand() (Command, bool) {
	element := invoker.redoCommands().Back()
	if element == nil {
		return nil, false
	}
	cmd, ok := invoker.redoCommands().Remove(element).(Command)
	return cmd, ok
}
