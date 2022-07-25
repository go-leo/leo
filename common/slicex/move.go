package slicex

import (
	"fmt"
)

// Move moves an array element from index to target
func Move(slice []any, index int, target int) ([]any, error) {
	if index < 0 || len(slice) <= index {
		return nil, fmt.Errorf(indexOutOfBoundsErrorFormat, index, len(slice))
	}
	if target < 0 || len(slice) <= target {
		return nil, fmt.Errorf("target out of bounds, Target: %d, Length: %d", target, len(slice))
	}
	if index == target {
		dst := make([]any, len(slice))
		copy(dst, slice)
		return dst, nil
	}
	dst := make([]any, len(slice))
	copy(dst, slice[:index])
	copy(dst[index:], slice[index+1:])
	copy(dst[target+1:], dst[target:len(slice)-1])
	dst[target] = slice[index]
	return dst, nil
}
