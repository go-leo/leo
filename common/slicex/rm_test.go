package slicex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveString(t *testing.T) {
	var tests = []struct {
		slice         []string
		index         int
		expected      []string
		expectedError string
	}{
		{
			slice:         nil,
			index:         0,
			expected:      nil,
			expectedError: fmt.Sprintf(indexOutOfBoundsErrorFormat, 0, 0),
		},
		{
			slice:         []string{},
			index:         0,
			expected:      []string{},
			expectedError: fmt.Sprintf(indexOutOfBoundsErrorFormat, 0, 0),
		},
		{
			slice:         []string{"0"},
			index:         0,
			expected:      []string{},
			expectedError: "",
		},
		{
			slice:         []string{"0", "2"},
			index:         1,
			expected:      []string{"0"},
			expectedError: "",
		},
		{
			slice:         []string{"0", "2", "4"},
			index:         3,
			expected:      nil,
			expectedError: fmt.Sprintf(indexOutOfBoundsErrorFormat, 3, 3),
		},
		{
			slice:         []string{"this", "is", "test", "slice"},
			index:         3,
			expected:      []string{"this", "is", "test"},
			expectedError: "",
		},
	}
	for _, test := range tests {
		actual, err := RemoveString(test.slice, test.index)
		if err != nil && err.Error() != test.expectedError {
			t.Errorf("expected has error %v, but actual error is %v", test.expectedError, err)
			continue
		}
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

	slice := []string{"this", "is", "test", "slice"}
	index := 3
	expected := []string{"this", "is", "test"}

	res := append(slice[:index], slice[index+1:]...)
	assert.Equal(t, expected, res)
}
