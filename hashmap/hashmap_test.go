package hashmap

import (
	"testing"

	"github.com/nomad-software/goad/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	m := New[string, int]()
	assert.True(t, m.Empty())

	m.Put("foo", 3)
	val, ok := m.Get("foo")
	assert.Eq(t, val, 3)
	assert.True(t, ok)

	m.Put("foo", 6)
	val, ok = m.Get("foo")
	assert.Eq(t, val, 6)
	assert.True(t, ok)

	assert.Eq(t, m.Count(), 1)
	assert.False(t, m.Empty())

	m.Remove("foo")
	val, ok = m.Get("foo")
	assert.Eq(t, val, 0)
	assert.False(t, ok)

	m.Remove("foo")
	m.Remove("foo")

	assert.Eq(t, m.Count(), 0)
}

func TestResizing(t *testing.T) {
	t.Parallel()

	m := New[string, int]()
	assert.Eq(t, m.capacity, 16)

	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)
	m.Put("d", 4)
	m.Put("e", 5)
	m.Put("f", 6)
	m.Put("g", 7)
	m.Put("h", 8)
	m.Put("i", 9)
	m.Put("j", 10)
	m.Put("k", 11)
	assert.Eq(t, m.Count(), 11)
	assert.Eq(t, m.capacity, 16)
	assert.False(t, m.Empty())

	m.Put("l", 12)
	assert.Eq(t, m.Count(), 12)
	assert.Eq(t, m.capacity, 32)

	m.Remove("l")
	assert.Eq(t, m.Count(), 11)
	assert.Eq(t, m.capacity, 16)
}

func TestFailedResize(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("no panic detected")
		}
	}()

	m := New[string, int]()
	m.resize(2)
}

func TestContains(t *testing.T) {
	t.Parallel()

	m := New[string, int]()
	assert.Eq(t, m.capacity, 16)

	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)
	m.Put("d", 4)
	m.Put("e", 5)
	assert.False(t, m.Empty())

	assert.True(t, m.ContainsKey("a"))
	assert.False(t, m.ContainsKey("f"))

	assert.True(t, m.ContainsValue(3))
	assert.False(t, m.ContainsValue(10))
}

func TestClearing(t *testing.T) {
	t.Parallel()

	m := New[string, string]()
	assert.Eq(t, m.capacity, 16)

	m.Put("a", "1")
	m.Put("b", "2")
	m.Put("c", "3")
	m.Put("d", "4")
	m.Put("e", "5")
	m.Put("f", "6")
	m.Put("g", "7")
	m.Put("h", "8")
	m.Put("i", "9")
	m.Put("j", "10")
	m.Put("k", "11")
	m.Put("l", "12")
	assert.Eq(t, m.Count(), 12)
	assert.Eq(t, m.capacity, 32)

	m.Clear()
	assert.Eq(t, m.Count(), 0)
	assert.Eq(t, m.capacity, 16)
}

func TestForEach(t *testing.T) {
	t.Parallel()

	m := New[string, int]()

	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)
	m.Put("d", 4)
	m.Put("e", 5)

	assert.Eq(t, m.Count(), 5)

	var i int
	m.ForEach(func(key string, val int) {
		i++
	})

	assert.Eq(t, i, 5)

	m.Clear()
	m.ForEach(func(key string, val int) {
		t.Errorf("hashmap not cleared")
	})
}

func TestKeysAndValues(t *testing.T) {
	m := New[string, int]()

	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)
	m.Put("d", 4)
	m.Put("e", 5)

	var i int
	for range m.Keys() {
		i++
	}
	assert.Eq(t, i, 5)

	i = 0
	for range m.Values() {
		i++
	}
	assert.Eq(t, i, 5)
}

func BenchmarkPut(b *testing.B) {
	m := New[string, int]()

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		m.Put("foo", x)
	}
}

func BenchmarkGet(b *testing.B) {
	m := New[string, int]()
	m.Put("foo", 1337)

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		m.Get("foo")
	}
}
