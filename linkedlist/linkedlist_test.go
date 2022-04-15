package linkedlist

import (
	"testing"

	"github.com/nomad-software/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	l := New[int]()

	assert.Eq(t, l.Count(), 0)
	assert.True(t, l.Empty())
}

func TestChannel(t *testing.T) {
	t.Parallel()

	s := New[chan string]()

	c1 := make(chan string)
	c2 := make(chan string)

	s.InsertLast(c1)
	s.InsertLast(c2)

	assert.False(t, s.Empty())
	assert.Eq(t, s.Count(), 2)
	assert.True(t, s.Contains(c1))
	assert.True(t, s.Contains(c2))
	assert.Eq(t, s.Last(), c2)
	s.RemoveLast()
	assert.Eq(t, s.Last(), c1)
	s.RemoveLast()
	assert.True(t, s.Empty())
}

func TestArray(t *testing.T) {
	t.Parallel()

	s := New[[3]bool]()

	a1 := [3]bool{true, false, true}
	a2 := [3]bool{false, false, true}

	s.InsertLast(a1)
	s.InsertLast(a2)

	assert.False(t, s.Empty())
	assert.Eq(t, s.Count(), 2)
	assert.True(t, s.Contains(a1))
	assert.True(t, s.Contains(a2))
	assert.Eq(t, s.Last(), a2)
	s.RemoveLast()
	assert.Eq(t, s.Last(), a1)
	s.RemoveLast()
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

	s.InsertLast(f1)
	s.InsertLast(f2)

	assert.False(t, s.Empty())
	assert.Eq(t, s.Count(), 2)
	assert.True(t, s.Contains(f1))
	assert.True(t, s.Contains(f2))
	assert.Eq(t, s.Last(), f2)
	s.RemoveLast()
	assert.Eq(t, s.Last(), f1)
	s.RemoveLast()
	assert.True(t, s.Empty())
}

func TestLargeCapacity(t *testing.T) {
	t.Parallel()

	s := New[uint]()
	limit := uint(1_000_000)

	for i := uint(1); i <= limit; i++ {
		s.InsertLast(i)
		assert.Eq(t, s.Last(), i)
		assert.Eq(t, s.Count(), int(i))
	}

	assert.Eq(t, s.Count(), int(limit))
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(limit))
	assert.False(t, s.Empty())

	for i := limit; i >= 1; i-- {
		assert.Eq(t, s.Count(), int(i))
		assert.Eq(t, s.Last(), i)
		s.RemoveLast()
	}

	assert.True(t, s.Empty())
	assert.Eq(t, s.Count(), 0)
}

func TestInsertFirst(t *testing.T) {
	t.Parallel()

	l := New[int]()

	l.InsertFirst(1)
	assert.Eq(t, l.First(), 1)
	assert.Eq(t, l.Last(), 1)

	l.InsertFirst(2)
	assert.Eq(t, l.First(), 2)
	assert.Eq(t, l.Last(), 1)

	l.InsertFirst(3)
	assert.Eq(t, l.First(), 3)
	assert.Eq(t, l.Last(), 1)

	assert.Eq(t, l.Count(), 3)
	assert.False(t, l.Empty())

	l.RemoveFirst()
	assert.Eq(t, l.First(), 2)
	assert.Eq(t, l.Last(), 1)

	l.RemoveFirst()
	assert.Eq(t, l.First(), 1)
	assert.Eq(t, l.Last(), 1)

	l.RemoveFirst()
	assert.Eq(t, l.Count(), 0)
	assert.True(t, l.Empty())
}

func TestFailedFirst(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	q := New[string]()
	q.First()
}

func TestInsertLast(t *testing.T) {
	t.Parallel()

	l := New[int]()

	l.InsertLast(1)
	assert.Eq(t, l.First(), 1)
	assert.Eq(t, l.Last(), 1)

	l.InsertLast(2)
	assert.Eq(t, l.First(), 1)
	assert.Eq(t, l.Last(), 2)

	l.InsertLast(3)
	assert.Eq(t, l.First(), 1)
	assert.Eq(t, l.Last(), 3)

	assert.Eq(t, l.Count(), 3)
	assert.False(t, l.Empty())

	l.RemoveLast()
	assert.Eq(t, l.First(), 1)
	assert.Eq(t, l.Last(), 2)

	l.RemoveLast()
	assert.Eq(t, l.First(), 1)
	assert.Eq(t, l.Last(), 1)

	l.RemoveLast()
	assert.Eq(t, l.Count(), 0)
	assert.True(t, l.Empty())
}

func TestFailedLast(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	q := New[string]()
	q.Last()
}

