package slicex

import "golang.org/x/exp/slices"

func Remove[S ~[]E, E comparable](s S, v E) S {
	if IsEmpty(s) {
		return slices.Clone(s)
	}
	return Delete(s, slices.Index(s, v))
}
