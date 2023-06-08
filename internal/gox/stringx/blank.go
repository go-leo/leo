package stringx

import "strings"

// IsBlank Checks if a string is empty ("") or whitespace only.
func IsBlank[S ~string](s S) bool {
	return len(strings.TrimSpace(string(s))) == 0
}

// IsNotBlank Checks if a string is not empty ("") and not whitespace only.
func IsNotBlank[S ~string](s S) bool {
	return !IsBlank(s)
}

// IsAllBlank Checks if all of the CharSequences are empty ("") or whitespace only.
func IsAllBlank[S ~string](ss ...S) bool {
	for _, s := range ss {
		if IsNotBlank(s) {
			return false
		}
	}
	return true
}

// IsAnyBlank Checks if any of the string are empty ("") or whitespace only.
func IsAnyBlank[S ~string](ss ...S) bool {
	for _, s := range ss {
		if IsBlank(s) {
			return true
		}
	}
	return false
}
