package http

import (
	"golang.org/x/exp/maps"
	"net/http"
	"strings"
)

func (x *Status) HttpHeader() http.Header {
	header := make(http.Header, len(x.GetHeaders()))
	for _, h := range x.GetHeaders() {
		header.Add(h.GetKey(), h.GetValue())
	}
	keys := maps.Keys(header)
	header.Add("X-Leo-Status-Keys", strings.Join(keys, ", "))
	return header
}
