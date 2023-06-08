package slicex

func IsSameLength[S ~[]E, E any](s1 S, s2 S) bool {
	return len(s1) == len(s2)
}
