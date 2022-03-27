package stack

import (
	"github.com/nomad-software/goad/constraint"
)

type Stack[T constraint.BuiltinTypes] struct {
	data []T
}

func New[T constraint.BuiltinTypes]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Count() int {
	return len(s.data)
}

func (s *Stack[T]) Empty() bool {
	return s.Count() == 0
}

func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
}

func (s *Stack[T]) Peek() T {
	return s.data[s.Count()-1]
}

func (s *Stack[T]) Pop() T {
	if s.Empty() {
		panic("stack empty, popping failed")
	}

	val := s.data[s.Count()-1]
	s.data = s.data[0 : s.Count()-1]
	return val
}

func (s *Stack[T]) Contains(needle T) bool {
	for _, v := range s.data {
		if v == needle {
			return true
		}
	}
	return false
}

func (s *Stack[T]) Clear() {
	s.data = s.data[:0]
}

func (l *Stack[T]) ForEach(f func(val T)) {
	for i := len(l.data) - 1; i >= 0; i-- {
		f(l.data[i])
	}
}
