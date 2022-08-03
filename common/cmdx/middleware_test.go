package cmdx

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type countKey struct{}

type CountCommand struct {
	count int
}

func (cmd *CountCommand) Execute(ctx context.Context) (context.Context, error) {
	cmd.count++
	log.Println("Execute count = ", cmd.count)
	return context.WithValue(ctx, countKey{}, cmd.count), nil
}

func TestSingleCommandChain(t *testing.T) {
	cmd := &CountCommand{}
	command := Chain(cmd)
	ctx, err := command.Execute(context.Background())
	assert.NoError(t, err)
	count := ctx.Value(countKey{})
	t.Log(count)
}

func TestSingleCommandFuncChain(t *testing.T) {
	type numberKey struct{}
	numberCommand := func(ctx context.Context) (context.Context, error) {
		return context.WithValue(ctx, numberKey{}, 1), nil
	}
	command := Chain(CommandFunc(numberCommand))
	ctx, err := command.Execute(context.Background())
	assert.NoError(t, err)
	number := ctx.Value(numberKey{})
	t.Log(number)
}

type logKey struct{}

type Log1Middleware struct{}

func (l *Log1Middleware) Decorate(cmd Command) Command {
	return CommandFunc(func(ctx context.Context) (context.Context, error) {
		logger := log.New(os.Stdout, "log", log.LstdFlags|log.Lshortfile)
		logger.Println("start log 1")
		defer logger.Println("end log 1")
		ctx = context.WithValue(ctx, logKey{}, logger)
		return cmd.Execute(ctx)
	})
}

func TestMiddlewareChain(t *testing.T) {
	cmd := &CountCommand{}
	command := Chain(cmd, &Log1Middleware{}, MiddlewareFunc(func(cmd Command) Command {
		return CommandFunc(func(ctx context.Context) (context.Context, error) {
			logger, ok := ctx.Value(logKey{}).(*log.Logger)
			if ok {
				logger.Println("start log 2")
				defer logger.Println("end log 2")
			}
			return cmd.Execute(ctx)
		})
	}))
	ctx, err := command.Execute(context.Background())
	assert.NoError(t, err)
	count := ctx.Value(countKey{})
	t.Log(count)
}

func TestMiddlewareWithErrorChain(t *testing.T) {
	cmd := &CountCommand{}
	err := errors.New("occur error")
	log2 := func(cmd Command) Command {
		return CommandFunc(func(ctx context.Context) (context.Context, error) {
			logger, ok := ctx.Value(logKey{}).(*log.Logger)
			if ok {
				logger.Println("start log 2")
				defer logger.Println("end log 2")
			}
			return cmd.Execute(ctx)
		})
	}
	errMdw := func(cmd Command) Command {
		return CommandFunc(func(ctx context.Context) (context.Context, error) {
			return ctx, err
		})
	}
	command := Chain(cmd, &Log1Middleware{}, MiddlewareFunc(errMdw), MiddlewareFunc(log2))
	ctx, e := command.Execute(context.Background())
	assert.ErrorIs(t, e, err)
	count := ctx.Value(countKey{})
	assert.Nil(t, count)
}
