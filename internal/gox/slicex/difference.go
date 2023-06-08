package slicex

import "golang.org/x/exp/slices"

// Difference 返回差集
func Difference[S ~[]E, E comparable](s1 S, s2 S) S {
	var r S
	for _, v := range s1 {
		if !slices.Contains(s2, v) {
			r = append(r, v)
		}
	}
	return r
}
