package decoder

import (
	"strings"

	"github.com/go-leo/gox/encodingx/jsonx"
	"golang.org/x/exp/slices"
)

type JSONDecoder struct{}

func (d *JSONDecoder) IsSupported(extension string) bool {
	if slices.Contains([]string{"json"}, strings.ToLower(extension)) {
		return true
	}
	return false
}

func (d *JSONDecoder) Decode(bytes []byte, m map[string]any) error {
	return jsonx.Unmarshal(bytes, &m)
}
