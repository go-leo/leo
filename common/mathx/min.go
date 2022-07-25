package mathx

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Min[N constraints.Integer | constraints.Float](x, y N) N {
	return N(math.Min(float64(x), float64(y)))
}
