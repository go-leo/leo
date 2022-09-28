package mathx

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Deprecated: Do not use. use github.com/go-leo/mathx instead.
func Max[N constraints.Integer | constraints.Float](x, y N) N {
	return N(math.Max(float64(x), float64(y)))
}
