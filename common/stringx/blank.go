package stringx

import "strings"

// IsBlank Checks if a string is empty ("") or whitespace only.
// Deprecated: Do not use. use github.com/go-leo/stringx instead.
func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsNotBlank Checks if a string is not empty ("") and not whitespace only.
// Deprecated: Do not use. use github.com/go-leo/stringx instead.
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// IsAllBlank Checks if all of the CharSequences are empty ("") or whitespace only.
// Deprecated: Do not use. use github.com/go-leo/stringx instead.
func IsAllBlank(ss ...string) bool {
	for _, s := range ss {
		if IsNotBlank(s) {
			return false
		}
	}
	return true
}

// IsAnyBlank Checks if any of the string are empty ("") or whitespace only.
// Deprecated: Do not use. use github.com/go-leo/stringx instead.
func IsAnyBlank(ss ...string) bool {
	for _, s := range ss {
		if IsBlank(s) {
			return true
		}
	}
	return false
}
