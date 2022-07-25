package slicex

import (
	"fmt"
)

const indexOutOfBoundsErrorFormat = "index out of bounds, Index: %d, Length: %d"

// RemoveString removes the element at the specified position from the specified array.
// All subsequent elements are shifted to the left (subtracts one from their indices).
func RemoveString(slice []string, index int) ([]string, error) {
	if index < 0 || len(slice) <= index {
		return nil, fmt.Errorf(indexOutOfBoundsErrorFormat, index, len(slice))
	}
	dst := make([]string, len(slice)-1)
	copy(dst, slice[:index])
	copy(dst[index:], slice[index+1:])
	return dst, nil
}
