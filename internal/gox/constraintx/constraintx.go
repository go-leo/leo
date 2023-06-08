package constraintx

import "golang.org/x/exp/constraints"

type Numeric interface {
	constraints.Integer | constraints.Float
}
