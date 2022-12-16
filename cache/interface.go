package cache

type Interface[K comparable, V any] interface {
	Get(key K) V
	Put(key K, value V)
}
