package slicex

import "golang.org/x/exp/slices"

func AppendIfNotContains[S ~[]E, E comparable](s S, v E) S {
	if slices.Contains(s, v) {
		return s
	}
	return append(s, v)
}
