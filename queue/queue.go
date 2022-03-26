package queue

import "github.com/nomad-software/goad/constraint"

type Queue[T constraint.BuiltinTypes] struct {
	data []T
}

func New[T constraint.BuiltinTypes]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Count() int {
	return len(q.data)
}

func (q *Queue[T]) Empty() bool {
	return q.Count() == 0
}

func (q *Queue[T]) Enqueue(val T) {
	q.data = append(q.data, val)
}

func (q *Queue[T]) Peek() T {
	return q.data[0]
}

func (q *Queue[T]) Dequeue() T {
	if q.Empty() {
		panic("queue empty, popping failed")
	}

	val := q.data[0]
	q.data = q.data[1:q.Count()]
	return val
}

func (q *Queue[T]) Contains(needle T) bool {
	for _, v := range q.data {
		if v == needle {
			return true
		}
	}
	return false
}

func (s *Queue[T]) Clear() {
	s.data = s.data[:0]
}
