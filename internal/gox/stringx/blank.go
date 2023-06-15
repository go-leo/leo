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
