package slicex

import (
	"math/rand"
	"time"
)

var shuffleRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// Shuffle 打乱数组顺序
func Shuffle[S ~[]E, E any](s S) S {
	shuffleRand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s
}
