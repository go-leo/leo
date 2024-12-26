package cqrs

import (
	"context"
	"github.com/go-leo/gox/cryptox/md5x"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockQuery struct {
	Text string
}

func (MockQuery) QueryName() string {
	return "MockQuery"
}

type MockQueryResult struct {
	Hash string
}

type MockQueryHandler QueryHandler[MockQuery, MockQueryResult]

type mockQueryHandler struct{}

func (m mockQueryHandler) Handle(ctx context.Context, q MockQuery) (MockQueryResult, error) {
	return MockQueryResult{Hash: md5x.TextMD5Hex(q.Text)}, nil
}

func TestQueryHandler(t *testing.T) {
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
