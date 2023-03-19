package slicex

import "golang.org/x/exp/slices"

// NotContains reports whether v is not present in s.
func NotContains[E comparable](s []E, v E) bool {
	return !slices.Contains(s, v)
}

// NotContainsFunc reports whether v is not present in s.
func NotContainsFunc[E any](s []E, f func(E) bool) bool {
	return !slices.ContainsFunc(s, f)
}
