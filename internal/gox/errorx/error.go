package errorx

func Must[T any](v T, err error) T {
	if err != nil {
		panic("must: " + err.Error())
	}
	return v
}

func Ignore[T any](v T, _ error) T {
	return v
}

func Concern[T any](_ T, err error) error {
	return err
}

func Quiet[T any](v T, err error) T {
	return Ignore(v, err)
}

func Silence(_ error) {}

func Stringfy(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
