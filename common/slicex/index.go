package slicex

import "golang.org/x/exp/slices"

// Index wrap slices.Index function
func Index[E comparable](s []E, v E) int {
	return slices.Index(s, v)
}

// IndexFunc wrap slices.IndexFunc function
func IndexFunc[E any](s []E, f func(E) bool) int {
	return slices.IndexFunc(s, f)
}
