package randx

import (
	"math/rand"
	"sync"
	"time"
)

var randPool = &sync.Pool{New: func() any { return rand.New(rand.NewSource(time.Now().UnixNano())) }}

func Get() *rand.Rand {
	r := randPool.Get().(*rand.Rand)
	return r
}

func Put(r *rand.Rand) {
	randPool.Put(r)
}
