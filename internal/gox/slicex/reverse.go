package slicex

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/mathx"
)

// Reverse the order of the given slice in the given range. if range is not supply, reverse all.
func Reverse[S ~[]E, E any](s S, indexes ...int) S {
	if s == nil {
		return nil
	}
	if len(indexes) <= 0 {
		indexes = []int{0, len(s)}
	}
	startIndexInclusive := indexes[0]
	endIndexExclusive := indexes[1]
	i := mathx.Max(startIndexInclusive, 0)
	j := mathx.Min(len(s), endIndexExclusive) - 1
	var tmp E
	for j > i {
		tmp = s[j]
		s[j] = s[i]
		s[i] = tmp
		j--
		i++
	}
	return s
}
