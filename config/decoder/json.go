package decoder

import (
	"strings"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/jsonx"
	"golang.org/x/exp/slices"
)

type JSON struct{}

func (JSON) IsSupported(extension string) bool {
	return slices.Contains([]string{"json", ".json"}, strings.ToLower(extension))
}

func (JSON) Decode(data []byte, m map[string]any) error {
	return jsonx.Unmarshal(data, &m)
}
