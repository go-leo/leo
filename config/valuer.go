package config

// Valuer is Gets the value of a key
type Valuer interface {
	// AddConfig add and merge config
	AddConfig(configs ...map[string]any)
	// Config return merged config.
	Config() map[string]any
	// Get can retrieve any value given the key to use.
	Get(key string) (any, error)
}
