package cqrs

import (
	"context"
	"errors"
	"github.com/go-leo/gox/syncx"
	"reflect"
	"sync"
	"sync/atomic"
)

var (
	ErrHandlerRegistered           = errors.New("cqrs: handler registered")
	ErrBusClosed                   = errors.New("cqrs: bus was closed")
	ErrHandlerInvalid              = errors.New("cqrs: handler is invalid")
	ErrUnimplementedCommandHandler = errors.New("cqrs: handler is not implement CommandHandler")
	ErrUnimplementedQueryHandler   = errors.New("cqrs: handler is not implement QueryHandler")
	ErrCommandHandlerUnregistered  = errors.New("cqrs: command handler unregistered")
	ErrQueryHandlerUnregistered    = errors.New("cqrs: query handler unregistered")
	ErrInvalidCommand              = errors.New("cqrs: invalid command type")
	ErrInvalidQuery                = errors.New("cqrs: invalid query type")
)

var _ Bus = (*SampleBus)(nil)

type SampleBus struct {
	// command handlers
	commandHandlers sync.Map
	// query handlers
	queryHandlers sync.Map
	// wait group
	wg sync.WaitGroup
	// true when bus is in shutdown
	closed atomic.Bool
}

func (b *SampleBus) RegisterCommand(handler any) error {
	if err := b.checkClosed(); err != nil {
		return err
	}
	refHandler, err := newReflectedCommandHandler(handler)
	if err != nil {
		return err
	}
	if _, loaded := b.commandHandlers.LoadOrStore(refHandler.InType(), refHandler); loaded {
		return ErrHandlerRegistered
	}
	return nil
}

func (b *SampleBus) RegisterQuery(handler any) error {
	if err := b.checkClosed(); err != nil {
		return err
	}
	refHandler, err := newReflectedQueryHandler(handler)
	if err != nil {
		return err
	}
	if _, loaded := b.queryHandlers.LoadOrStore(refHandler.InType(), refHandler); loaded {
		return ErrHandlerRegistered
	}
	return nil
}

func (b *SampleBus) Exec(ctx context.Context, command any) error {
	if err := b.checkClosed(); err != nil {
		return err
	}
	handler, err := b.loadCommandHandler(command)
	if err != nil {
		return err
	}
	return handler.Exec(ctx, command)
}

// AsyncExec executes a command asynchronously.
func (b *SampleBus) AsyncExec(ctx context.Context, command any, errC chan<- error) error {
	if err := b.checkClosed(); err != nil {
		return err
	}
	handler, err := b.loadCommandHandler(command)
	if err != nil {
		return err
	}
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()
		if err := handler.Exec(ctx, command); err != nil {
			errC <- err
		}
	}()
	return nil
}

func (b *SampleBus) Query(ctx context.Context, query any) (any, error) {
	if err := b.checkClosed(); err != nil {
		return nil, err
	}
	handler, err := b.loadQueryHandler(query)
	if err != nil {
		return nil, err
	}
	return handler.Query(ctx, query)
}

func (b *SampleBus) AsyncQuery(ctx context.Context, query any, resultC chan<- any, errC chan<- error) error {
	if err := b.checkClosed(); err != nil {
		return err
	}
	handler, err := b.loadQueryHandler(query)
	if err != nil {
		return err
	}
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()
		result, err := handler.Query(ctx, query)
		if err != nil {
			errC <- err
			return
		}
		resultC <- result
	}()
	return nil
}

func (b *SampleBus) Close(ctx context.Context) error {
	if b.closed.CompareAndSwap(false, true) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-syncx.WaitNotify(&b.wg):
			return nil
		}
	}
	return nil
}

func (b *SampleBus) checkClosed() error {
	if b.closed.Load() {
		return ErrBusClosed
	}
	return nil
}

func (b *SampleBus) loadCommandHandler(command any) (*reflectedHandler, error) {
	value, ok := b.commandHandlers.Load(reflect.TypeOf(command))
	if !ok {
		return nil, ErrCommandHandlerUnregistered
	}
	return value.(*reflectedHandler), nil
}

func (b *SampleBus) loadQueryHandler(query any) (*reflectedHandler, error) {
	value, ok := b.queryHandlers.Load(reflect.TypeOf(query))
	if !ok {
		return nil, ErrQueryHandlerUnregistered
	}
	return value.(*reflectedHandler), nil
}
