package stringx

// IsEmpty checks if a string is empty ("")
func IsEmpty[S ~string](s S) bool {
	return len(s) == 0
}

// IsNotEmpty Checks if a string is not empty ("")
func IsNotEmpty[S ~string](s S) bool {
	return !IsEmpty(s)
}

// IsAllEmpty Checks if all of the strings are empty ("")
func IsAllEmpty[S ~string](ss ...S) bool {
	for _, s := range ss {
		if IsNotEmpty(s) {
			return false
		}
	}
	return true
}

// IsAnyEmpty Checks if any of the strings are empty ("")
func IsAnyEmpty[S ~string](ss ...S) bool {
	for _, s := range ss {
		if IsEmpty(s) {
			return true
		}
	}
	return false
}
