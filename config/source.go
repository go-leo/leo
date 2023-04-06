package config

import (
	"sync"

	"github.com/go-leo/gox/slicex"
)

type Source struct {
	// Name of the config
	name string
	//  Value from config
	value []byte
	// Extension identifies the type of Value
	extension string
	observers []SourceObserver
	mutex     sync.RWMutex
}

func (s *Source) Name() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.name
}

func (s *Source) SetName(name string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.name = name
	s.notifyAll()
}

func (s *Source) Value() []byte {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.value
}

func (s *Source) SetValue(value []byte) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.value = value
	s.notifyAll()
}

func (s *Source) Extension() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.extension
}

func (s *Source) SetExtension(extension string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.extension = extension
	s.notifyAll()
}

func (s *Source) AddObserver(so SourceObserver) {
	if so == nil {
		return
	}
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	s.observers = append(s.observers, so)
}

func (s *Source) DeleteObserver(so SourceObserver) {
	if so == nil {
		return
	}
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	s.observers = slicex.Remove(s.observers, so)
}

func (s *Source) notifyAll() {
	for _, observer := range s.observers {
		observer.Update(s)
	}
}

type SourceObserver interface {
	Update(s *Source)
}

func NewSource(name string, value []byte, extension string) *Source {
	return &Source{name: name, value: value, extension: extension}
}
