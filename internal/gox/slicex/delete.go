package slicex

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func Delete[S ~[]E, E any](s S, i int) S {
	return slices.Delete(s, i, i+1)
}

func DeleteAll[S ~[]E, E any](array S, indices ...int) S {
	length := len(array)
	diff := 0 // number of distinct indexes, i.e. number of entries that will be removed
	slices.Sort(indices)

	// identify length of result array
	if IsNotEmpty(indices) {
		i := len(indices) - 1
		prevIndex := length
		for i >= 0 {
			index := indices[i]
			if index < 0 || index >= length {
				panic(fmt.Errorf("index: %d, length: %d", index, length))
			}
			if index >= prevIndex {
				continue
			}
			diff++
			prevIndex = index
			i--
		}
	}

	// create result array
	result := make(S, length-diff)
	if diff < length {
		end := length         // index just after last copy
		dest := length - diff // number of entries so far not copied
		for i := len(indices) - 1; i >= 0; i-- {
			index := indices[i]
			if end-index > 1 { // same as (cp > 0)
				cp := end - index - 1
				dest -= cp
				copy(result[dest:dest+cp], array[index+1:index+1+cp])
				// After this copy, we still have room for dest items.
			}
			end = index
		}
		if end > 0 {
			copy(result[:end], array[:end])
		}
	}
	return result
}
