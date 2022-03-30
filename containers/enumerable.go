package containers

type EnumerableWithIndex[T any] interface {
	Each(func(index int, value T))
	Any(func(index int, value T) bool) bool
	All(func(index int, value T) bool) bool
	Find(func(index int, value T) bool) (int, T)
}

type EnumerableWithKey[K, V any] interface {
	Each(func(key K, value V))
	Any(func(key K, value V) bool) bool
	All(func(key K, value V) bool) bool
	Find(func(key K, value V) bool) (K, V)
}
