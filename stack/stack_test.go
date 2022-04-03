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

func TestChannel(t *testing.T) {
	t.Parallel()

	s := New[chan string]()

	c1 := make(chan string)
	c2 := make(chan string)

	s.Push(c1)
	s.Push(c2)

	assert.Eq(t, s.Empty(), false)
	assert.Eq(t, s.Count(), 2)
	assert.Eq(t, s.Contains(c1), true)
	assert.Eq(t, s.Contains(c2), true)
	assert.Eq(t, s.Pop(), c2)
	assert.Eq(t, s.Pop(), c1)
	assert.Eq(t, s.Empty(), true)
}

func TestArray(t *testing.T) {
	t.Parallel()

	s := New[[3]bool]()

	a1 := [3]bool{true, false, true}
	a2 := [3]bool{false, false, true}

	s.Push(a1)
	s.Push(a2)

	assert.Eq(t, s.Empty(), false)
	assert.Eq(t, s.Count(), 2)
	assert.Eq(t, s.Contains(a1), true)
	assert.Eq(t, s.Contains(a2), true)
	assert.Eq(t, s.Pop(), a2)
	assert.Eq(t, s.Pop(), a1)
	assert.Eq(t, s.Empty(), true)
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

	s.Push(f1)
	s.Push(f2)

	assert.Eq(t, s.Empty(), false)
	assert.Eq(t, s.Count(), 2)
	assert.Eq(t, s.Contains(f1), true)
	assert.Eq(t, s.Contains(f2), true)
	assert.Eq(t, s.Pop(), f2)
	assert.Eq(t, s.Pop(), f1)
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

func TestForEach(t *testing.T) {
	t.Parallel()

	s := New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	i := s.Count()
	s.ForEach(func(val int) {
		assert.Eq(t, val, i)
		i--
	})

	s.Clear()
	s.ForEach(func(val int) {
		t.Errorf("stack not cleared")
	})
}

func BenchmarkStack(b *testing.B) {
	s := New[int]()

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		s.Push(x)
		s.Pop()
	}
}
