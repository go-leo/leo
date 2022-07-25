package errorx

func String(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
