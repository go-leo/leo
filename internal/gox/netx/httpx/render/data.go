package render

import "net/http"

// Data writes data with custom ContentType.
func Data(w http.ResponseWriter, data []byte, contentType string) error {
	writeContentType(w, []string{contentType})
	_, err := w.Write(data)
	return err
}
