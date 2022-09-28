package errorx

// Deprecated: Do not use. use github.com/go-leo/errorx instead.
func Must[T any](v T, err error) T {
	if err != nil {
		panic("must: " + err.Error())
	}
	return v
}
