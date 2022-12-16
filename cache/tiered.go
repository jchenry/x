package cache

import "reflect"

type tieredCache[K comparable, V any] struct {
	inner Interface[K, V]
	outer Interface[K, V]
}

func NewTieredCache[K comparable, V any](inner, outer Interface[K, V]) Interface[K, V] {
	return &tieredCache[K, V]{
		inner: inner,
		outer: outer,
	}
}

func (t *tieredCache[K, V]) Get(key K) V {
	var zero, value V
	value = t.inner.Get(key)
	if reflect.DeepEqual(value, zero) {
		value = t.outer.Get(key)
		// if required, add value to inner cache for future requests
	}
	return value
}

func (t *tieredCache[K, V]) Put(key K, value V) {
	t.inner.Put(key, value)

	// add key to outer cache asynchronously
	go func(key K) {
		t.outer.Put(key, value)
	}(key)
}
