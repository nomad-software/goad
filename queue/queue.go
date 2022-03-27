package queue

import "github.com/nomad-software/goad/constraint"

// Queue is the main queue type.
type Queue[T constraint.BuiltinTypes] struct {
	data []T
}

// New is used to create a new queue.
func New[T constraint.BuiltinTypes]() Queue[T] {
	return Queue[T]{}
}

// Count returns the amount of entries in the queue.
func (q *Queue[T]) Count() int {
	return len(q.data)
}

// Empty returns true if the queue is empty, false if not.
func (q *Queue[T]) Empty() bool {
	return q.Count() == 0
}

// Enqueue add a value to the queue.
func (q *Queue[T]) Enqueue(val T) {
	q.data = append(q.data, val)
}

// Peek returns the first value.
func (q *Queue[T]) Peek() T {
	return q.data[0]
}

// Dequeue returns the first value and removes it.
func (q *Queue[T]) Dequeue() T {
	if q.Empty() {
		panic("queue empty, popping failed")
	}
	val := q.data[0]
	q.data = q.data[1:q.Count()]
	return val
}

// Contains returns true if the value exists in the queue, false if not.
func (q *Queue[T]) Contains(val T) bool {
	for _, v := range q.data {
		if v == val {
			return true
		}
	}
	return false
}

// Clear empties the entire queue.
func (s *Queue[T]) Clear() {
	s.data = s.data[:0]
}

// ForEach iterates over the dataset within the queue, calling the passed
// function for each value.
func (l *Queue[T]) ForEach(f func(val T)) {
	for _, v := range l.data {
		f(v)
	}
}
