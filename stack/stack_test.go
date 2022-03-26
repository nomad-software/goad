package stack

import (
	"testing"

	"github.com/nomad-software/goad/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Eq(t, s.Empty(), false)
	assert.Eq(t, s.Count(), 3)
	assert.Eq(t, s.Pop(), 3)
	assert.Eq(t, s.Peek(), 2)
	assert.Eq(t, s.Pop(), 2)
	assert.Eq(t, s.Pop(), 1)
	assert.Eq(t, s.Empty(), true)
}

func TestLargeCapacity(t *testing.T) {
	t.Parallel()

	s := New[int]()
	limit := 1_000_000

	for i := 1; i <= limit; i++ {
		s.Push(i)
		assert.Eq(t, s.Peek(), i)
		assert.Eq(t, s.Count(), i)
	}

	assert.Eq(t, s.Peek(), limit)
	assert.Eq(t, s.Count(), limit)
	assert.Eq(t, s.Contains(1), true)
	assert.Eq(t, s.Contains(limit), true)
	assert.Eq(t, s.Empty(), false)

	for i := limit; i >= 1; i-- {
		assert.Eq(t, s.Count(), i)
		assert.Eq(t, s.Peek(), i)
		assert.Eq(t, s.Pop(), i)
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
	s.Pop()
}

func TestContains(t *testing.T) {
	t.Parallel()

	s := New[string]()
	s.Push("foo")
	s.Push("bar")
	s.Push("baz")
	s.Push("qux")

	assert.Eq(t, s.Contains("bar"), true)
	assert.Eq(t, s.Contains("fuz"), false)
}

func TestClearing(t *testing.T) {
	t.Parallel()

	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Eq(t, s.Count(), 3)

	s.Clear()
	assert.Eq(t, s.Count(), 0)

	s.Push(1)
	s.Push(2)
	assert.Eq(t, s.Count(), 2)
}

func BenchmarkStack(b *testing.B) {
	s := New[int]()

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		s.Push(1)
		s.Pop()
	}
}
