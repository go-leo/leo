package cqrs

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-leo/gox/contextx"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/gox/syncx/gopher"
	"github.com/go-leo/gox/syncx/gopher/sample"
	"reflect"
	"sync"
	"sync/atomic"

	"github.com/go-leo/gox/syncx"
)

var (
	// ErrHandlerNil CommandHandler or QueryHandler is nil
	ErrHandlerNil = errors.New("cqrs: handler is nil")

	// ErrRegistered not register CommandHandler or QueryHandler
	ErrRegistered = errors.New("cqrs: handler registered")

	// ErrUnregistered is not register CommandHandler or QueryHandler
	ErrUnregistered = errors.New("cqrs: handler unregistered")

	// ErrArgsNil arguments is nil
	ErrArgsNil = errors.New("cqrs: arguments is nil")

	// ErrBusClosed bus was closed
	ErrBusClosed = errors.New("cqrs: bus was closed")

	// ErrUnimplemented handler is not implement CommandHandler or QueryHandler
	ErrUnimplemented = errors.New("cqrs: handler is not implement CommandHandler or QueryHandler")
)

var _ Bus = (*defaultBus)(nil)

type defaultBus struct {
	handlers   *sync.Map
	wg         *sync.WaitGroup
	inShutdown *atomic.Bool // true when bus is in shutdown
	options    *option
}

func (b *defaultBus) RegisterCommand(handler any) error {
	if err := b.checkHandler(handler); err != nil {
		return err
	}
	handlerRef, err := newReflectedHandler(handler, "command")
	if err != nil {
		return err
	}
	if _, loaded := b.handlers.LoadOrStore(handlerRef.InType(), handlerRef); loaded {
		return ErrRegistered
	}
	return nil
}

func (b *defaultBus) RegisterQuery(handler any) error {
	if err := b.checkHandler(handler); err != nil {
		return err
	}
	handlerRef, err := newReflectedHandler(handler, "query")
	if err != nil {
		return err
	}
	if _, loaded := b.handlers.LoadOrStore(handlerRef.InType(), handlerRef); loaded {
		return ErrRegistered
	}
	return nil
}

func (b *defaultBus) Exec(ctx context.Context, args any) error {
	if err := b.checkArgs(args); err != nil {
		return err
	}
	info, err := b.loadHandler(args)
	if err != nil {
		return err
	}
	return info.Exec(ctx, args)
}

func (b *defaultBus) Query(ctx context.Context, args any) (any, error) {
	if err := b.checkArgs(args); err != nil {
		return nil, err
	}
	info, err := b.loadHandler(args)
	if err != nil {
		return nil, err
	}
	return info.Query(ctx, args)
}

func (b *defaultBus) AsyncExec(ctx context.Context, args any) (Future, error) {
	if err := b.checkArgs(args); err != nil {
		return nil, err
	}
	info, err := b.loadHandler(args)
	if err != nil {
		return nil, err
	}
	f := newDefaultFuture()
	b.wg.Add(1)
	if err := b.options.Pool.Go(func() {
		defer b.wg.Done()
		f.OnExec(info.Exec(ctx, args))
	}); err != nil {
		return nil, err
	}
	return f, nil
}

func (b *defaultBus) AsyncQuery(ctx context.Context, args any) (Future, error) {
	if err := b.checkArgs(args); err != nil {
		return nil, err
	}
	info, err := b.loadHandler(args)
	if err != nil {
		return nil, err
	}
	f := newDefaultFuture()
	b.wg.Add(1)
	if err := b.options.Pool.Go(func() {
		defer b.wg.Done()
		f.OnQuery(info.Query(ctx, args))
	}); err != nil {
		return nil, err
	}
	return f, nil
}

func (b *defaultBus) Close(ctx context.Context) error {
	if b.inShutdown.CompareAndSwap(false, true) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-syncx.WaitNotify(b.wg):
			return nil
		}
	}
	return ErrBusClosed
}

