package valuer

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/derekparker/trie"
	"github.com/spf13/cast"

	"github.com/hmldd/leo/config"
)

var _ config.Valuer = new(TrieValuer)

type TrieValuer struct {
	config map[string]any
	tree   *trie.Trie
}

func NewTrieTreeValuer() *TrieValuer {
	return &TrieValuer{
		config: make(map[string]any),
		tree:   trie.New(),
	}
}

func (v *TrieValuer) AddConfig(configs ...map[string]any) {
	for _, config := range configs {
		v.mergeMaps(config, v.config)
	}
	v.addToTree(v.config, v.tree)
}

func (v *TrieValuer) Config() map[string]any {
	return v.config
}

func (v *TrieValuer) Get(key string) (any, error) {
	node, ok := v.tree.Find(key)
	if !ok {
		return nil, fmt.Errorf("not found %s value", key)
	}
	return node.Meta(), nil
}

func (v *TrieValuer) addToTree(config map[string]any, tree *trie.Trie) {
	v.mapToTrie("", config, tree)
}

func (v *TrieValuer) mapToTrie(keyPrefix string, config any, tree *trie.Trie) {
	tree.Add(keyPrefix, config)
	switch m := config.(type) {
	case map[any]any:
		for subKey, conf := range m {
			key := v.getKey(keyPrefix, cast.ToString(subKey))
			v.mapToTrie(key, conf, tree)
		}
	case map[string]any:
		for subKey, conf := range m {
			key := v.getKey(keyPrefix, subKey)
			v.mapToTrie(key, conf, tree)
		}
	case []any:
		for i, val := range m {
			key := v.getKey(keyPrefix, strconv.Itoa(i))
			v.mapToTrie(key, val, tree)
		}
	}
}

func (v *TrieValuer) getKey(keyPrefix string, subKey string) string {
	var key string
	if keyPrefix != "" {
		key = fmt.Sprintf("%s.%s", keyPrefix, subKey)
	} else {
		key = subKey
	}
	return key
}

// mergeMaps merges two string maps.
func (v *TrieValuer) mergeMaps(src, target map[string]any) {
	for key, srcVal := range src {
		// if target not contains srcKey, add srcVal and continue
		targetVal, ok := target[key]
		if !ok {
			target[key] = srcVal
			continue
		}

		// if targetVal and srcVal is not same type, ignore this val, and log it
		srcValType := reflect.TypeOf(srcVal)
		targetValType := reflect.TypeOf(targetVal)
		if targetValType != nil && srcValType != targetValType {
			fmt.Printf(
				"merge map, srcValType != targetValType; key=%s, srcValType=%v, targetValType=%v, srcVal=%v, targetVal=%v",
				key, srcValType, targetValType, srcVal, targetVal)
			continue
		}

		// if targetVal is `map[string]any`, merge sub string map.
		// else add srcVal and continue
		switch subTargetVal := targetVal.(type) {
		case map[string]any:
			subSrvVal := srcVal.(map[string]any)
			v.mergeMaps(subSrvVal, subTargetVal)
		default:
			target[key] = srcVal
			continue
		}
	}
}
