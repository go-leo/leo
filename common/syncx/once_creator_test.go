package syncx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type A struct {
	Name string
}

func TestNewOnceMap(t *testing.T) {
	creator, err := NewOnceCreator[*A](func() *A { return new(A) })
	assert.NoError(t, err)
	a := creator.LoadOrCreate("a")
	_a := creator.LoadOrCreate("a")
	assert.True(t, a == _a)

	b := creator.LoadOrCreate("b")
	_b := creator.LoadOrCreate("b")
	assert.True(t, b == _b)

	assert.True(t, a != b)
	assert.True(t, a != _b)
	assert.True(t, _a != b)
	assert.True(t, _a != _b)
}
