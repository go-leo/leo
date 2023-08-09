package decoder

import (
	"errors"
	"github.com/joho/godotenv"
	"golang.org/x/exp/maps"
	"strings"

	"golang.org/x/exp/slices"
)

type ENV struct{}

func (ENV) IsSupported(extension string) bool {
	return slices.Contains([]string{"env", ".env"}, strings.ToLower(extension))
}

func (ENV) Decode(data []byte, m map[string]any) error {
	envMap := make(map[string]string)
	err := Unmarshal(data, envMap)
	if err != nil {
		return err
	}
	for key, val := range envMap {
		m[key] = val
	}
	return nil
}

func Unmarshal(data []byte, val any) error {
	envMap, ok := val.(map[string]string)
	if !ok {
		return errors.New("any not convert to map[string]string")
	}
	m, err := godotenv.UnmarshalBytes(data)
	if err != nil {
		return err
	}
	maps.Copy(envMap, m)
	return nil
}
