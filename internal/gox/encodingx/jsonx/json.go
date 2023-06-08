//go:build !jsoniter && !go_json
// +build !jsoniter,!go_json

package jsonx

import (
	"encoding/json"
	"io"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func NewEncoder(w io.Writer) JSONEncoder {
	return json.NewEncoder(w)
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return json.NewDecoder(r)
}
