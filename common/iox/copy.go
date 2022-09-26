package iox

import "io"

// Deprecated: Do not use. use github.com/go-leo/iox instead.
func Copy(dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	return err
}
