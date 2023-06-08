package randx

import (
	"math/rand"
	"time"
)

var globalRand *rand.Rand

func init() {
	globalRand = rand.New(NewSyncSource(time.Now().UnixNano()))
}

func Rand() rand.Source {
	return globalRand
}
