package mapx

import (
	"sync"
	"testing"
)

func TestStdMap(t *testing.T) {
	var m Map[string, int] = NewStdMap[string, int](10)
	m.Store("a", 1)
	m.Store("b", 2)
	m.Store("c", 3)
	value, ok := m.Load("a")
	if !ok {
		t.Error("a was not found while getting a")
	}
	if value != 1 {
		t.Error("a was not equal 1")
	}

	s := new(sync.Map)
	s.Store("1", nil)
	load, o := s.Load("1")
	t.Log(load, o)
	s.Store(nil, 1)
	load, o = s.Load(nil)
	t.Log(load, o)

}
