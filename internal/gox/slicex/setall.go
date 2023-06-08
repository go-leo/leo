package slicex

func SetAll[S ~[]E, E any](s S, f func(int) E) S {
	for i := range s {
		s[i] = f(i)
	}
	return s
}
