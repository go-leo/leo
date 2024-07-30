package cqrs

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-leo/gox/contextx"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/gox/syncx/groupx"
	"github.com/go-leo/leo/v3/metadatax"
	"reflect"
	"sync"
	"sync/atomic"
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
	handlerRef, err := b.newReflectedHandler(handler, "command")
	if err != nil {
		return err
	}
	return b.registerHandler(handlerRef)
}

func (b *defaultBus) RegisterQuery(handler any) error {
	if err := b.registerCheck(handler); err != nil {
		return err
	}
	handlerRef, err := b.newReflectedHandler(handler, "query")
	if err != nil {
		return err
	}
	return b.registerHandler(handlerRef)
}

func (b *defaultBus) Exec(ctx context.Context, args any) (metadatax.Metadata, error) {
	if err := b.invokeCheck(args); err != nil {
		return nil, err
	}
	info, err := b.loadHandler(args)
	if err != nil {
		return nil, err
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
		case <-groupx.WaitNotify(b.wg):
			return nil
		}
	}
	return nil
}

func (b *defaultBus) registerHandler(handlerRef *reflectedHandler) error {
	if _, loaded := b.handlers.LoadOrStore(handlerRef.InType(), handlerRef); loaded {
		return errors.New("cqrs: handler registered")
	}
	return nil
}

func (b *defaultBus) checkClosed() error {
	if b.closed.Load() {
		return errors.New("cqrs: bus was closed")
	}
	return nil
}

func (b *defaultBus) registerCheck(handler any) error {
	if handler == nil {
		return errors.New("cqrs: handler is nil")
	}
	return b.checkClosed()
}

func (b *defaultBus) invokeCheck(args any) error {
	if args == nil {
		return errors.New("cqrs: arguments is nil")
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

func (b *defaultBus) newReflectedHandler(handler any, kind string) (*reflectedHandler, error) {
	handlerVal := reflect.ValueOf(handler)
	method, ok := handlerVal.Type().MethodByName("Handle")
	if !ok {
		return nil, b.unimplemented()
	}
	switch kind {
	case "command":
		if method.Type.NumIn() != 3 {
			return nil, b.unimplemented()
		}
		if !method.Type.In(1).Implements(contextx.ContextType) {
			return nil, b.unimplemented()
		}
		if method.Type.NumOut() != 2 {
			return nil, b.unimplemented()
		}
		if !method.Type.Out(0).Implements(metadatax.Type) {
			return nil, b.unimplemented()
		}
		if !method.Type.Out(1).Implements(errorx.ErrorType) {
			return nil, b.unimplemented()
		}
	case "query":
		if method.Type.NumIn() != 3 {
			return nil, b.unimplemented()
		}
		if !method.Type.In(1).Implements(contextx.ContextType) {
			return nil, b.unimplemented()
		}
		if method.Type.NumOut() != 2 {
			return nil, b.unimplemented()
		}
		if !method.Type.Out(1).Implements(errorx.ErrorType) {
			return nil, b.unimplemented()
		}
	default:
		return nil, fmt.Errorf("cqrs: unknown kind %s", kind)
	}
	inType := method.Type.In(2)
	return &reflectedHandler{
		value:  handlerVal,
		method: method,
		inType: inType,
	}, nil
}

func (b *defaultBus) unimplemented() error {
	return errors.New("cqrs: handler is not implement CommandHandler or QueryHandler")
}

func NewBus() Bus {
	return &defaultBus{
		handlers: &sync.Map{},
		wg:       &sync.WaitGroup{},
		closed:   &atomic.Bool{},
	}
}

type reflectedHandler struct {
	value  reflect.Value
	method reflect.Method
	inType reflect.Type
}

func (handler *reflectedHandler) Exec(ctx context.Context, args any) (metadatax.Metadata, error) {
	resultValues := handler.method.Func.Call(
		[]reflect.Value{
			handler.value,
			reflect.ValueOf(ctx),
			reflect.ValueOf(args),
		})
	err := resultValues[1].Interface()
	if err != nil {
		return nil, err.(error)
	}
	md := resultValues[0].Interface()
	if md == nil {
		return nil, nil
	}
	return md.(metadatax.Metadata), nil
}

func (handler *reflectedHandler) Query(ctx context.Context, args any) (any, error) {
	resultValues := handler.method.Func.Call(
		[]reflect.Value{
			handler.value,
			reflect.ValueOf(ctx),
			reflect.ValueOf(args),
		})
	err := resultValues[1].Interface()
	if err != nil {
		return nil, err.(error)
	}
	return resultValues[0].Interface(), nil
}

func (handler *reflectedHandler) InType() reflect.Type {
	return handler.inType
}
