package decoder

import (
	"strings"

	"github.com/go-leo/gox/encodingx/tomlx"
	"golang.org/x/exp/slices"
)

type TOMLDecoder struct{}

func (d *TOMLDecoder) IsSupported(extension string) bool {
	if slices.Contains([]string{"toml"}, strings.ToLower(extension)) {
		return true
	}
	return false
}

func (d *TOMLDecoder) Decode(bytes []byte, m map[string]any) error {
	return tomlx.Unmarshal(bytes, &m)
}
