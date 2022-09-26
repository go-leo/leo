package iox

import "io"

// Deprecated: Do not use. use github.com/go-leo/iox instead.
func QuiteClose(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}
