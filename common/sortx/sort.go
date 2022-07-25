package sortx

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Asc sorts a slice of any ordered type in ascending order.
func Asc[E constraints.Ordered](x []E) {
	slices.Sort(x)
}

// Desc sorts a slice of any ordered type in descending order.
func Desc[E constraints.Ordered](x []E) {
	slices.Sort(x)
	Reverse(x, 0, len(x))
}

// IsAsc reports whether x is sorted in ascending order.
func IsAsc[E constraints.Ordered](x []E) bool {
	return slices.IsSorted(x)
}

// IsDesc reports whether x is sorted in descending order.
func IsDesc[E constraints.Ordered](x []E) bool {
	for i := len(x) - 1; i > 0; i-- {
		if x[i] > x[i-1] {
			return false
		}
	}
	return true
}
