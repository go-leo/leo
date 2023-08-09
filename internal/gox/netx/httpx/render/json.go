package render

import (
	"encoding/json"
	"net/http"
)

// JSON marshals the given interface object and writes it with custom ContentType.
func JSON(w http.ResponseWriter, data any) (err error) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(data)
}