func (b *defaultBus) shuttingDown() bool {
	return b.inShutdown.Load()
}

func (b *defaultBus) checkHandler(handler any) error {
	if handler == nil {
		return ErrHandlerNil
	}
	if b.shuttingDown() {
		return ErrBusClosed
	}
	return nil
}

func (b *defaultBus) checkArgs(args any) error {
	if args == nil {
		return ErrArgsNil
	}
	if b.shuttingDown() {
		return ErrBusClosed
	}
	return nil
}

func (b *defaultBus) loadHandler(args any) (*reflectedHandler, error) {
	value, ok := b.handlers.Load(reflect.TypeOf(args))
	if !ok {
		return nil, ErrUnregistered
	}
	info := value.(*reflectedHandler)
	return info, nil
}

type option struct {
	Pool gopher.Gopher
}

func newOption(opts ...Option) *option {
	o := &option{}
	for _, opt := range opts {
		opt(o)
	}
	if o.Pool == nil {
		o.Pool = sample.Gopher{}
	}
	return o
}

type Option func(*option)

func Pool(pool gopher.Gopher) Option {
	return func(o *option) {
		o.Pool = pool
	}
}

func NewBus(opts ...Option) Bus {
	return &defaultBus{
		handlers:   &sync.Map{},
		wg:         &sync.WaitGroup{},
		inShutdown: &atomic.Bool{},
		options:    newOption(opts...),
	}
}

type reflectedHandler struct {
	value  reflect.Value
	method reflect.Method
	inType reflect.Type
}

func (handler *reflectedHandler) Exec(ctx context.Context, args any) error {
	resultValues := handler.method.Func.Call(
		[]reflect.Value{
			handler.value,
			reflect.ValueOf(ctx),
			reflect.ValueOf(args),
		})
	err := resultValues[0].Interface()
	if err != nil {
		return err.(error)
	}
	return nil
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

func newReflectedHandler(handler any, kind string) (*reflectedHandler, error) {
	handlerVal := reflect.ValueOf(handler)
	method, ok := handlerVal.Type().MethodByName("Handle")
	if !ok {
		return nil, ErrUnimplemented
	}
	switch kind {
	case "command":
		if method.Type.NumIn() != 3 {
			return nil, ErrUnimplemented
		}
		if !method.Type.In(1).Implements(contextx.ContextType) {
			return nil, ErrUnimplemented
		}
		if method.Type.NumOut() != 1 {
			return nil, ErrUnimplemented
		}
		if !method.Type.Out(0).Implements(errorx.ErrorType) {
			return nil, ErrUnimplemented
		}
	case "query":
		if method.Type.NumIn() != 3 {
			return nil, ErrUnimplemented
		}
		if !method.Type.In(1).Implements(contextx.ContextType) {
			return nil, ErrUnimplemented
		}
		if method.Type.NumOut() != 2 {
			return nil, ErrUnimplemented
		}
		if !method.Type.Out(1).Implements(errorx.ErrorType) {
			return nil, ErrUnimplemented
		}
	default:
		return nil, fmt.Errorf("unknown kind %s", kind)
	}
	inType := method.Type.In(2)
	return &reflectedHandler{
		value:  handlerVal,
		method: method,
		inType: inType,
	}, nil
}

type defaultFuture struct {
	resC chan any
	errC chan error
}

func (f *defaultFuture) Get(ctx context.Context) (any, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-f.resC:
		return res, nil
	case err := <-f.errC:
		return nil, err
	}
}

func (f *defaultFuture) OnExec(err error) {
	f.errC <- err
}

func (f *defaultFuture) OnQuery(res any, err error) {
	if err != nil {
		f.errC <- err
		return
	}
	f.resC <- res
}

func newDefaultFuture() *defaultFuture {
	return &defaultFuture{
		resC: make(chan any, 1),
		errC: make(chan error, 1),
	}
}
