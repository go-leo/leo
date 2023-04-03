package decoder

import (
	"strings"

	"github.com/go-leo/gox/encodingx/envx"
	"golang.org/x/exp/slices"
)

type ENV struct{}

func (ENV) IsSupported(extension string) bool {
	if slices.Contains([]string{"env", ".env"}, strings.ToLower(extension)) {
		return true
	}
	return false
}

func (ENV) Decode(data []byte, m map[string]any) error {
	envMap := make(map[string]string)
	err := envx.Unmarshal(data, envMap)
	if err != nil {
		return err
	}
	for key, val := range envMap {
		m[key] = val
	}
	return nil
}
