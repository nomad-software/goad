package set

import "github.com/nomad-software/goad/hashmap"

// Set is the main set type.
type Set[T comparable] struct {
	data hashmap.HashMap[T, any]
}

// New is used to create a new stack.
func New[T comparable]() Set[T] {
	return Set[T]{
		data: hashmap.New[T, any](),
	}
}

// Count returns the amount of entries in the set.
func (s Set[T]) Count() int {
	return s.data.Count()
}

// Empty returns true if the set is empty, false if not.
func (s Set[T]) Empty() bool {
	return s.data.Empty()
}

// Add adds a value to the set.
func (s *Set[T]) Add(value T) {
	s.data.Put(value, nil)
}

// Remove removes a value from the set.
func (s *Set[T]) Remove(value T) {
	s.data.Remove(value)
}

// Contains returns true if the value exists in the set, false if not.
func (s Set[T]) Contains(val T) bool {
	return s.data.ContainsKey(val)
}

// Clear empties the entire set.
func (s *Set[T]) Clear() {
	s.data.Clear()
}

// ForEach iterates over the dataset within the set, calling the passed
// function for each value.
func (s Set[T]) ForEach(f func(val T)) {
	s.data.ForEach(func(key T, val any) {
		f(key)
	})
}
