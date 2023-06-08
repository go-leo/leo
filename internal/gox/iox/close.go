package iox

import "io"

func QuiteClose(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}
