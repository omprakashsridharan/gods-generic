package containers

type IteratorWithIndex[T any] interface {
	Next() bool
	Index() int
	Begin()
	First() bool
	Value() T
}
