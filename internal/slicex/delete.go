package slicex

import "golang.org/x/exp/slices"

func Remove[S ~[]E, E comparable](s S, v E) S {
	i := slices.Index(s, v)
	if i < 0 {
		return s
	}
	return slices.Delete(s, i, i)
}
