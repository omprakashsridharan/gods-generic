package arraylist

type Iterator[T comparable] struct {
	list  *List[T]
	index int
}

func (list *List[T]) Iterator() Iterator[T] {
	return Iterator[T]{
		list:  list,
		index: -1,
	}
}

func (iterator Iterator[T]) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	return iterator.list.withinRange(iterator.index)
}

func (iterator *Iterator[T]) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.list.withinRange(iterator.index)
}

func (iterator *Iterator[T]) Value() T {
	return iterator.list.elements[iterator.index]
}

func (iterator *Iterator[T]) Index() int {
	return iterator.index
}

func (iterator *Iterator[T]) Begin() {
	iterator.index = -1
}

func (iterator *Iterator[T]) End() {
	iterator.index = iterator.list.size
}

func (iterator *Iterator[T]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

func (iterator *Iterator[T]) Last() bool {
	iterator.End()
	return iterator.Prev()
}
