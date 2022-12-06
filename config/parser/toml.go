package parser

import (
	"fmt"

	"github.com/BurntSushi/toml"

	"github.com/go-leo/leo/v2/config"
)

var _ config.Parser = new(TomlParser)

type TomlParser struct {
	config map[string]any
}

func (p *TomlParser) Parse(rawData []byte) error {
	configMap := make(map[string]any)
	err := toml.Unmarshal(rawData, &configMap)
	if err != nil {
		return fmt.Errorf("failed parse toml config, %w", err)
	}
	p.config, _ = standardized(configMap).(map[string]any)
	return nil
}

func (p *TomlParser) ConfigMap() map[string]any {
	return p.config
}

func (p *TomlParser) Support(contentType string) bool {
	return TOML == contentType
}

func NewTomlParser() *TomlParser {
	return &TomlParser{}
}
