package metadatax

import "reflect"

// Metadata is a mapping from metadata keys to values.
type Metadata interface {
	// Set sets the value of a given key with a slice of values.
	Set(key string, value ...string)

	// Append adds the values to key, not overwriting what was already stored at that key.
	Append(key string, value ...string)

	// Get gets the first value associated with the given key.
	Get(key string) string

	// Values returns all values associated with the given key.
	Values(key string) []string

	// Keys returns the keys of the Metadata.
	Keys() []string

	// Range calls f sequentially for each key and value present in the Metadata.
	Range(f func(key string, values []string) bool)

	// Delete removes the values for a given key.
	Delete(key string)

	// Len returns the number of items in Metadata.
	Len() int

	// Clone returns a copy of Metadata or nil if Metadata is nil.
	Clone() Metadata
}

var Type = reflect.TypeOf((*Metadata)(nil)).Elem()
