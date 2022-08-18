package mapx

import (
	"sync"
)

type SyncMap[K comparable, V any] struct {
	m *sync.Map
}

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{m: new(sync.Map)}
}

func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	if !ok {
		return value, ok
	}
	value, _ = v.(V)
	return value, ok
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := m.m.LoadOrStore(key, value)
	actual, _ = v.(V)
	return actual, loaded
}

func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.m.LoadAndDelete(key)
	if !loaded {
		return value, loaded
	}
	value, _ = v.(V)
	return value, loaded
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *SyncMap[K, V]) DeleteFunc(del func(K, V) bool) {
	m.m.Range(func(key, value any) bool {
		k, _ := key.(K)
		v, _ := value.(V)
		if del(k, v) {
			m.m.Delete(key)
		}
		return true
	})
}

func (m *SyncMap[K, V]) Map() map[K]V {
	resMap := make(map[K]V)
	m.m.Range(func(key, value any) bool {
		k, _ := key.(K)
		v, _ := value.(V)
		resMap[k] = v
		return true
	})
	return resMap
}

func (m *SyncMap[K, V]) Len() int {
	length := 0
	m.m.Range(func(_, _ any) bool {
		length++
		return true
	})
	return length
}

func (m *SyncMap[K, V]) Keys() []K {
	var keys []K
	m.m.Range(func(key, _ any) bool {
		k, _ := key.(K)
		keys = append(keys, k)
		return true
	})
	return keys
}

func (m *SyncMap[K, V]) Values() []V {
	var values []V
	m.m.Range(func(_, value any) bool {
		v, _ := value.(V)
		values = append(values, v)
		return true
	})
	return values
}

func (m *SyncMap[K, V]) Clear() {
	m.m.Range(func(key, _ any) bool {
		m.m.Delete(key)
		return true
	})
}

func (m *SyncMap[K, V]) Clone() Map[K, V] {
	newMap := new(sync.Map)
	m.m.Range(func(key, value any) bool {
		newMap.Store(key, value)
		return true
	})
	return &SyncMap[K, V]{m: newMap}
}
