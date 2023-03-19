package slicex

import "golang.org/x/exp/slices"

func Uniq[S ~[]E, E comparable](s S) S {
	if s == nil {
		return nil
	}
	length := len(s)
	r := make(S, 0, length)
	if length <= 8 {
		for _, v := range s {
			if !slices.Contains(r, v) {
				r = append(r, v)
			}
		}
		return r
	}
	m := make(map[E]struct{}, length)
	for _, v := range s {
		m[v] = struct{}{}
	}
	for k := range m {
		r = append(r, k)
	}
	return r
}
