package config

import (
	"strings"

	"github.com/spf13/cast"
)

type cleaner struct{}

func (c cleaner) Clean(configMap map[string]any) {
	c.cleanMap(configMap)
}

func (c cleaner) cleanMap(configMap map[string]any) {
	for key, val := range configMap {
		val = c.cleanValue(val)
		lower := strings.ToLower(key)
		if key != lower {
			// remove old key (not lower-cased)
			delete(configMap, key)
		}
		// update map
		configMap[lower] = val
	}
}

func (c cleaner) cleanValue(val any) any {
	switch val.(type) {
	case map[any]any:
		// nested map: cast and recursively clean
		val = cast.ToStringMap(val)
		c.cleanMap(val.(map[string]any))
	case map[string]any:
		// nested map: recursively clean
		c.cleanMap(val.(map[string]any))
	case []any:
		// nested array: recursively clean
		c.cleanSlice(val.([]any))
	}
	return val
}

func (c cleaner) cleanSlice(a []any) {
	for i, val := range a {
		a[i] = c.cleanValue(val)
	}
}
