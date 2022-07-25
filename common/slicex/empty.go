package slicex

// IsEmptyString Checks if an slice is nil or length equals 0
func IsEmptyString(slice []string) bool {
	return len(slice) <= 0
}
