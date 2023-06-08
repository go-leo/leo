package decoder

import (
	"strings"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/tomlx"
	"golang.org/x/exp/slices"
)

type TOML struct{}

func (TOML) IsSupported(extension string) bool {
	return slices.Contains([]string{"toml", ".toml"}, strings.ToLower(extension))
}

func (TOML) Decode(data []byte, m map[string]any) error {
	return tomlx.Unmarshal(data, &m)
}
