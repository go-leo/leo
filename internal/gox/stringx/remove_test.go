package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveChar(t *testing.T) {
	tests := []struct {
		title     string
		candidate string
		chars     string
		expected  string
	}{
		{
			title:     "when we have one char to replace",
			candidate: "hello foobar",
			chars:     "a",
			expected:  "hello foobr",
		},
		{
			title:     "when we have multiple chars to replace",
			candidate: "hello foobar",
			chars:     "all",
			expected:  "heo foobr",
		},
		{
			title:     "when we have no chars to replace",
			candidate: "hello foobar",
			chars:     "x",
			expected:  "hello foobar",
		},
		{
			title:     "when we have an empty string",
			candidate: "",
			chars:     "x",
			expected:  "",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			assert.Equal(t, test.expected, Remove(test.candidate, test.chars))
		})
	}
}
