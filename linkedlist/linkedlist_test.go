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

func TestLargeCapacity(t *testing.T) {
	t.Parallel()

	s := New[uint]()
	limit := uint(1_000_000)

	for i := uint(1); i <= limit; i++ {
		s.InsertLast(i)
		assert.Eq(t, s.Last(), i)
		assert.Eq(t, s.Count(), i)
	}

	assert.Eq(t, s.Count(), limit)
	assert.Eq(t, s.Contains(1), true)
	assert.Eq(t, s.Contains(limit), true)
	assert.Eq(t, s.Empty(), false)

	for i := limit; i >= 1; i-- {
		assert.Eq(t, s.Count(), i)
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

	l.Remove(2)
	assert.Eq(t, l.Get(0), 2)
	assert.Eq(t, l.Get(1), 3)
	assert.Eq(t, l.Count(), 2)
}

func TestFailedRemove(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	q := New[string]()
	q.Remove(1)
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
