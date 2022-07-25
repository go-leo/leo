package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	tests := []struct {
		slice    []any
		index    int
		target   int
		expected []any
	}{
		{
			slice:    []any{"1", "2", "3"},
			index:    0,
			target:   0,
			expected: []any{"1", "2", "3"},
		},
		{
			slice:    []any{"1", "2", "3"},
			index:    0,
			target:   1,
			expected: []any{"2", "1", "3"},
		},
		{
			slice:    []any{"1", "2", "3"},
			index:    0,
			target:   2,
			expected: []any{"2", "3", "1"},
		},
		{
			slice:    []any{"1", "2", "3"},
			index:    1,
			target:   1,
			expected: []any{"1", "2", "3"},
		},
		{
			slice:    []any{"1", "2", "3"},
			index:    1,
			target:   2,
			expected: []any{"1", "3", "2"},
		},
		{
			slice:    []any{"1", "2", "3"},
			index:    1,
			target:   0,
			expected: []any{"2", "1", "3"},
		},
		{
			slice:    []any{"1", "2", "3"},
			index:    2,
			target:   2,
			expected: []any{"1", "2", "3"},
		},
		{
			slice:    []any{"1", "2", "3"},
			index:    2,
			target:   1,
			expected: []any{"1", "3", "2"},
		},
		{
			slice:    []any{"1", "2", "3"},
			index:    2,
			target:   0,
			expected: []any{"3", "1", "2"},
		},
	}
	for _, test := range tests {
		actual, err := Move(test.slice, test.index, test.target)
		assert.NoError(t, err)
		assert.ElementsMatch(t, actual, test.expected)
	}
}
