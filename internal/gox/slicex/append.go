package slicex

import "golang.org/x/exp/slices"

func AppendFirst[S ~[]E, E any](s S, e E) S {
	return slices.Insert(s, 0, e)
}

func AppendIfNotContains[S ~[]E, E comparable](s S, v E) S {
	if slices.Contains(s, v) {
		return s
	}
	return append(s, v)
}
