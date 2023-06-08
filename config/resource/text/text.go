package text

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/slicex"
	"golang.org/x/exp/slices"
	"sync"
	"sync/atomic"
)

type Text struct {
	value     atomic.Value
	mu        sync.RWMutex
	observers []Observer
}

type Observer interface {
	onChanged(newText, oldText string)
}

type ObserverFunc func(newText, oldText string)

func (f ObserverFunc) onChanged(newText, oldText string) {
	f(newText, oldText)
}

func NewText(t string) *Text {
	text := &Text{}
	text.value.Store(t)
	return text
}

func (t *Text) Get() string {
	return t.value.Load().(string)
}

func (t *Text) Set(text string) {
	old := t.value.Swap(text).(string)
	var observers []Observer
	t.mu.RLock()
	observers = slices.Clone(t.observers)
	t.mu.RUnlock()
	for _, observer := range observers {
		observer.onChanged(text, old)
	}
}

func (t *Text) AddObserver(observer Observer) {
	t.mu.Lock()
	t.observers = append(t.observers, observer)
	t.mu.Unlock()
}

func (t *Text) RemoveObserver(observer Observer) {
	t.mu.Lock()
	t.observers = slicex.Remove(t.observers, observer)
	t.mu.Unlock()
}
