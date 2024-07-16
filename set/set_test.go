package set

import (
	"strconv"
	"testing"

	"github.com/nomad-software/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	s := New[int]()
	s.Add(1)
	s.Add(1)
	s.Add(2)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	assert.False(t, s.Empty())
	assert.Eq(t, s.Count(), 3)
	s.Remove(1)
	assert.Eq(t, s.Count(), 2)
	s.Remove(2)
	s.Remove(3)
	assert.True(t, s.Empty())
}

func TestContains(t *testing.T) {
	t.Parallel()

	s := New[string]()
	s.Add("foo")
	s.Add("bar")
	s.Add("baz")
	s.Add("qux")

	assert.True(t, s.Contains("bar"))
	assert.False(t, s.Contains("fuz"))
}

func TestClearing(t *testing.T) {
	t.Parallel()

	s := New[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	assert.Eq(t, s.Count(), 3)

	s.Clear()
	assert.Eq(t, s.Count(), 0)

	s.Add(1)
	s.Add(2)
	assert.Eq(t, s.Count(), 2)
}

func TestForEach(t *testing.T) {
	t.Parallel()

	s := New[int]()

	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s.Add(5)

	var i int
	s.ForEach(func(val int) {
		i++
	})

	assert.Eq(t, i, 5)

	s.Clear()
	s.ForEach(func(val int) {
		t.Errorf("set not cleared")
	})
}

func BenchmarkSetAdd(b *testing.B) {
	m := New[string]()

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		m.Add("foo")
	}
}

func BenchmarkSetRemove(b *testing.B) {
	m := New[string]()

	for x := 0; x < 1_000_000; x++ {
		m.Add(strconv.Itoa(x))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		m.Remove("500000")
	}
}

func BenchmarkSetForEach(b *testing.B) {
	m := New[int]()

	for x := 0; x < 1_000_000; x++ {
		m.Add(x)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for x := 0; x < b.N; x++ {
		m.ForEach(func(val int) {})
	}
}
