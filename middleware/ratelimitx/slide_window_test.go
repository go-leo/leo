package ratelimitx

import (
	"sync"
	"testing"
)

func TestLock(t *testing.T) {
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	t.Log("hello")
	mu.Lock()
	defer mu.Unlock()
}
