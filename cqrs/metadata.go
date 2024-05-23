package cqrs

import (
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type metadata map[string][]string

func (m metadata) Set(key string, value ...string) {
	m[key] = value
}

func (m metadata) Append(key string, value ...string) {
	m[key] = append(m[key], value...)
}

func (m metadata) Get(key string) string {
	if m == nil {
		return ""
	}
	v := m[key]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func (m metadata) Values(key string) []string {
	if m == nil {
		return nil
	}
	return m[key]
}

func (m metadata) Keys() []string {
	return maps.Keys(m)
}

func (m metadata) Delete(key string) {
	delete(m, key)
}

func (m metadata) Len() int {
	return len(m)
}

func (m metadata) Clone() Metadata {
	clonedMd := metadata{}
	for _, key := range m.Keys() {
		clonedMd[key] = slices.Clone(m[key])
	}
	return clonedMd
}

func NewMetadata() Metadata {
	return metadata{}
}
