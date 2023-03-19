package mapx

import (
	"sync"

	"golang.org/x/exp/maps"
)

type StdMap[K comparable, V any] struct {
	m map[K]V
	sync.RWMutex
}

func NewStdMap[K comparable, V any](initSize int) *StdMap[K, V] {
	return &StdMap[K, V]{m: make(map[K]V, initSize)}
}

func (m *StdMap[K, V]) Load(key K) (value V, ok bool) {
	m.RLock()
	v, ok := m.m[key]
	m.RUnlock()
	return v, ok
}

func (m *StdMap[K, V]) Store(key K, value V) {
	m.Lock()
	m.m[key] = value
	m.Unlock()
}

func (m *StdMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	m.Lock()
	actual, loaded = m.m[key]
	if !loaded {
		m.m[key] = value
		actual = value
	}
	m.Unlock()
	return actual, loaded
}

func (m *StdMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	m.Lock()
	value, loaded = m.m[key]
	if loaded {
		delete(m.m, key)
	}
	m.Unlock()
	return value, loaded
}

func (m *StdMap[K, V]) Delete(key K) {
	m.Lock()
	delete(m.m, key)
	m.Unlock()
}

func (m *StdMap[K, V]) DeleteFunc(del func(K, V) bool) {
	m.Lock()
	maps.DeleteFunc(m.m, del)
	m.Unlock()
}

func (m *StdMap[K, V]) Map() map[K]V {
	m.RLock()
	resMap := maps.Clone(m.m)
	m.RUnlock()
	return resMap
}

func (m *StdMap[K, V]) Len() int {
	m.RLock()
	length := len(m.m)
	m.RUnlock()
	return length
}

func (m *StdMap[K, V]) Keys() []K {
	m.RLock()
	keys := maps.Keys(m.m)
	m.RUnlock()
	return keys
}

func (m *StdMap[K, V]) Values() []V {
	m.RLock()
	values := maps.Values(m.m)
	m.RUnlock()
	return values
}

func (m *StdMap[K, V]) Clear() {
	m.Lock()
	maps.Clear(m.m)
	m.Unlock()
}

func (m *StdMap[K, V]) Clone() Map[K, V] {
	m.RLock()
	newMap := maps.Clone(m.m)
	m.RUnlock()
	return &StdMap[K, V]{m: newMap}
}
