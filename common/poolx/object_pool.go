package poolx

import (
	"errors"
	"sync"
)

var ErrClosed = errors.New("object pool is closed")

const (
	closeStateNormal uint32 = 0
	closeStateClosed uint32 = 1
)

// ObjectPool is a object pool
type ObjectPool struct {
	// createObject creates an origin object that can be served by the pool.
	createObject func() (any, error)
	// destroyObject destroys an origin object no longer needed by the pool.
	destroyObject func(obj any) error
	// activateObject reinitializes an instance to be returned by the pool.
	activateObject func(obj any) error
	// checkObject ensures that the origin object is safe to be returned by the pool.
	checkObject func(obj any) error
	// passivateObject uninitializes an origin object to be returned to the pool.
	passivateObject func(obj any) error
	mu              sync.RWMutex
	pooledObjectC   chan *pooledObject
	allObjects      sync.Map
	closed          uint32
}

type Option func(pool *ObjectPool)

func ObjectDestroyer(destroyObject func(obj any) error) Option {
	return func(pool *ObjectPool) {
		pool.destroyObject = destroyObject
	}
}

func ObjectChecker(checkObject func(obj any) error) Option {
	return func(pool *ObjectPool) {
		pool.checkObject = checkObject
	}
}

func ObjectActivator(activateObject func(obj any) error) Option {
	return func(pool *ObjectPool) {
		pool.activateObject = activateObject
	}
}

func ObjectPassivate(passivateObject func(obj any) error) Option {
	return func(pool *ObjectPool) {
		pool.passivateObject = passivateObject
	}
}

func NewObjectPool(
	maxTotal int64,
	createObject func() (any, error),
	opts ...Option,
) *ObjectPool {
	pool := &ObjectPool{
		createObject:    createObject,
		destroyObject:   func(obj any) error { return nil },
		checkObject:     func(obj any) error { return nil },
		activateObject:  func(obj any) error { return nil },
		passivateObject: func(obj any) error { return nil },
		pooledObjectC:   make(chan *pooledObject, maxTotal),
		closed:          closeStateNormal,
	}
	pool.apply(opts)
	pool.init()
	return pool
}

func (pool *ObjectPool) apply(opts []Option) {
	for _, opt := range opts {
		opt(pool)
	}
}

func (pool *ObjectPool) init() {

}

// Borrow borrows an object from the pool.
func (pool *ObjectPool) Borrow() (any, error) {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	if pool.isClosed() {
		return nil, ErrClosed
	}

	var p *pooledObject
Loop:
	for {
		var isCreate bool
		select {
		case p = <-pool.pooledObjectC:
			// get p from the pool.
		default:
			// the pool is emtpy, create an object
			pooledObject, err := pool.createPooledObject()
			if err != nil {
				return nil, err
			}
			pool.allObjects.Store(pooledObject.getObject(), pooledObject)
			isCreate = true
		}

		if p.isUnusable() {
			// if p is unusable. destroy int and continue loop
			_ = pool.destroyPooledObject(p)
			continue Loop
		}

		if err := pool.activateObject(p.getObject()); err != nil {
			// if object failed to reset, destroy it
			_ = pool.destroyPooledObject(p)
			if isCreate {
				// if the object was created, return error
				return nil, errors.New("failed to reset object")
			}
			// if the object was from poll, continue loop
			continue Loop
		}

		if err := pool.checkObject(p.getObject()); err != nil {
			// if object failed to check, destroy it
			_ = pool.destroyPooledObject(p)
			if isCreate {
				// if the object was created, return error
				return nil, errors.New("failed to reset object")
			}
			// if the object was from poll, continue loop
			continue Loop
		}
		break
	}
	return p.getObject(), nil
}

// Return returns the object to the pool.
func (pool *ObjectPool) Return(obj any) error {
	if obj == nil {
		return errors.New("object is nil")
	}
	pool.mu.Lock()
	defer pool.mu.Unlock()
	if pool.isClosed() {
		return ErrClosed
	}

	var p *pooledObject
	// gat pooledObject from allObjects by origin object
	value, ok := pool.allObjects.Load(obj)
	if !ok {
		p = new(pooledObject)
		p.wrapObject(obj)
		pool.allObjects.Store(obj, p)
		value = p
	}
	p, _ = value.(*pooledObject)

	if err := pool.passivateObject(obj); err != nil {
		_ = pool.destroyObject(obj)
		return err
	}
	select {
	case pool.pooledObjectC <- p:
		// put the object back into the pool.
		return nil
	default:
		// the pool is full, destroy the object
		return pool.destroyPooledObject(p)
	}
}

// Len returns the current number of object of the pool.
func (pool *ObjectPool) Len() (int, error) {
	pool.mu.RLock()
	defer pool.mu.RUnlock()
	if pool.isClosed() {
		return 0, ErrClosed
	}
	return len(pool.pooledObjectC), nil
}

// Clear clears the pool
func (pool *ObjectPool) Clear() error {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	if pool.isClosed() {
		return ErrClosed
	}
	return pool.clear()
}

// Close clears and closes the pool. After Close() the pool is no longer usable.
func (pool *ObjectPool) Close() error {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	if pool.isClosed() {
		return ErrClosed
	}
	pool.closed = closeStateClosed
	close(pool.pooledObjectC)
	return pool.clear()
}

// IsClosed returns whether the pool is closed.
func (pool *ObjectPool) IsClosed() bool {
	pool.mu.RLock()
	defer pool.mu.RUnlock()
	return pool.isClosed()
}

func (pool *ObjectPool) createPooledObject() (*pooledObject, error) {
	obj, err := pool.createObject()
	if err != nil {
		return nil, err
	}
	if obj == nil {
		return nil, errors.New("failed to create object")
	}
	poolObject := new(pooledObject)
	poolObject.wrapObject(obj)
	return poolObject, nil
}

func (pool *ObjectPool) destroyPooledObject(p *pooledObject) error {
	p.markUnusable()
	obj := p.getObject()
	pool.allObjects.Delete(obj)
	return pool.destroyObject(obj)
}

func (pool *ObjectPool) clear() error {
	for p := range pool.pooledObjectC {
		_ = pool.destroyPooledObject(p)
	}
	return nil
}

func (pool *ObjectPool) isClosed() bool {
	return pool.closed == closeStateClosed
}
