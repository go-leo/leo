package errorx

func Quiet[T any](v T, _ error) T {
	return v
}

func Silence(_ error) {
	return
}
