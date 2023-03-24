package config

type Source struct {
	// Name of the config
	Name string
	//  Value from config
	Value []byte
	// Extension identifies the type of Value
	Extension string
}
