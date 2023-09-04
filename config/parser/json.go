package parser

import (
	"encoding/json"
	"fmt"

	"github.com/hmldd/leo/config"
)

var _ config.Parser = new(JsonParser)

type JsonParser struct {
	config map[string]any
}

func (p *JsonParser) Parse(rawData []byte) error {
	configMap := make(map[string]any)
	err := json.Unmarshal(rawData, &configMap)
	if err != nil {
		return fmt.Errorf("failed parse json config, %w", err)
	}
	p.config, _ = standardized(configMap).(map[string]any)
	return nil
}

func (p *JsonParser) ConfigMap() map[string]any {
	return p.config
}

func (p *JsonParser) Support(contentType string) bool {
	return JSON == contentType
}

func NewJsonParser() *JsonParser {
	return &JsonParser{}
}
