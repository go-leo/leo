package slicex

func Reverse[S ~[]E, E comparable](s S) S {
	if s == nil {
		return nil
	}
	r := make(S, len(s))
	for i, e := range s {
		r[len(s)-1-i] = e
	}
	return r
}
