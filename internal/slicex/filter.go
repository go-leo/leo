package slicex

func Filter[S ~[]E, E any](s S, f func(int, E) bool) S {
	var r S
	for i, e := range s {
		if f(i, e) {
			r = append(r, e)
		}
	}
	return r
}
