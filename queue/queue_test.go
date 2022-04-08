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

	assert.False(t, q.Empty())
	assert.Eq(t, q.Count(), 3)
	assert.Eq(t, q.Dequeue(), 1)
	assert.Eq(t, q.Peek(), 2)
	assert.Eq(t, q.Dequeue(), 2)
	assert.Eq(t, q.Dequeue(), 3)
	assert.True(t, q.Empty())
}

func TestChannel(t *testing.T) {
	t.Parallel()

	s := New[chan string]()

	c1 := make(chan string)
	c2 := make(chan string)

	s.Enqueue(c1)
	s.Enqueue(c2)

	assert.False(t, s.Empty())
	assert.Eq(t, s.Count(), 2)
	assert.True(t, s.Contains(c1))
	assert.True(t, s.Contains(c2))
	assert.Eq(t, s.Dequeue(), c1)
	assert.Eq(t, s.Dequeue(), c2)
	assert.True(t, s.Empty())
}

func TestArray(t *testing.T) {
	t.Parallel()

	s := New[[3]bool]()

	a1 := [3]bool{true, false, true}
	a2 := [3]bool{false, false, true}

	s.Enqueue(a1)
	s.Enqueue(a2)

	assert.False(t, s.Empty())
	assert.Eq(t, s.Count(), 2)
	assert.True(t, s.Contains(a1))
	assert.True(t, s.Contains(a2))
	assert.Eq(t, s.Dequeue(), a1)
	assert.Eq(t, s.Dequeue(), a2)
	assert.True(t, s.Empty())
}

func TestStruct(t *testing.T) {
	t.Parallel()

	type Foo struct {
		Foo string
		Bar string
	}

	s := New[Foo]()

	f1 := Foo{Foo: "foo", Bar: "bar"}
	f2 := Foo{Foo: "baz", Bar: "qux"}

	s.Enqueue(f1)
	s.Enqueue(f2)

	assert.False(t, s.Empty())
	assert.Eq(t, s.Count(), 2)
	assert.True(t, s.Contains(f1))
	assert.True(t, s.Contains(f2))
	assert.Eq(t, s.Dequeue(), f1)
	assert.Eq(t, s.Dequeue(), f2)
	assert.True(t, s.Empty())
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
	assert.True(t, q.Contains(1))
	assert.True(t, q.Contains(limit))
	assert.False(t, q.Empty())

	for i := 1; i <= limit; i++ {
		assert.Eq(t, q.Peek(), i)
		assert.Eq(t, q.Dequeue(), i)
		assert.Eq(t, q.Count(), limit-i)
	}

	assert.True(t, q.Empty())
	assert.Eq(t, q.Count(), 0)
}

func TestFailedDequeue(t *testing.T) {
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

	assert.True(t, q.Contains("bar"))
	assert.False(t, q.Contains("fuz"))
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

	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	i := 1
	q.ForEach(func(val int) {
		assert.Eq(t, val, i)
		i++
	})

	q.Clear()
	q.ForEach(func(val int) {
		t.Errorf("queue not cleared")
	})
}

func TestValues(t *testing.T) {
	t.Parallel()

	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	i := 1
	for val := range q.Values() {
		assert.Eq(t, val, i)
		i++
	}

	q.Clear()
	for range q.Values() {
		t.Errorf("queue not cleared")
	}
}

func BenchmarkQueue(b *testing.B) {
	q := New[int]()

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		q.Enqueue(x)
		q.Dequeue()
	}
}
