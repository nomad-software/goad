package queue

import "github.com/nomad-software/goad/constraint"

type Queue[T constraint.BuiltinTypes] struct {
	data []T
}

func New[T constraint.BuiltinTypes]() Queue[T] {
	return Queue[T]{}
}

func (s *Queue[T]) Count() int {
	return len(s.data)
}

func (s *Queue[T]) Empty() bool {
	return s.Count() == 0
}

func (s *Queue[T]) Enqueue(val T) {
	s.data = append(s.data, val)
}

func (s *Queue[T]) Peek() T {
	return s.data[0]
}

func (s *Queue[T]) Dequeue() T {
	if s.Empty() {
		panic("queue empty, popping failed")
	}

	val := s.data[0]
	s.data = s.data[1:s.Count()]
	return val
}

func (s *Queue[T]) Contains(needle T) bool {
	for _, v := range s.data {
		if v == needle {
			return true
		}
	}
	return false
}

func (s *Queue[T]) Clear() {
	s.data = s.data[:0]
}
