package decoder

import (
	"gopkg.in/yaml.v3"
	"strings"

	"golang.org/x/exp/slices"
)

type YAML struct{}

func (YAML) IsSupported(extension string) bool {
	return slices.Contains([]string{"yaml", ".yaml", "yml", ".yml"}, strings.ToLower(extension))
}

func (YAML) Decode(data []byte, m map[string]any) error {
	return yaml.Unmarshal(data, &m)
}
