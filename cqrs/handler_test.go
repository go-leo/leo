package cqrs

import (
	"context"
	"github.com/go-leo/gox/cryptox/md5x"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type commandHandlerWithoutHandle struct{}

type commandHandlerWithInvalidParamsCount struct{}

func (h commandHandlerWithInvalidParamsCount) Handle(ctx context.Context) error {
	return nil
}

type commandHandlerWithoutContentParams struct{}

func (h commandHandlerWithoutContentParams) Handle(ctx any, cmd any) error {
	return nil
}

type commandHandlerWithInvalidReturnCount struct{}

func (h commandHandlerWithInvalidReturnCount) Handle(ctx context.Context, cmd any) (any, error) {
	return nil, nil
}

type commandHandlerWithoutErrorReturn struct{}

func (h commandHandlerWithoutErrorReturn) Handle(ctx context.Context, cmd any) any {
	return nil
}

type commandHandlerWithValidSignature struct{}

func (h commandHandlerWithValidSignature) Handle(ctx context.Context, cmd any) error {
	return nil
}

func TestNewReflectedCommandHandler(t *testing.T) {
	tests := []struct {
		handler any
		wantErr error
	}{
		{commandHandlerWithoutHandle{}, ErrUnimplemented},
		{commandHandlerWithInvalidParamsCount{}, ErrUnimplemented},
		{commandHandlerWithoutContentParams{}, ErrUnimplemented},
		{commandHandlerWithInvalidReturnCount{}, ErrUnimplemented},
		{commandHandlerWithoutErrorReturn{}, ErrUnimplemented},
		{commandHandlerWithValidSignature{}, nil},
	}

	for _, tt := range tests {
		handler, err := newReflectedCommandHandler(tt.handler)
		if tt.wantErr != nil {
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Nil(t, handler)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, handler)
		}
	}
}

type queryHandlerWithoutHandle struct{}

type queryHandlerWithInvalidParamsCount struct{}

func (h queryHandlerWithInvalidParamsCount) Handle(ctx context.Context) (any, error) {
	return nil, nil
}

type queryHandlerWithoutContentParams struct{}

func (h queryHandlerWithoutContentParams) Handle(ctx any, cmd any) (any, error) {
	return nil, nil
}

type queryHandlerWithInvalidReturn struct{}

func (h queryHandlerWithInvalidReturn) Handle(ctx context.Context, req string) error {
	return nil
}

type queryHandlerWithInvalidSecondReturn struct{}

func (h queryHandlerWithInvalidSecondReturn) Handle(ctx context.Context, req string) (string, string) {
	return "", ""
}

type queryHandlerWithValidSignature struct{}

func (h queryHandlerWithValidSignature) Handle(ctx context.Context, req any) (any, error) {
	return nil, nil
}

func TestNewReflectedQueryHandler(t *testing.T) {
	tests := []struct {
		handler any
		wantErr error
	}{
		{queryHandlerWithoutHandle{}, ErrUnimplemented},
		{queryHandlerWithInvalidParamsCount{}, ErrUnimplemented},
		{queryHandlerWithoutContentParams{}, ErrUnimplemented},
		{queryHandlerWithInvalidReturn{}, ErrUnimplemented},
		{queryHandlerWithInvalidSecondReturn{}, ErrUnimplemented},
		{queryHandlerWithValidSignature{}, nil},
	}

	for _, tt := range tests {
		handler, err := newReflectedQueryHandler(tt.handler)
		if tt.wantErr != nil {
			assert.ErrorIs(t, err, tt.wantErr)
		} else {
			require.NoError(t, err)
			assert.NotNil(t, handler)
		}
	}
}

type MockQuery struct {
	Text string
}

type MockQueryResult struct {
	Hash string
}

type MockQueryHandler QueryHandler[MockQuery, MockQueryResult]

type mockQueryHandler struct{}

func (m mockQueryHandler) Handle(ctx context.Context, q MockQuery) (MockQueryResult, error) {
	return MockQueryResult{Hash: md5x.TextMD5Hex(q.Text)}, nil
}

func TestQuery(t *testing.T) {
	handler, err := newReflectedQueryHandler(mockQueryHandler{})
	assert.NoError(t, err)
	text := "hello"
	result, err := handler.Query(context.Background(), MockQuery{Text: text})
	assert.NoError(t, err)
	assert.Equal(t, md5x.TextMD5Hex(text), result.(MockQueryResult).Hash)

	handler, err = newReflectedQueryHandler(&mockQueryHandler{})
	assert.NoError(t, err)
	text = "world"
	result, err = handler.Query(context.Background(), MockQuery{Text: text})
	assert.NoError(t, err)
	assert.Equal(t, md5x.TextMD5Hex(text), result.(MockQueryResult).Hash)
}
