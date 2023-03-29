package decoder

import (
	"strings"

	"github.com/go-leo/gox/encodingx/yamlx"
	"golang.org/x/exp/slices"
)

type YAML struct{}

func (YAML) IsSupported(extension string) bool {
	if slices.Contains([]string{"yaml", ".yaml", "yml", ".yml"}, strings.ToLower(extension)) {
		return true
	}
	return false
}

func (YAML) Decode(bytes []byte, m map[string]any) error {
	return yamlx.Unmarshal(bytes, &m)
}
