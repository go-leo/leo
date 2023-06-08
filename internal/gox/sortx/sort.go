package sortx

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/slicex"
)

// Asc sorts a slice of any ordered type in ascending order.
func Asc[E constraints.Ordered](x []E) []E {
	slices.Sort(x)
	return x
}

// Desc sorts a slice of any ordered type in descending order.
func Desc[E constraints.Ordered](x []E) []E {
	return slicex.Reverse(x)
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

func BubbleSort[E constraints.Ordered](x []E) []E {
	for i := 0; i < len(x)-1; i++ {
		for j := 1; j < len(x)-i; j++ {
			if x[j] < x[j-1] {
				x[j], x[j-1] = x[j-1], x[j]
			}
		}
	}
	return x
}

func SelectSort[E constraints.Ordered](x []E) []E {
	for i := 0; i < len(x); i++ {
		min := i
		for j := i + 1; j < len(x); j++ {
			if x[min] > x[j] {
				min = j
			}
		}
		x[i], x[min] = x[min], x[i]
	}
	return x
}

func InsertSort[E constraints.Ordered](x []E) []E {
	for i := 0; i < len(x); i++ {
		for j := i - 1; j >= 0; j-- {
			if x[j+1] < x[j] {
				x[j+1], x[j] = x[j], x[j+1]
			} else {
				break
			}
		}
	}
	return x
}
