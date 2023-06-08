package slicex

import "golang.org/x/exp/slices"

// ContainsAny checks if any of the elem are in the given slice.
func ContainsAny[E comparable](s []E, vs ...E) bool {
	for _, v := range vs {
		if slices.Contains(s, v) {
			return true
		}
	}
	return false
}
