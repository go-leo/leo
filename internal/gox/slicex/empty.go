package slicex

// IsEmpty Checks if an slice is nil or length equals 0
func IsEmpty[S ~[]E, E any](s S) bool {
	return len(s) <= 0
}

func IsNotEmpty[S ~[]E, E any](s S) bool {
	return len(s) > 0
}
