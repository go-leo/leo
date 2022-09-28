package syncx

import (
	"errors"
	"sync"

	"github.com/spf13/cast"
	"golang.org/x/sync/singleflight"
)

// Deprecated: Do not use. use github.com/go-leo/syncx instead.
type OnceCreator[T any] struct {
	m            sync.Map
	sfg          singleflight.Group
	createObject func() T
}

// Deprecated: Do not use. use github.com/go-leo/syncx instead.
func NewOnceCreator[T any](createObject func() T) (*OnceCreator[T], error) {
	var cache *OnceCreator[T]
	if createObject == nil {
		return cache, errors.New("new function is nil")
	}
	cache = &OnceCreator[T]{
		m:            sync.Map{},
		sfg:          singleflight.Group{},
		createObject: createObject,
	}
	return cache, nil
}

// Deprecated: Do not use. use github.com/go-leo/syncx instead.
func (m *OnceCreator[T]) LoadOrCreate(key string) T {
	v, ok := m.m.Load(key)
	if ok {
		return v.(T)
	}
	v, _, _ = m.sfg.Do(cast.ToString(key), func() (any, error) {
		if v, ok := m.m.Load(key); ok {
			return v, nil
		}
		v := m.createObject()
		m.m.Store(key, v)
		return v, nil
	})
	return v.(T)
}
