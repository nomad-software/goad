package queue

// Queue is the main queue type.
type Queue[T comparable] struct {
	data []T
}

// New is used to create a new queue.
func New[T comparable]() Queue[T] {
	return Queue[T]{}
}

// Count returns the amount of entries in the queue.
func (q Queue[T]) Count() int {
	return len(q.data)
}

// Empty returns true if the queue is empty, false if not.
func (q Queue[T]) Empty() bool {
	return q.Count() == 0
}

// Enqueue add a value to the queue.
func (q *Queue[T]) Enqueue(val T) {
	q.data = append(q.data, val)
}

// Peek returns the first value.
func (q Queue[T]) Peek() T {
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
func (q Queue[T]) Contains(val T) bool {
	for _, v := range q.data {
		if v == val {
			return true
		}
	}
	return false
}

// Clear empties the entire queue.
func (q *Queue[T]) Clear() {
	q.data = q.data[:0]
}

// ForEach iterates over the dataset within the queue, calling the passed
// function for each value.
func (q Queue[T]) ForEach(f func(val T)) {
	for _, v := range q.data {
		f(v)
	}
}

// Values returns the values delivered through a channel. This is safe to be
// called in a for/range loop as it only creates one channel.
func (q Queue[T]) Values() chan T {
	c := make(chan T)

	go func() {
		for _, v := range q.data {
			c <- v
		}
		close(c)
	}()

	return c
}
