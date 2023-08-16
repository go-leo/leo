package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	m1 := map[string]any{
		"hello": map[string]any{
			"largenum": 765432101234567,
			"pop":      37890,
			"world":    []any{"us", "uk", "fr", "de"},
		},
	}
	m2 := map[string]any{
		"hello": map[string]any{
			"pop": 1234,
		},
		"world": map[any]any{
			"rock": 345,
		},
	}

	m := make(map[string]any)
	err := merge(m1, m)
	assert.NoError(t, err)

	m1Hello := m1["hello"].(map[string]any)
	mHello := m["hello"].(map[string]any)
	assert.Equal(t, m1Hello["largenum"], mHello["largenum"])
	assert.Equal(t, m1Hello["pop"], mHello["pop"])
	assert.Equal(t, m1Hello["world"], mHello["world"])

	err = merge(m2, m)
	assert.NoError(t, err)

	m2Hello := m2["hello"].(map[string]any)
	m2World := m2["world"].(map[any]any)
	mWorld := m["world"].(map[any]any)
	assert.Equal(t, m1Hello["largenum"], mHello["largenum"])
	assert.Equal(t, m2Hello["pop"], mHello["pop"])
	assert.Equal(t, m1Hello["world"], mHello["world"])
	assert.Equal(t, m2World["rock"], mWorld["rock"])
}
