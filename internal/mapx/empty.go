package mapx

func IsEmpty[M ~map[K]V, K comparable, V any](m M) bool {
	return len(m) <= 0
}

func IsNotEmpty[M ~map[K]V, K comparable, V any](m M) bool {
	return len(m) > 0
}
