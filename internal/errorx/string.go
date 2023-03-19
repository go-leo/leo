package errorx

func Stringfy(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
