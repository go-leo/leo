package stringx

import (
	"strings"
)

func Indices(str, substr string) []int {
	s := str
	var indices []int
	subStrLen := len(substr)
	var index int
	for {
		i := strings.Index(s, substr)
		if i < 0 {
			break
		}
		indices = append(indices, i+index)
		s = s[i+subStrLen:]
		index = i + index + subStrLen
	}
	return indices
}
