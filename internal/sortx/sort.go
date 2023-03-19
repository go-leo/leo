package sortx

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Asc sorts a slice of any ordered type in ascending order.
func Asc[E constraints.Ordered](x []E) {
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})
}

// Desc sorts a slice of any ordered type in descending order.
func Desc[E constraints.Ordered](x []E) {
	sort.Slice(x, func(i, j int) bool {
		return x[i] > x[j]
	})
}

// IsAsc reports whether x is sorted in ascending order.
func IsAsc[E constraints.Ordered](x []E) bool {
	for i := len(x) - 1; i > 0; i-- {
		if x[i] < x[i-1] {
			return false
		}
	}
	return true
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
