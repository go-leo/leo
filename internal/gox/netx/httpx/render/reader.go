// Copyright 2018 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package render

import (
	"io"
	"net/http"
	"strconv"
)

// Reader writes data with custom ContentType and headers.
func Reader(
	w http.ResponseWriter,
	reader io.Reader,
	contentLength int64,
	contentType string,
	headers map[string]string,
) error {
	writeContentType(w, []string{contentType})
	if contentLength >= 0 {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Content-Length"] = strconv.FormatInt(contentLength, 10)
	}
	header := w.Header()
	for k, v := range headers {
		if header.Get(k) == "" {
			header.Set(k, v)
		}
	}
	_, err := io.Copy(w, reader)
	return err
}
