package config

type Source struct {
	// Name of the config
	name string
	//  Value from config
	value []byte
	// Extension identifies the type of Value
	extension string
}

func (s *Source) Name() string {
	return s.name
}

func (s *Source) Value() []byte {
	return s.value
}

func (s *Source) Extension() string {
	return s.extension
}

func NewSource(name string, value []byte, extension string) *Source {
	return &Source{name: name, value: value, extension: extension}
}
