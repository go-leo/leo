package config

import (
	"fmt"
)

// Decoder decodes the config data into config map.
type Decoder interface {
	// IsSupported return true if the extension that the Decoder support (excluding the '.').
	IsSupported(extension string) bool
	// Decode bytes to map
	Decode([]byte, map[string]any) error
}

type parser struct {
	Decoders []Decoder
}

func (p *parser) Parse(s *Source) (*Data, error) {
	for _, decoder := range p.Decoders {
		if !decoder.IsSupported(s.Extension) {
			continue
		}
		configMap := make(map[string]any)
		err := decoder.Decode(s.Value, configMap)
		if err != nil {
			return nil, err
		}
		d := newData(configMap)
		return d, nil
	}
	return nil, fmt.Errorf("unknown extension %s", s.Extension)
}
