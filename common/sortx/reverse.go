package sortx

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Deprecated: Do not use. use github.com/go-leo/sortx instead.
func Reverse[E constraints.Ordered](x []E, start, end int) {
	if len(x) == 0 {
		return
	}
	a := int(math.Max(float64(start), 0))
	b := int(math.Min(float64(len(x)), float64(end)) - 1)
	for b > a {
		x[a], x[b] = x[b], x[a]
		b--
		a++
	}
}
