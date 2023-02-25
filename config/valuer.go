package config

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/derekparker/trie"
	"github.com/go-leo/gox/stringx"
	"github.com/spf13/cast"
)

var ValueNotFound = errors.New("value not found")

// Valuer is Gets the value of a key
type Valuer interface {
	Merge(configMaps ...map[string]any) error
	// Value can retrieve any value given the key to use.
	Value(key string) (*Value, error)
}

func NewValuer() Valuer {
	return &valuer{configMap: make(map[string]any), tree: trie.New()}
}

type valuer struct {
	configMap map[string]any
	tree      *trie.Trie
}

func (valuer *valuer) Merge(configMaps ...map[string]any) error {
	for _, configMap := range configMaps {
		err := mergeMaps(configMap, valuer.configMap, nil)
		if err != nil {
			return err
		}
	}
	rootKeyPrefix := ""
	mapToTrie(rootKeyPrefix, valuer.configMap, valuer.tree)
	return nil
}

func (valuer *valuer) Value(key string) (*Value, error) {
	if stringx.IsBlank(key) {
		return &Value{val: valuer.configMap}, nil
	}
	node, ok := valuer.tree.Find(key)
	if !ok {
		return nil, ValueNotFound
	}
	return &Value{val: node.Meta()}, nil
}

// mergeMaps merges two maps. The `itgt` parameter is for handling go-yaml's
// insistence on parsing nested structures as `map[any]any`
// instead of using a `string` as the key for nest structures beyond one level
// deep. Both map types are supported as there is a go-yaml fork that uses
// `map[string]any` instead.
func mergeMaps(src, tgt map[string]any, itgt map[any]any) error {
	for sk, sv := range src {
		tk := keyExists(sk, tgt)
		if tk == "" {
			tgt[sk] = sv
			if itgt != nil {
				itgt[sk] = sv
			}
			continue
		}

		tv, ok := tgt[tk]
		if !ok {
			tgt[sk] = sv
			if itgt != nil {
				itgt[sk] = sv
			}
			continue
		}

		svType := reflect.TypeOf(sv)
		tvType := reflect.TypeOf(tv)

		switch ttv := tv.(type) {
		case map[any]any:
			tsv, ok := sv.(map[any]any)
			if !ok {
				return fmt.Errorf(`could not cast sv to map[any]any, key: %v, st: %v, tt: %v, sv: %v, tv: %v`, sk, svType, tvType, sv, tv)
			}

			ssv := castToMapStringInterface(tsv)
			stv := castToMapStringInterface(ttv)
			err := mergeMaps(ssv, stv, ttv)
			if err != nil {
				return err
			}
		case map[string]any:
			tsv, ok := sv.(map[string]any)
			if !ok {
				return fmt.Errorf(`could not cast sv to map[string]any, key: %v, st: %v, tt: %v, sv: %v, tv: %v`, sk, svType, tvType, sv, tv)
			}
			err := mergeMaps(tsv, ttv, nil)
			if err != nil {
				return err
			}
		default:
			tgt[tk] = sv
			if itgt != nil {
				itgt[tk] = sv
			}
		}
	}
	return nil
}

func keyExists(k string, m map[string]any) string {
	lk := strings.ToLower(k)
	for mk := range m {
		lmk := strings.ToLower(mk)
		if lmk == lk {
			return mk
		}
	}
	return ""
}

func castToMapStringInterface(src map[any]any) map[string]any {
	tgt := map[string]any{}
	for k, v := range src {
		tgt[fmt.Sprintf("%v", k)] = v
	}
	return tgt
}

func mapToTrie(keyPrefix string, config any, tree *trie.Trie) {
	tree.Add(keyPrefix, config)
	switch m := config.(type) {
	case map[any]any:
		for subKey, conf := range m {
			key := getKey(keyPrefix, cast.ToString(subKey))
			mapToTrie(key, conf, tree)
		}
	case map[string]any:
		for subKey, conf := range m {
			key := getKey(keyPrefix, subKey)
			mapToTrie(key, conf, tree)
		}
	case []any:
		for i, val := range m {
			key := getKey(keyPrefix, strconv.Itoa(i))
			mapToTrie(key, val, tree)
		}
	}
}

func getKey(keyPrefix string, subKey string) string {
	var key string
	if keyPrefix != "" {
		key = fmt.Sprintf("%s.%s", keyPrefix, subKey)
	} else {
		key = subKey
	}
	return key
}
