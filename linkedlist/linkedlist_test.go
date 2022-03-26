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

func TestRange(t *testing.T) {
	t.Parallel()

	l := New[int]()

	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)
	l.InsertLast(4)
	l.InsertLast(5)

	i := 1
	l.Range(func(val int) {
		assert.Eq(t, val, i)
		i++
	})
}
