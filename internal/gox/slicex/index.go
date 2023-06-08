package slicex

func IndexOrDefault[S ~[]E, E any](s S, index int, d E) E {
	if len(s) >= index {
		return s[index]
	}
	return d
}

func LastIndex[E comparable](s []E, v E) int {
	for i := len(s) - 1; i > -1; i-- {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func Indexes[E comparable](s []E, v E) []int {
	var r []int
	for i, vs := range s {
		if v == vs {
			r = append(r, i)
		}
	}
	return r
}

func IndexesFunc[E any](s []E, f func(E) bool) []int {
	var r []int
	for i, v := range s {
		if f(v) {
			r = append(r, i)
		}
	}
	return r
}
