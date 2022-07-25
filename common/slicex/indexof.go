package slicex

import "strings"

const (
	// IndexNotFound means that no elements were found
	IndexNotFound = -1
)

// IndexOfString Finds the index of the given value in the slice.
func IndexOfString(slice []string, value string, startIndex int) int {
	if IsEmpty(slice) {
		return IndexNotFound
	}
	if startIndex < 0 {
		startIndex = 0
	}
	length := len(slice)
	for i := startIndex; i < length; i++ {
		if value == slice[i] {
			return i
		}
	}
	return IndexNotFound
}

// PrefixIndexOfString Finds the index of the prefix of the given value in the prefix slice.
func PrefixIndexOfString(slice []string, value string, startIndex int) int {
	if IsEmpty(slice) {
		return IndexNotFound
	}
	if startIndex < 0 {
		startIndex = 0
	}
	length := len(slice)
	for i := startIndex; i < length; i++ {
		prefix := slice[i]
		if strings.HasPrefix(value, prefix) {
			return i
		}
	}
	return IndexNotFound
}
