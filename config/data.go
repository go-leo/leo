package config

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-leo/gox/containerx/trie"
	"github.com/go-leo/gox/convx"
)

const allConfigKey = ""

type Data struct {
	configMap  map[string]any
	configTree *trie.Trie
}

func (d *Data) AsMap() map[string]any {
	return d.configMap
}

func (d *Data) AsTree() *trie.Trie {
	return d.configTree
}

func (d *Data) Merge(o *Data) (*Data, error) {
	configMap := make(map[string]any)
	if err := d.merge(d.configMap, configMap); err != nil {
		return nil, err
	}
	if err := d.merge(o.configMap, configMap); err != nil {
		return nil, err
	}
	return newData(configMap), nil
}

func (d *Data) init() {
	d.clean()
	d.tree()
}

func (d *Data) clean() {
	d.cleanMap(d.configMap)
}

func (d *Data) cleanMap(configMap map[string]any) {
	for key, val := range configMap {
		val = d.cleanValue(val)
		configMap[key] = val
	}
}

func (d *Data) cleanValue(val any) any {
	switch realVal := val.(type) {
	case map[any]any:
		// nested map: cast and recursively clean
		d.cleanMap(convx.ToStringMap(val))
	case map[string]any:
		// nested map: recursively clean
		d.cleanMap(realVal)
	case []any:
		// nested array: recursively clean
		d.cleanSlice(realVal)
	}
	return val
}

func (d *Data) cleanSlice(a []any) {
	for i, val := range a {
		a[i] = d.cleanValue(val)
	}
}

func (d *Data) tree() {
	d.mapToTrie("", d.configMap, d.configTree)
}

func (d *Data) mapToTrie(key string, config any, tree *trie.Trie) {
	tree.Add(key, config)
	switch m := config.(type) {
	case map[any]any:
		for subKey, conf := range m {
			key := d.getKey(key, convx.ToString(subKey))
			d.mapToTrie(key, conf, tree)
		}
	case map[string]any:
		for subKey, conf := range m {
			key := d.getKey(key, subKey)
			d.mapToTrie(key, conf, tree)
		}
	case []any:
		for i, val := range m {
			key := d.getKey(key, strconv.Itoa(i))
			d.mapToTrie(key, val, tree)
		}
	}
}

func (d *Data) getKey(key string, subKey string) string {
	if key == allConfigKey {
		return subKey
	} else {
		return fmt.Sprintf("%s.%s", key, subKey)
	}
}

func (d *Data) merge(src, tgt map[string]any) error {
	return d.mergeMaps(src, tgt, nil)
}

func (d *Data) mergeMaps(src, tgt map[string]any, itgt map[any]any) error {
	for srcKey, srcVal := range src {
		// key ignore case, not exist, set
		tgtKey, ok := d.keyExists(srcKey, tgt)
		if !ok {
			d.setToTarget(srcKey, srcVal, tgt, itgt)
			continue
		}

		// care about case, not exist, set
		tgtVal, ok := tgt[tgtKey]
		if !ok {
			d.setToTarget(srcKey, srcVal, tgt, itgt)
			continue
		}

		srcValType := reflect.TypeOf(srcVal)
		tgtValType := reflect.TypeOf(tgtVal)
		switch tmpTgtVal := tgtVal.(type) {
		case map[any]any:
			tmpSrcVal, ok := srcVal.(map[any]any)
			if !ok {
				return fmt.Errorf(`failed to cast src value to map[any]any, srcKey: %s, srcValType: %s, tgtValType: %s, srcVal: %s, tgtVal: %s`, srcKey, srcValType, tgtValType, srcVal, tgtVal)
			}
			err := d.mergeMaps(d.anyMapToStringMap(tmpSrcVal), d.anyMapToStringMap(tmpTgtVal), tmpTgtVal)
			if err != nil {
				return err
			}
		case map[string]any:
			tmpSrcVal, ok := srcVal.(map[string]any)
			if !ok {
				return fmt.Errorf(`failed to cast src value to map[string]any, srcKey: %s, srcValType: %s, tgtValType: %s, srcVal: %s, tgtVal: %s`, srcKey, srcValType, tgtValType, srcVal, tgtVal)
			}
			err := d.mergeMaps(tmpSrcVal, tmpTgtVal, nil)
			if err != nil {
				return err
			}
		default:
			tgt[tgtKey] = srcVal
			if itgt != nil {
				itgt[tgtKey] = srcVal
			}
			d.setToTarget(tgtKey, srcVal, tgt, itgt)
		}
	}

	return nil
}

func (d *Data) setToTarget(key string, val any, tgt map[string]any, itgt map[any]any) {
	tgt[key] = val
	if itgt != nil {
		itgt[key] = val
	}
}

func (d *Data) anyMapToStringMap(src map[any]any) map[string]any {
	tgt := map[string]any{}
	for k, v := range src {
		tgt[fmt.Sprintf("%v", k)] = v
	}
	return tgt
}

func (d *Data) keyExists(key string, tgt map[string]any) (string, bool) {
	lowerKey := strings.ToLower(key)
	for configKey := range tgt {
		if strings.ToLower(configKey) == lowerKey {
			return configKey, true
		}
	}
	return "", false
}

func newData(configMap map[string]any) *Data {
	d := &Data{configMap: configMap, configTree: trie.New()}
	d.init()
	return d
}

func mutilData(ds ...*Data) (*Data, error) {
	md := newData(map[string]any{})
	for _, d := range ds {
		nd, err := md.Merge(d)
		if err != nil {
			return nil, err
		}
		md = nd
	}
	return md, nil
}
