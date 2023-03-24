package decoder

import (
	"strings"

	"github.com/go-leo/gox/encodingx/yamlx"
	"golang.org/x/exp/slices"
)

type YAMLDecoder struct{}

func (d *YAMLDecoder) IsSupported(extension string) bool {
	if slices.Contains([]string{"yaml", "yml"}, strings.ToLower(extension)) {
		return true
	}
	return false
}

func (d *YAMLDecoder) Decode(bytes []byte, m map[string]any) error {
	return yamlx.Unmarshal(bytes, &m)
}
