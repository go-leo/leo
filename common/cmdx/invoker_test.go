package cmdx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvokerCommandFunc(t *testing.T) {
	invoker := new(Invoker)
	i := 0
	_, err := invoker.Call(context.Background(), CommandFunc(func(ctx context.Context) (context.Context, error) {
		i++
		return ctx, nil
	}))
	assert.NoError(t, err)
	assert.Equal(t, 1, i)

	_, err = invoker.Undo(context.Background())
	assert.ErrorIs(t, err, ErrNotUndoCommand)
	assert.Equal(t, 1, i)

	_, err = invoker.Redo(context.Background())
	assert.ErrorIs(t, err, ErrNotFoundRedoCommand)
	assert.Equal(t, 1, i)
}

type testUndoCommand struct {
	number int
}

func (cmd *testUndoCommand) Execute(ctx context.Context) (context.Context, error) {
	cmd.number++
	return ctx, nil
}

func (cmd *testUndoCommand) Undo(ctx context.Context) (context.Context, error) {
	cmd.number--
	return ctx, nil
}

func TestInvokerUndoCommand(t *testing.T) {
	invoker := new(Invoker)
	cmd := new(testUndoCommand)
	_, err := invoker.Call(context.Background(), cmd)
	assert.NoError(t, err)
	assert.Equal(t, 1, cmd.number)

	_, err = invoker.Undo(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 0, cmd.number)

	_, err = invoker.Redo(context.Background())
	assert.ErrorIs(t, err, ErrNotRedoCommand)
	assert.Equal(t, 0, cmd.number)
}

type testRedoCommand struct {
	number int
}

func (cmd *testRedoCommand) Execute(ctx context.Context) (context.Context, error) {
	cmd.number++
	return ctx, nil
}

func (cmd *testRedoCommand) Undo(ctx context.Context) (context.Context, error) {
	cmd.number--
	return ctx, nil
}

func (cmd *testRedoCommand) Redo(ctx context.Context) (context.Context, error) {
	return cmd.Execute(ctx)
}

func TestInvokerRedoCommand(t *testing.T) {
	invoker := new(Invoker)
	cmd := new(testRedoCommand)
	_, err := invoker.Call(context.Background(), cmd)
	assert.NoError(t, err)
	assert.Equal(t, 1, cmd.number)

	_, err = invoker.Undo(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 0, cmd.number)

	_, err = invoker.Redo(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, cmd.number)
}

func TestInvokerMultiCommand(t *testing.T) {
	invoker := new(Invoker)
	number := 0
	_, err := invoker.Call(context.Background(), NewDefaultCommand(
		func(ctx context.Context) (context.Context, error) {
			number++
			return ctx, nil
		},
		func(ctx context.Context) (context.Context, error) {
			number--
			return ctx, nil
		},
		func(ctx context.Context) (context.Context, error) {
			number++
			return ctx, nil
		},
	))
	assert.NoError(t, err)
	assert.Equal(t, 1, number)

	_, err = invoker.Call(context.Background(), NewDefaultCommand(
		func(ctx context.Context) (context.Context, error) {
			number += 10
			return ctx, nil
		},
		func(ctx context.Context) (context.Context, error) {
			number -= 10
			return ctx, nil
		},
		func(ctx context.Context) (context.Context, error) {
			number += 10
			return ctx, nil
		},
	))
	assert.NoError(t, err)
	assert.Equal(t, 11, number)

	_, err = invoker.Undo(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, number)

	_, err = invoker.Redo(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 11, number)

	_, err = invoker.Undo(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, number)

	_, err = invoker.Redo(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 11, number)

	_, err = invoker.Redo(context.Background())
	assert.ErrorIs(t, ErrNotFoundRedoCommand, err)
	assert.Equal(t, 11, number)

	_, err = invoker.Call(context.Background(), NewDefaultCommand( // nolint
		func(ctx context.Context) (context.Context, error) {
			number += 20
			return ctx, nil
		},
		func(ctx context.Context) (context.Context, error) {
			number -= 20
			return ctx, nil
		},
		func(ctx context.Context) (context.Context, error) {
			number += 20
			return ctx, nil
		},
	))

	_, err = invoker.Redo(context.Background())
	assert.ErrorIs(t, ErrNotFoundRedoCommand, err)
	assert.Equal(t, 31, number)

	_, err = invoker.Undo(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 11, number)

	_, err = invoker.Undo(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, number)

	_, err = invoker.Undo(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 0, number)

	_, err = invoker.Undo(context.Background())
	assert.ErrorIs(t, err, ErrNotFoundUndoCommand)
	assert.Equal(t, 0, number)
}
