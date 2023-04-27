package decoder

import (
	"strings"

	"github.com/go-leo/gox/encodingx/yamlx"
	"golang.org/x/exp/slices"
)

type YAML struct{}

func (YAML) IsSupported(extension string) bool {
	return slices.Contains([]string{"yaml", ".yaml", "yml", ".yml"}, strings.ToLower(extension))
}

func (YAML) Decode(data []byte, m map[string]any) error {
	return yamlx.Unmarshal(data, &m)
}
