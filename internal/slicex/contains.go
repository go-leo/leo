package slicex

import "golang.org/x/exp/slices"

// ContainsFunc reports whether v is present in s.
func ContainsFunc[E any](s []E, f func(E) bool) bool {
	return slices.ContainsFunc(s, f)
}
