package parser

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/hmldd/leo/config"
)

var _ config.Parser = new(YamlParser)

type YamlParser struct {
	config map[string]any
}

func (p *YamlParser) Parse(rawData []byte) error {
	configMap := make(map[string]any)
	err := yaml.Unmarshal(rawData, &configMap)
	if err != nil {
		return fmt.Errorf("failed parse yaml config, %w", err)
	}
	p.config, _ = standardized(configMap).(map[string]any)
	return nil
}

func (p *YamlParser) ConfigMap() map[string]any {
	return p.config
}

func (p *YamlParser) Support(contentType string) bool {
	return YAML == contentType || YML == contentType
}

func NewYamlParser() config.Parser {
	return &YamlParser{}
}