func TestInsert(t *testing.T) {
	t.Parallel()

	l := New[int]()

	l.Insert(1, 0)
	assert.Eq(t, l.First(), 1)
	assert.Eq(t, l.Last(), 1)

	l.Insert(2, 1)
	assert.Eq(t, l.First(), 1)
	assert.Eq(t, l.Last(), 2)

	l.Insert(3, 0)
	assert.Eq(t, l.First(), 3)
	assert.Eq(t, l.Last(), 2)

	l.Insert(4, 1)
	assert.Eq(t, l.First(), 3)
	assert.Eq(t, l.Last(), 2)

	assert.Eq(t, l.Get(0), 3)
	assert.Eq(t, l.Get(1), 4)
	assert.Eq(t, l.Get(2), 1)
	assert.Eq(t, l.Get(3), 2)

	assert.Eq(t, l.Count(), 4)
	assert.False(t, l.Empty())
}

func TestFailedInsert(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	q := New[string]()
	q.Insert("foo", 1)
}

func TestFailedGet(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	q := New[string]()
	q.Get(1)
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	l := New[int]()

	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)
	l.InsertLast(4)

	assert.Eq(t, l.Get(0), 1)
	assert.Eq(t, l.Get(1), 2)
	assert.Eq(t, l.Get(2), 3)
	assert.Eq(t, l.Get(3), 4)

	l.Update(2, 5)

	assert.Eq(t, l.Get(0), 1)
	assert.Eq(t, l.Get(1), 2)
	assert.Eq(t, l.Get(2), 5)
	assert.Eq(t, l.Get(3), 4)

	assert.Eq(t, l.Count(), 4)
	assert.False(t, l.Empty())
}

func TestFailedUpdate(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	q := New[string]()
	q.Update(1, "foo")
}

func TestRemove(t *testing.T) {
	t.Parallel()

	l := New[int]()

	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)
	l.InsertLast(4)
	l.InsertLast(5)

	assert.Eq(t, l.Get(0), 1)
	assert.Eq(t, l.Get(1), 2)
	assert.Eq(t, l.Get(2), 3)
	assert.Eq(t, l.Get(3), 4)
	assert.Eq(t, l.Get(4), 5)

	assert.Eq(t, l.Count(), 5)
	assert.False(t, l.Empty())

	l.Remove(4)
	assert.Eq(t, l.Get(0), 1)
	assert.Eq(t, l.Get(1), 2)
	assert.Eq(t, l.Get(2), 3)
	assert.Eq(t, l.Get(3), 4)

	assert.Eq(t, l.Count(), 4)
	assert.False(t, l.Empty())

	l.Remove(0)
	assert.Eq(t, l.Get(0), 2)
	assert.Eq(t, l.Get(1), 3)
	assert.Eq(t, l.Get(2), 4)

	assert.Eq(t, l.Count(), 3)
	assert.False(t, l.Empty())

	l.Remove(1)
	assert.Eq(t, l.Get(0), 2)
	assert.Eq(t, l.Get(1), 4)

	assert.Eq(t, l.Count(), 2)
	assert.False(t, l.Empty())
}

func TestFailedRemoveOnEmptyList(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	l := New[string]()
	l.Remove(1)
}

func TestFailedRemoveOutsideOfBounds(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	l := New[string]()
	l.InsertLast("foo")
	l.Remove(5)
}

func TestContains(t *testing.T) {
	t.Parallel()

	l := New[int]()

	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)

	assert.False(t, l.Contains(0))
	assert.True(t, l.Contains(1))
	assert.True(t, l.Contains(2))
	assert.True(t, l.Contains(3))
	assert.False(t, l.Contains(4))
}

func TestClear(t *testing.T) {
	t.Parallel()

	l := New[int]()

	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)

	l.Clear()
	assert.Eq(t, l.Count(), 0)
	assert.True(t, l.Empty())

	l.InsertLast(4)
	l.InsertLast(5)
	l.InsertLast(6)

	assert.Eq(t, l.Count(), 3)
	assert.False(t, l.Empty())
}

func TestForEach(t *testing.T) {
	t.Parallel()

	l := New[int]()

	l.InsertLast(0)
	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)
	l.InsertLast(4)

	l.ForEach(func(i int, val int) {
		assert.Eq(t, val, i)
	})

	l.Clear()
	l.ForEach(func(i int, val int) {
		t.Errorf("linked list not cleared")
	})
}

func BenchmarkLinkedList(b *testing.B) {
	l := New[int]()

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		l.InsertLast(x)
		l.RemoveLast()
	}
}

func BenchmarkForEach(b *testing.B) {
	l := New[int]()

	for x := 0; x < 1_000_000; x++ {
		l.InsertLast(x)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		l.ForEach(func(index int, val int) {
		})
	}
}
