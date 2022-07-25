package config

// Parser encode raw data to map `map[string]any`
type Parser interface {
	// Parse unmarshal raw data into a config that type of `map[string]any`
	// must ensure the sub value is type of `map[string]any`
	Parse(rawData []byte) error
	// ConfigMap return config that type of `map[string]any`
	ConfigMap() map[string]any
	// Support returns true if the Parser supports this contentType, return false if the Parser not supports this contentType
	Support(contentType string) bool
}
