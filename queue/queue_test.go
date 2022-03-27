package queue

import (
	"testing"

	"github.com/nomad-software/goad/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Eq(t, q.Empty(), false)
	assert.Eq(t, q.Count(), 3)
	assert.Eq(t, q.Dequeue(), 1)
	assert.Eq(t, q.Peek(), 2)
	assert.Eq(t, q.Dequeue(), 2)
	assert.Eq(t, q.Dequeue(), 3)
	assert.Eq(t, q.Empty(), true)
}

func TestLargeCapacity(t *testing.T) {
	t.Parallel()

	q := New[int]()
	limit := 1_000_000

	for i := 1; i <= limit; i++ {
		q.Enqueue(i)
		assert.Eq(t, q.Peek(), 1)
		assert.Eq(t, q.Count(), i)
	}

	assert.Eq(t, q.Peek(), 1)
	assert.Eq(t, q.Count(), limit)
	assert.Eq(t, q.Contains(1), true)
	assert.Eq(t, q.Contains(limit), true)
	assert.Eq(t, q.Empty(), false)

	for i := 1; i <= limit; i++ {
		assert.Eq(t, q.Peek(), i)
		assert.Eq(t, q.Dequeue(), i)
		assert.Eq(t, q.Count(), limit-i)
	}

	assert.Eq(t, q.Empty(), true)
	assert.Eq(t, q.Count(), 0)
}

func TestFailedPop(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	q := New[string]()
	q.Dequeue()
}

func TestContains(t *testing.T) {
	t.Parallel()

	q := New[string]()
	q.Enqueue("foo")
	q.Enqueue("bar")
	q.Enqueue("baz")
	q.Enqueue("qux")

	assert.Eq(t, q.Contains("bar"), true)
	assert.Eq(t, q.Contains("fuz"), false)
}

func TestClearing(t *testing.T) {
	t.Parallel()

	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Eq(t, q.Count(), 3)

	q.Clear()
	assert.Eq(t, q.Count(), 0)

	q.Enqueue(1)
	q.Enqueue(2)
	assert.Eq(t, q.Count(), 2)
}

func TestForEach(t *testing.T) {
	t.Parallel()

	l := New[int]()

	l.Enqueue(1)
	l.Enqueue(2)
	l.Enqueue(3)
	l.Enqueue(4)
	l.Enqueue(5)

	i := 1
	l.ForEach(func(val int) {
		assert.Eq(t, val, i)
		i++
	})

	l.Clear()
	l.ForEach(func(val int) {
		t.Errorf("queue not cleared")
	})
}

func BenchmarkQueue(b *testing.B) {
	q := New[int]()

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		q.Enqueue(1)
		q.Dequeue()
	}
}
