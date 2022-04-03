package linkedlist

import (
	"testing"

	"github.com/nomad-software/goad/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	l := New[int]()

	assert.Eq(t, l.Count(), 0)
	assert.Eq(t, l.Empty(), true)
}

func TestChannel(t *testing.T) {
	t.Parallel()

	s := New[chan string]()

	c1 := make(chan string)
	c2 := make(chan string)

	s.InsertLast(c1)
	s.InsertLast(c2)

	assert.Eq(t, s.Empty(), false)
	assert.Eq(t, s.Count(), 2)
	assert.Eq(t, s.Contains(c1), true)
	assert.Eq(t, s.Contains(c2), true)
	assert.Eq(t, s.Last(), c2)
	s.RemoveLast()
	assert.Eq(t, s.Last(), c1)
	s.RemoveLast()
	assert.Eq(t, s.Empty(), true)
}

func TestArray(t *testing.T) {
	t.Parallel()

	s := New[[3]bool]()

	a1 := [3]bool{true, false, true}
	a2 := [3]bool{false, false, true}

	s.InsertLast(a1)
	s.InsertLast(a2)

	assert.Eq(t, s.Empty(), false)
	assert.Eq(t, s.Count(), 2)
	assert.Eq(t, s.Contains(a1), true)
	assert.Eq(t, s.Contains(a2), true)
	assert.Eq(t, s.Last(), a2)
	s.RemoveLast()
	assert.Eq(t, s.Last(), a1)
	s.RemoveLast()
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

	s.InsertLast(f1)
	s.InsertLast(f2)

	assert.Eq(t, s.Empty(), false)
	assert.Eq(t, s.Count(), 2)
	assert.Eq(t, s.Contains(f1), true)
	assert.Eq(t, s.Contains(f2), true)
	assert.Eq(t, s.Last(), f2)
	s.RemoveLast()
	assert.Eq(t, s.Last(), f1)
	s.RemoveLast()
	assert.Eq(t, s.Empty(), true)
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
	assert.Eq(t, s.Contains(1), true)
	assert.Eq(t, s.Contains(limit), true)
	assert.Eq(t, s.Empty(), false)

	for i := limit; i >= 1; i-- {
		assert.Eq(t, s.Count(), int(i))
		assert.Eq(t, s.Last(), i)
		s.RemoveLast()
	}

	assert.Eq(t, s.Empty(), true)
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
	assert.Eq(t, l.Empty(), false)

	l.RemoveFirst()
	assert.Eq(t, l.First(), 2)
	assert.Eq(t, l.Last(), 1)

	l.RemoveFirst()
	assert.Eq(t, l.First(), 1)
	assert.Eq(t, l.Last(), 1)

	l.RemoveFirst()
	assert.Eq(t, l.Count(), 0)
	assert.Eq(t, l.Empty(), true)
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
	assert.Eq(t, l.Empty(), false)

	l.RemoveLast()
	assert.Eq(t, l.First(), 1)
	assert.Eq(t, l.Last(), 2)

	l.RemoveLast()
	assert.Eq(t, l.First(), 1)
	assert.Eq(t, l.Last(), 1)

	l.RemoveLast()
	assert.Eq(t, l.Count(), 0)
	assert.Eq(t, l.Empty(), true)
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
	assert.Eq(t, l.Empty(), false)
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

	l.Update(5, 2)

	assert.Eq(t, l.Get(0), 1)
	assert.Eq(t, l.Get(1), 2)
	assert.Eq(t, l.Get(2), 5)
	assert.Eq(t, l.Get(3), 4)

	assert.Eq(t, l.Count(), 4)
	assert.Eq(t, l.Empty(), false)
}

func TestFailedUpdate(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	q := New[string]()
	q.Update("foo", 1)
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
	assert.Eq(t, l.Empty(), false)

	l.Remove(4)
	assert.Eq(t, l.Get(0), 1)
	assert.Eq(t, l.Get(1), 2)
	assert.Eq(t, l.Get(2), 3)
	assert.Eq(t, l.Get(3), 4)

	assert.Eq(t, l.Count(), 4)
	assert.Eq(t, l.Empty(), false)

	l.Remove(0)
	assert.Eq(t, l.Get(0), 2)
	assert.Eq(t, l.Get(1), 3)
	assert.Eq(t, l.Get(2), 4)

	assert.Eq(t, l.Count(), 3)
	assert.Eq(t, l.Empty(), false)

	l.Remove(1)
	assert.Eq(t, l.Get(0), 2)
	assert.Eq(t, l.Get(1), 4)

	assert.Eq(t, l.Count(), 2)
	assert.Eq(t, l.Empty(), false)
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

	assert.Eq(t, l.Contains(0), false)
	assert.Eq(t, l.Contains(1), true)
	assert.Eq(t, l.Contains(2), true)
	assert.Eq(t, l.Contains(3), true)
	assert.Eq(t, l.Contains(4), false)
}

func TestClear(t *testing.T) {
	t.Parallel()

	l := New[int]()

	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)

	l.Clear()
	assert.Eq(t, l.Count(), 0)
	assert.Eq(t, l.Empty(), true)

	l.InsertLast(4)
	l.InsertLast(5)
	l.InsertLast(6)

	assert.Eq(t, l.Count(), 3)
	assert.Eq(t, l.Empty(), false)
}

func TestForEach(t *testing.T) {
	t.Parallel()

	l := New[int]()

	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)
	l.InsertLast(4)
	l.InsertLast(5)

	i := 1
	l.ForEach(func(val int) {
		assert.Eq(t, val, i)
		i++
	})

	l.Clear()
	l.ForEach(func(val int) {
		t.Errorf("linked list not cleared")
	})
}

func BenchmarkLinkedList(b *testing.B) {
	s := New[int]()

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		s.InsertLast(x)
		s.RemoveLast()
	}
}
