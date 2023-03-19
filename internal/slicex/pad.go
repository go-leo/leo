package slicex

import "golang.org/x/exp/slices"

// PadStart 如果slice长度小于 length 则在左侧填充val。
func PadStart[S ~[]E, E any](s S, size int, val E) S {
	if size <= len(s) {
		return slices.Clone(s)
	}
	r := make(S, 0, size)
	for i := 0; i < (size - len(s)); i++ {
		r = append(r, val)
	}
	r = append(r, s...)
	return s
}

// PadEnd 如果slice长度小于 length 则在右侧填充val。
func PadEnd[S ~[]E, E any](s S, size int, val E) S {
	if size <= len(s) {
		return slices.Clone(s)
	}
	r := make(S, 0, size)
	r = append(r, s...)
	for i := 0; i < (size - len(s)); i++ {
		r = append(r, val)
	}
	return r
}
