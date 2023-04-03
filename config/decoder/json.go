package decoder

import (
	"strings"

	"github.com/go-leo/gox/encodingx/jsonx"
	"golang.org/x/exp/slices"
)

type JSON struct{}

func (JSON) IsSupported(extension string) bool {
	if slices.Contains([]string{"json", ".json"}, strings.ToLower(extension)) {
		return true
	}
	return false
}

func (JSON) Decode(data []byte, m map[string]any) error {
	return jsonx.Unmarshal(data, &m)
}
