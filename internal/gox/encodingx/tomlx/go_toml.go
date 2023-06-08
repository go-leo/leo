//go:build !toml
// +build !toml

package tomlx

import (
	"io"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx"
)

func Marshal(v any) ([]byte, error) {
	return toml.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return toml.Unmarshal(data, v)
}

func NewEncoder(w io.Writer) encodingx.Encoder {
	return toml.NewEncoder(w)
}

func NewDecoder(r io.Reader) encodingx.Decoder {
	return toml.NewDecoder(r)
}
