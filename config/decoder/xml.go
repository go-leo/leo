package decoder

import (
	"strings"

	"github.com/go-leo/gox/encodingx/xmlx"
	"golang.org/x/exp/slices"
)

type XMLDecoder struct{}

func (d *XMLDecoder) IsSupported(extension string) bool {
	if slices.Contains([]string{"xml"}, strings.ToLower(extension)) {
		return true
	}
	return false
}

func (d *XMLDecoder) Decode(bytes []byte, m map[string]any) error {
	return xmlx.Unmarshal(bytes, &m)
}
