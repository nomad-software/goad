package hashmap

import (
	"testing"

	"github.com/nomad-software/goad/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	m := New[string, int]()

	m.Put("foo", 3)
	val, ok := m.Get("foo")
	assert.Eq(t, val, 3)
	assert.Eq(t, ok, true)

	m.Put("foo", 6)
	val, ok = m.Get("foo")
	assert.Eq(t, val, 6)
	assert.Eq(t, ok, true)

	assert.Eq(t, m.Count(), 1)

	m.Remove("foo")
	val, ok = m.Get("foo")
	assert.Eq(t, val, 0)
	assert.Eq(t, ok, false)

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

	m.Put("l", 12)
	assert.Eq(t, m.Count(), 12)
	assert.Eq(t, m.capacity, 32)

	m.Remove("l")
	assert.Eq(t, m.Count(), 11)
	assert.Eq(t, m.capacity, 16)
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
