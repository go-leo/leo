package mathx

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Max[N constraints.Integer | constraints.Float](x, y N) N {
	return N(math.Max(float64(x), float64(y)))
}
