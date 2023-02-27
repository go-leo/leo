package config

// Loader is a loader that can be used to load config data.
type Loader interface {
	Load() ([]byte, error)
}

type noopLoader struct{}

func (noopLoader) Load() ([]byte, error) {
	return []byte{}, nil
}
