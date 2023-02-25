package config

import (
	"strings"

	"github.com/spf13/cast"
)

type Cleaner interface {
	Clean(configMap any) (map[string]any, error)
}

type cleaner struct{}

func newCleaner() Cleaner {
	return &cleaner{}
}

func (c cleaner) Clean(configMap any) (map[string]any, error) {

	return nil, nil
}

func insensitiviseMap(m map[string]any) {
	for key, val := range m {
		val = insensitiviseVal(val)
		lower := strings.ToLower(key)
		if key != lower {
			// remove old key (not lower-cased)
			delete(m, key)
		}
		// update map
		m[lower] = val
	}
}

func insensitiviseVal(val any) any {
	switch val.(type) {
	case map[any]any:
		// nested map: cast and recursively insensitivise
		val = cast.ToStringMap(val)
		insensitiviseMap(val.(map[string]any))
	case map[string]any:
		// nested map: recursively insensitivise
		insensitiviseMap(val.(map[string]any))
	case []any:
		// nested array: recursively insensitivise
		insensitiveArray(val.([]any))
	}
	return val
}

func insensitiveArray(a []any) {
	for i, val := range a {
		a[i] = insensitiviseVal(val)
	}
}
