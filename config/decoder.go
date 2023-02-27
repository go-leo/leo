package config

// Decoder decodes the config data into config map.
type Decoder interface {
	Decode(data []byte, configMap map[string]any) error
}

type noopDecoder struct{}

func (noopDecoder) Decode(data []byte, configMap map[string]any) error {
	return nil
}
