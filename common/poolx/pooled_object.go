package poolx

import "sync"

// pooledObject wrap origin object to be pooled
type pooledObject struct {
	obj      any
	pool     *ObjectPool // nolint
	unusable bool
	mu       sync.RWMutex
}

// getObject return the origin object
func (p *pooledObject) getObject() any {
	return p.obj
}

// WrapObject wrap an origin object to a pooledObject
func (p *pooledObject) wrapObject(obj any) {
	p.obj = obj
}

// markUnusable marks the object not usable anymore.
func (p *pooledObject) markUnusable() {
	p.mu.Lock()
	p.unusable = true
	p.mu.Unlock()
}

// isUnusable returns whether this object is unusable
func (p *pooledObject) isUnusable() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.unusable
}
