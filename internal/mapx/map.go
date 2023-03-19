package mapx

type Map[K comparable, V any] interface {
	Load(key K) (value V, ok bool)
	Store(key K, value V)
	LoadOrStore(key K, value V) (actual V, loaded bool)
	LoadAndDelete(key K) (value V, loaded bool)
	Delete(key K)
	DeleteFunc(del func(K, V) bool)
	Map() map[K]V
	Len() int
	Keys() []K
	Values() []V
	Clear()
	Clone() Map[K, V]
}
