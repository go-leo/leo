package config

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/derekparker/trie"
	"github.com/go-leo/gox/stringx"
	"github.com/spf13/cast"
)

var ValueNotFound = errors.New("value not found")

var Nil = errors.New("value is nil")

const allConfigKey = ""

// Valuer is Gets the value of a key
type Valuer struct {
	configMap map[string]any
	tree      *trie.Trie
	cleaner   cleaner
	merger    merger
}

func newValuer() *Valuer {
	return &Valuer{configMap: make(map[string]any), tree: trie.New()}
}

func (valuer *Valuer) Merge(configMaps ...map[string]any) error {
	for _, configMap := range configMaps {
		valuer.cleaner.Clean(configMap)
		err := valuer.merger.Merge(configMap, valuer.configMap)
		if err != nil {
			return err
		}
	}
	mapToTrie(allConfigKey, valuer.configMap, valuer.tree)
	return nil
}

func (valuer *Valuer) Value(key string) (*Value, error) {
	if stringx.IsBlank(key) {
		return &Value{val: valuer.configMap}, nil
	}
	node, ok := valuer.tree.Find(key)
	if !ok {
		return nil, ValueNotFound
	}
	meta := node.Meta()
	if meta == nil {
		return nil, Nil
	}
	return &Value{val: meta}, nil
}

func mapToTrie(key string, config any, tree *trie.Trie) {
	tree.Add(key, config)
	switch m := config.(type) {
	case map[any]any:
		for subKey, conf := range m {
			key := getKey(key, cast.ToString(subKey))
			mapToTrie(key, conf, tree)
		}
	case map[string]any:
		for subKey, conf := range m {
			key := getKey(key, subKey)
			mapToTrie(key, conf, tree)
		}
	case []any:
		for i, val := range m {
			key := getKey(key, strconv.Itoa(i))
			mapToTrie(key, val, tree)
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
