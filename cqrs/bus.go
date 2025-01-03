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
	ErrHandlerRegistered = errors.New("cqrs: handler registered")
	ErrBusClosed         = errors.New("cqrs: bus was closed")
	ErrHandlerNil        = errors.New("cqrs: handler is nil")
	ErrArgNil            = errors.New("cqrs: argument is nil")
	ErrUnimplemented     = errors.New("cqrs: handler is not implement CommandHandler or QueryHandler")
)

var _ Bus = (*defaultBus)(nil)

type defaultBus struct {
	handlers *sync.Map
	wg       *sync.WaitGroup
	closed   *atomic.Bool // true when bus is in shutdown
}

func (b *defaultBus) RegisterCommand(handler any) error {
	if err := b.registerCheck(handler); err != nil {
		return err
	}
	handlerRef, err := newReflectedCommandHandler(handler)
	if err != nil {
		return err
	}
	return b.registerHandler(handlerRef)
}

func (b *defaultBus) RegisterQuery(handler any) error {
	if err := b.registerCheck(handler); err != nil {
		return err
	}
	handlerRef, err := newReflectedQueryHandler(handler)
	if err != nil {
		return err
	}
	return b.registerHandler(handlerRef)
}

func (b *defaultBus) Exec(ctx context.Context, args any) error {
	if err := b.invokeCheck(args); err != nil {
		return err
	}
	info, err := b.loadHandler(args)
	if err != nil {
		return err
	}
	return info.Exec(ctx, args)
}

func (b *defaultBus) Query(ctx context.Context, args any) (any, error) {
	if err := b.invokeCheck(args); err != nil {
		return nil, err
	}
	info, err := b.loadHandler(args)
	if err != nil {
		return nil, err
	}
	return info.Query(ctx, args)
}

func (b *defaultBus) Close(ctx context.Context) error {
	if b.closed.CompareAndSwap(false, true) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-syncx.WaitNotify(b.wg):
			return nil
		}
	}
	return nil
}

func (b *defaultBus) registerHandler(handlerRef *reflectedHandler) error {
	if _, loaded := b.handlers.LoadOrStore(handlerRef.InType(), handlerRef); loaded {
		return ErrHandlerRegistered
	}
	return nil
}

func (b *defaultBus) checkClosed() error {
	if b.closed.Load() {
		return ErrBusClosed
	}
	return nil
}

func (b *defaultBus) registerCheck(handler any) error {
	if handler == nil {
		return ErrHandlerNil
	}
	return b.checkClosed()
}

func (b *defaultBus) invokeCheck(args any) error {
	if args == nil {
		return ErrArgNil
	}
	return b.checkClosed()
}

func (b *defaultBus) loadHandler(args any) (*reflectedHandler, error) {
	value, ok := b.handlers.Load(reflect.TypeOf(args))
	if !ok {
		return nil, errors.New("cqrs: handler unregistered")
	}
	info := value.(*reflectedHandler)
	return info, nil
}

func NewBus() Bus {
	return &defaultBus{
		handlers: &sync.Map{},
		wg:       &sync.WaitGroup{},
		closed:   &atomic.Bool{},
	}
}
