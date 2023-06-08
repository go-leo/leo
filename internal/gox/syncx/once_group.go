package syncx

import (
	"sync"
)

type OnceGroup struct {
	m sync.Map
}

func (o *OnceGroup) Do(key string, f func()) {
	actual, loaded := o.m.LoadOrStore(key, &sync.Once{})
	if !loaded {
		return
	}
	once := actual.(*sync.Once)
	once.Do(f)
}
