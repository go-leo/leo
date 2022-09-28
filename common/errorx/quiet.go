package errorx

// Deprecated: Do not use. use github.com/go-leo/errorx instead.
func Quiet[T any](v T, _ error) T {
	return v
}

// Deprecated: Do not use. use github.com/go-leo/errorx instead.
func Silence(_ error) {}
