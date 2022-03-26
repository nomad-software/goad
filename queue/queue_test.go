package queue

import (
	"testing"

	"github.com/nomad-software/goad/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	s := New[int]()
	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)

	assert.Eq(t, s.Count(), 3)
	assert.Eq(t, s.Dequeue(), 1)
	assert.Eq(t, s.Peek(), 2)
	assert.Eq(t, s.Dequeue(), 2)
	assert.Eq(t, s.Dequeue(), 3)
}

func TestLargeCapacity(t *testing.T) {
	t.Parallel()

	s := New[int]()
	limit := 1_000_000

	for i := 1; i <= limit; i++ {
		s.Enqueue(i)
		assert.Eq(t, s.Peek(), 1)
		assert.Eq(t, s.Count(), i)
	}

	assert.Eq(t, s.Peek(), 1)
	assert.Eq(t, s.Count(), limit)
	assert.Eq(t, s.Contains(1), true)
	assert.Eq(t, s.Contains(limit), true)
	assert.Eq(t, s.Empty(), false)

	for i := 1; i <= limit; i++ {
		assert.Eq(t, s.Peek(), i)
		assert.Eq(t, s.Dequeue(), i)
		assert.Eq(t, s.Count(), limit-i)
	}

	assert.Eq(t, s.Empty(), true)
	assert.Eq(t, s.Count(), 0)
}

func TestFailedPop(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	s := New[string]()
	s.Dequeue()
}

func TestContains(t *testing.T) {
	t.Parallel()

	s := New[string]()
	s.Enqueue("foo")
	s.Enqueue("bar")
	s.Enqueue("baz")
	s.Enqueue("qux")

	assert.Eq(t, s.Contains("bar"), true)
	assert.Eq(t, s.Contains("fuz"), false)
}

func TestClearing(t *testing.T) {
	t.Parallel()

	s := New[int]()
	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)
	assert.Eq(t, s.Count(), 3)

	s.Clear()
	assert.Eq(t, s.Count(), 0)

	s.Enqueue(1)
	s.Enqueue(2)
	assert.Eq(t, s.Count(), 2)
}

func BenchmarkStack(b *testing.B) {
	s := New[int]()

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		s.Enqueue(1)
		s.Dequeue()
	}
}
