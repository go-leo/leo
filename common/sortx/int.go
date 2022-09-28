package sortx

import "sort"

// Deprecated: Do not use. use github.com/go-leo/sortx instead.
func IntsDesc(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
}

// Deprecated: Do not use. use github.com/go-leo/sortx instead.
func IntsAsc(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}
