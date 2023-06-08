package slicex

// ToMap 方法创建一个Map，这个Map由原数组中的每个元素都调用一次提供的函数后的返回值作为Key、每个元素作为Value组成。
func ToMap[S ~[]E, M ~map[K]E, E any, K comparable](s S, f func(int, E) K) M {
	m := make(M, len(s))
	for i, v := range s {
		m[f(i, v)] = v
	}
	return m
}
