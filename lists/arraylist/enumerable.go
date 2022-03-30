package arraylist

func (list *List[T]) Each(f func(index int, value T)) {
	iterator := list.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

func (list *List[T]) Select(f func(index int, value T) bool) *List[T] {
	newList := &List[T]{}
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newList.Add(iterator.Value())
		}
	}
	return newList
}

func (list *List[T]) Any(f func(index int, value T) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

func (list *List[T]) All(f func(index int, value T) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

func (list *List[T]) Find(f func(index int, value T) bool) (int, T) {
	iterator := list.Iterator()
	var result T
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	return -1, result
}
