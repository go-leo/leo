package cqrs

import (
	"context"
	"github.com/go-leo/gox/contextx"
	"github.com/go-leo/gox/errorx"
	"reflect"
)

type reflectedHandler struct {
	receiver reflect.Value
	method   reflect.Method
	inType   reflect.Type
}

func (handler *reflectedHandler) Exec(ctx context.Context, command any) error {
	resultValues := handler.method.Func.Call(
		[]reflect.Value{
			handler.receiver,
			reflect.ValueOf(ctx),
			reflect.ValueOf(command),
		})
	err := resultValues[1].Interface()
	if err != nil {
		return err.(error)
	}
	return nil
}

func (handler *reflectedHandler) Query(ctx context.Context, query any) (any, error) {
	resultValues := handler.method.Func.Call(
		[]reflect.Value{
			handler.receiver,
			reflect.ValueOf(ctx),
			reflect.ValueOf(query),
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

// newReflectedCommandHandler creates a reflectedHandler for a CommandHandler.
func newReflectedCommandHandler(handler any) (*reflectedHandler, error) {
	handlerVal := reflect.ValueOf(handler)
	method, ok := handlerVal.Type().MethodByName("Handle")
	if !ok {
		return nil, ErrUnimplemented
	}
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
	inType := method.Type.In(2)
	return &reflectedHandler{
		receiver: handlerVal,
		method:   method,
		inType:   inType,
	}, nil
}

// newReflectedQueryHandler creates a reflectedHandler for a QueryHandler.
func newReflectedQueryHandler(handler any) (*reflectedHandler, error) {
	handlerVal := reflect.ValueOf(handler)
	method, ok := handlerVal.Type().MethodByName("Handle")
	if !ok {
		return nil, ErrUnimplemented
	}
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
	inType := method.Type.In(2)
	return &reflectedHandler{
		receiver: handlerVal,
		method:   method,
		inType:   inType,
	}, nil
}
