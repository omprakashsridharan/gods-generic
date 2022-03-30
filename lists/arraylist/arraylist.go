package arraylist

import (
	"fmt"
	"strings"
)

type List[T comparable] struct {
	elements []T
	size     int
}

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List[T]) Add(values ...T) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

func (list *List[T]) growBy(n int) {
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

func (list *List[T]) resize(cap int) {
	newElements := make([]T, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *List[T]) Get(index int) (T, bool) {
	var result T
	if !list.withinRange(index) {
		return result, false
	}
	return list.elements[index], true
}

func (list *List[T]) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *List[T]) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	list.shrink()
}

func (list *List[T]) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}

func (list List[T]) Contains(values ...T) bool {
	for _, searchValue := range values {
		found := false
		for _, element := range list.elements {
			if element == searchValue {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *List[T]) Values() []T {
	newElements := make([]T, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

func (list *List[T]) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.elements {
		if element == value {
			return index
		}
	}
	return -1
}

func (list *List[T]) Empty() bool {
	return list.size == 0
}

func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) Clear() {
	list.size = 0
	list.elements = []T{}
}

func (list *List[T]) Insert(index int, values ...T) {

	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-l])
	copy(list.elements[index:], values)
}

func (list *List[T]) Set(index int, value T) {

	if !list.withinRange(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}

	list.elements[index] = value
}

func (list *List[T]) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
