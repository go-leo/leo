package slicex

import (
	"testing"
)

func TestRemoveElementString(t *testing.T) {
	var tests = []struct {
		slice    []string
		element  string
		expected []string
	}{
		{
			slice:    nil,
			element:  "a",
			expected: nil,
		},
		{
			slice:    []string{},
			element:  "a",
			expected: []string{},
		},
		{
			slice:    []string{"a"},
			element:  "b",
			expected: []string{"a"},
		},
		{
			slice:    []string{"a", "b"},
			element:  "a",
			expected: []string{"b"},
		},
		{
			slice:    []string{"a", "b", "a"},
			element:  "a",
			expected: []string{"b", "a"},
		},
	}
	for _, test := range tests {
		actual := RemoveElementString(test.slice, test.element)
		if len(actual) != len(test.expected) {
			t.Errorf("expected is %v, but actual is %v", test.expected, actual)
			continue
		}
		for i := range actual {
			if actual[i] != test.expected[i] {
				t.Errorf("expected is %v, but actual is %v", test.expected, actual)
				break
			}
		}
	}
}
