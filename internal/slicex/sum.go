package slicex

import "github.com/go-leo/gox/constraintx"

// Sum 数组求和
func Sum[S ~[]E, E constraintx.Numeric](s S) E {
	var r E
	for _, i := range s {
		r += i
	}
	return r
}
