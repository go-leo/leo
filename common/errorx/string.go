package errorx

// Deprecated: Do not use. use github.com/go-leo/errorx instead.
func String(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
