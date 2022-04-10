package binaryheap

import (
	"testing"

	"github.com/nomad-software/goad/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	b := New(func(a, b int) bool { return a < b })
	b.Insert(6)
	b.Insert(5)
	b.Insert(2)
	b.Insert(9)
	b.Insert(4)
	b.Insert(8)
	b.Insert(7)
	b.Insert(1)
	b.Insert(3)
	b.Insert(10)

	assert.False(t, b.Empty())
	assert.Eq(t, b.Count(), 10)

	for i := 1; i <= 10; i++ {
		assert.Eq(t, b.Peek(), i)
		assert.Eq(t, b.Extract(), i)
	}

	assert.True(t, b.Empty())
	assert.Eq(t, b.Count(), 0)
}

func TestStruct(t *testing.T) {
	t.Parallel()

	type Foo struct {
		Foo int
		Bar string
	}

	b := New(func(a, b Foo) bool { return a.Foo < b.Foo })

	f1 := Foo{Foo: 2, Bar: "bar"}
	f2 := Foo{Foo: 4, Bar: "qux"}
	f3 := Foo{Foo: 3, Bar: "baz"}
	f4 := Foo{Foo: 1, Bar: "foo"}

	b.Insert(f1)
	b.Insert(f2)
	b.Insert(f3)
	b.Insert(f4)

	assert.False(t, b.Empty())
	assert.Eq(t, b.Count(), 4)

	assert.True(t, b.Contains(f1))
	assert.True(t, b.Contains(f2))
	assert.True(t, b.Contains(f3))
	assert.True(t, b.Contains(f4))

	assert.Eq(t, b.Extract(), f4)
	assert.Eq(t, b.Extract(), f1)
	assert.Eq(t, b.Extract(), f3)
	assert.Eq(t, b.Extract(), f2)

	assert.True(t, b.Empty())
}

func TestLargeCapacity(t *testing.T) {
	t.Parallel()

	b := New(func(a, b int) bool { return a < b })
	limit := 1_000_000

	for i := 1; i <= limit; i++ {
		b.Insert(i)
		assert.Eq(t, b.Peek(), 1)
		assert.Eq(t, b.Count(), i)
	}

	assert.Eq(t, b.Peek(), 1)
	assert.Eq(t, b.Count(), limit)
	assert.True(t, b.Contains(1))
	assert.True(t, b.Contains(limit))
	assert.False(t, b.Empty())

	c := b.Count()
	for i := 1; i <= limit; i++ {
		assert.Eq(t, b.Count(), c)
		assert.Eq(t, b.Peek(), i)
		assert.Eq(t, b.Extract(), i)
		c--
	}

	assert.True(t, b.Empty())
	assert.Eq(t, b.Count(), 0)
}

func TestFailedExtract(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	b := New(func(a, b int) bool { return a < b })
	b.Extract()
}

func TestContains(t *testing.T) {
	t.Parallel()

	b := New(func(a, b string) bool { return a < b })
	b.Insert("foo")
	b.Insert("bar")
	b.Insert("baz")
	b.Insert("qux")

	assert.True(t, b.Contains("bar"))
	assert.False(t, b.Contains("fuz"))
}

func TestClearing(t *testing.T) {
	t.Parallel()

	b := New(func(a, b string) bool { return a < b })
	b.Insert("foo")
	b.Insert("bar")
	b.Insert("baz")
	b.Insert("qux")
	assert.Eq(t, b.Count(), 4)

	b.Clear()
	assert.Eq(t, b.Count(), 0)

	b.Insert("foo")
	b.Insert("bar")
	assert.Eq(t, b.Count(), 2)
}

func TestForEach(t *testing.T) {
	t.Parallel()

	b := New(func(a, b int) bool { return a < b })
	b.Insert(6)
	b.Insert(5)
	b.Insert(2)
	b.Insert(9)
	b.Insert(4)
	b.Insert(8)
	b.Insert(7)
	b.Insert(1)
	b.Insert(3)
	b.Insert(10)

	assert.False(t, b.Empty())
	assert.Eq(t, b.Count(), 10)

	i := 1
	b.ForEach(func(val int) {
		assert.Eq(t, val, i)
		i++
	})

	assert.False(t, b.Empty())
	assert.Eq(t, b.Count(), 10)

	b.Clear()
	b.ForEach(func(val int) {
		t.Errorf("stack not cleared")
	})
}

func BenchmarkBinaryHeap(b *testing.B) {
	h := New(func(a, b int) bool { return a < b })

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		h.Insert(x)
	}
}

func BenchmarkForEach(b *testing.B) {
	h := New(func(a, b int) bool { return a < b })

	for x := 0; x < 1_000_000; x++ {
		h.Insert(x)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		h.ForEach(func(val int) {
		})
	}
}
