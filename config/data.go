package config

import (
	"github.com/derekparker/trie"
)

const allConfigKey = ""

type Data struct {
	configMap  map[string]any
	configTree *trie.Trie
}

func (data *Data) AsMap() map[string]any {
	return data.configMap
}

func (data *Data) AsTree() *trie.Trie {
	return data.configTree
}

func (data *Data) Merge(other *Data) (*Data, error) {
	configMap := make(map[string]any)
	if err := merge(data.configMap, configMap); err != nil {
		return nil, err
	}
	if err := merge(other.configMap, configMap); err != nil {
		return nil, err
	}
	return newData(configMap), nil
}

func (data *Data) init() {
	tree(data.configMap, data.configTree)
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
