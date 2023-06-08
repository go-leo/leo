package slicex

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/mathx/randx"
)

// Shuffle 打乱数组顺序
func Shuffle[S ~[]E, E any](s S) S {
	r := randx.Get()
	r.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	randx.Put(r)
	return s
}
