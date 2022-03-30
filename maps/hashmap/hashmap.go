package hashmap

import "fmt"

type Map[K comparable, V any] struct {
	m map[K]V
}

func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m: make(map[K]V),
	}
}

func (m *Map[K, V]) Put(Key K, value V) {
	m.m[Key] = value
}

func (m *Map[K, V]) Get(key K) (value V, found bool) {
	value, found = m.m[key]
	return
}

func (m *Map[K, V]) Remove(key K) {
	delete(m.m, key)
}

func (m *Map[K, V]) Empty() bool {
	return m.Size() == 0
}

func (m *Map[K, V]) Size() int {
	return len(m.m)
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, m.Size())
	count := 0
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

func (m *Map[K, V]) Values() []V {
	values := make([]V, m.Size())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count++
	}
	return values
}

func (m *Map[K, V]) Clear() {
	m.m = make(map[K]V)
}

func (m *Map[K, V]) String() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.m)
	return str
}
