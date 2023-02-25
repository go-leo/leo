package config

// Decoder decodes the config data into config map.
type Decoder interface {
	Decode(data []byte, configMap map[string]interface{}) error
}
