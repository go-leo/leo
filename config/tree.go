package config

import (
	"fmt"
	"github.com/derekparker/trie"
	"reflect"

	"github.com/spf13/cast"
	"strconv"
)

func tree(configMap map[string]any, tree *trie.Trie) {
	mapToTrie(allConfigKey, configMap, tree)
}

func mapToTrie(key string, config any, tree *trie.Trie) {
	tree.Add(key, config)

	var configValue reflect.Value
	if t := reflect.TypeOf(config); t.Kind() == reflect.Pointer {
		v := reflect.ValueOf(config)
		for v.Kind() == reflect.Pointer && !v.IsNil() {
			v = v.Elem()
		}
		configValue = v
	} else {
		configValue = reflect.ValueOf(config)
	}

	switch configValue.Kind() {
	case reflect.Map:
		keys := configValue.MapKeys()
		for _, keyValue := range keys {
			key := getKey(key, cast.ToString(keyValue.Interface()))
			mapToTrie(key, configValue.MapIndex(keyValue).Interface(), tree)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < configValue.Len(); i++ {
			key := getKey(key, strconv.Itoa(i))
			mapToTrie(key, configValue.Index(i).Interface(), tree)
		}
	}
}

func getKey(key string, subKey string) string {
	if key == allConfigKey {
		return subKey
	} else {
		return fmt.Sprintf("%s.%s", key, subKey)
	}
}
