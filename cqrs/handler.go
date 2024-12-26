package cqrs

import (
	"context"
	"github.com/go-leo/gox/contextx"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/leo/v3/metadatax"
	"reflect"
)

type reflectedHandler struct {
	receiver reflect.Value
	method   reflect.Method
	inType   reflect.Type
}

func (handler *reflectedHandler) Exec(ctx context.Context, args any) (metadatax.Metadata, error) {
	resultValues := handler.method.Func.Call(
		[]reflect.Value{
			handler.receiver,
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
			handler.receiver,
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
	if method.Type.NumOut() != 2 {
		return nil, ErrUnimplemented
	}
	if !method.Type.Out(0).Implements(metadatax.Type) {
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
