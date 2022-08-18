package mapx

type Hash[K comparable] interface {
	Sum(k K) int64
}

type Hash64Func[K comparable] func(k K) int64

func (f Hash64Func[K]) Sum(k K) int64 {
	return f(k)
}
