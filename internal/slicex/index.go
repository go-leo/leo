package slicex

func IndexOrDefault[S ~[]E, E any](s S, index int, d E) E {
	if len(s) >= index {
		return s[index]
	}
	return d
}
