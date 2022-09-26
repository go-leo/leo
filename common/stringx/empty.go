package stringx

// IsEmpty checks if a string is empty ("")
// Deprecated: Do not use. use github.com/go-leo/stringx instead.
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsNotEmpty Checks if a string is not empty ("")
// Deprecated: Do not use. use github.com/go-leo/stringx instead.
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// IsAllEmpty Checks if all of the strings are empty ("")
// Deprecated: Do not use. use github.com/go-leo/stringx instead.
func IsAllEmpty(ss ...string) bool {
	for _, s := range ss {
		if IsNotEmpty(s) {
			return false
		}
	}
	return true
}

// IsAnyEmpty Checks if any of the strings are empty ("")
// Deprecated: Do not use. use github.com/go-leo/stringx instead.
func IsAnyEmpty(ss ...string) bool {
	for _, s := range ss {
		if IsEmpty(s) {
			return true
		}
	}
	return false
}
