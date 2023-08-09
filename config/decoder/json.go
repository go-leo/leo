package decoder

import (
	"encoding/json"
	"golang.org/x/exp/slices"
	"strings"
)

type JSON struct{}

func (JSON) IsSupported(extension string) bool {
	return slices.Contains([]string{"json", ".json"}, strings.ToLower(extension))
}

func (JSON) Decode(data []byte, m map[string]any) error {
	return json.Unmarshal(data, &m)
}
