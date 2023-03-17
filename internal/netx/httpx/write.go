package httpx

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func WriteJSON[T any](writer http.ResponseWriter, obj T) error {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(obj)
	if err != nil {
		return err
	}
	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}
