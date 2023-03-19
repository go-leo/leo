package errorx

func Must[T any](v T, err error) T {
	if err != nil {
		panic("must: " + err.Error())
	}
	return v
}
