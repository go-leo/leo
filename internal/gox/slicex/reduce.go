package slicex

func Reduce[S ~[]E, E any, R any](s S, initValue R, f func(R, int, E) R) R {
	var r = initValue
	for i, e := range s {
		r = f(r, i, e)
	}
	return r
}
