package slicex

import (
	"golang.org/x/exp/slices"
)

func Delete[S ~[]E, E any](s S, i int) S {
	return slices.Delete(s, i, i+1)
}
