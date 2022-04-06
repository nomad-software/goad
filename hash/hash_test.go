package hash

import (
	"bytes"
	"testing"

	"github.com/nomad-software/goad/assert"
)

func TestHashingBooleans(t *testing.T) {
	t.Parallel()

	assert.True(t, Hash(true) > 0)
	assert.True(t, Hash(false) > 0)
}

func TestHashingNumbers(t *testing.T) {
	t.Parallel()

	assert.True(t, Hash(123) > 0)
	assert.True(t, Hash(int(123)) > 0)
	assert.True(t, Hash(int8(123)) > 0)
	assert.True(t, Hash(int16(123)) > 0)
	assert.True(t, Hash(int32(123)) > 0)
	assert.True(t, Hash(int64(123)) > 0)
	assert.True(t, Hash(uint(123)) > 0)
	assert.True(t, Hash(uint8(123)) > 0)
	assert.True(t, Hash(uint16(123)) > 0)
	assert.True(t, Hash(uint32(123)) > 0)
	assert.True(t, Hash(uint64(123)) > 0)
	assert.True(t, Hash(3.1415927) > 0)
	assert.True(t, Hash(float32(3.1415927)) > 0)
	assert.True(t, Hash(float64(3.1415927)) > 0)
	assert.True(t, Hash(complex(3, -5)) > 0)
	assert.True(t, Hash(complex64(complex(3, -5))) > 0)
	assert.True(t, Hash(complex128(complex(3, -5))) > 0)
}

func TestHashingStrings(t *testing.T) {
	t.Parallel()

	assert.True(t, Hash("foo") > 0)
	assert.True(t, Hash("bar") > 0)
}

func TestHashingPointers(t *testing.T) {
	t.Parallel()

	p1 := true
	p2 := false
	p3 := 246
	p4 := int(123)
	p5 := int8(123)
	p6 := int16(123)
	p7 := int32(123)
	p8 := int64(123)
	p9 := uint(123)
	p10 := uint8(123)
	p11 := uint16(123)
	p12 := uint32(123)
	p13 := uint64(123)
	p14 := "foo bar baz qux"
	p15 := uintptr(0xDEADBEEF)

	assert.True(t, Hash(&p1) > 0)
	assert.True(t, Hash(&p2) > 0)
	assert.True(t, Hash(&p3) > 0)
	assert.True(t, Hash(&p4) > 0)
	assert.True(t, Hash(&p5) > 0)
	assert.True(t, Hash(&p6) > 0)
	assert.True(t, Hash(&p7) > 0)
	assert.True(t, Hash(&p8) > 0)
	assert.True(t, Hash(&p9) > 0)
	assert.True(t, Hash(&p10) > 0)
	assert.True(t, Hash(&p11) > 0)
	assert.True(t, Hash(&p12) > 0)
	assert.True(t, Hash(&p13) > 0)
	assert.True(t, Hash(&p14) > 0)
	assert.True(t, Hash(p15) > 0)
}

func TestHashingChannels(t *testing.T) {
	t.Parallel()

	c1 := make(chan int)
	c2 := c1
	c3 := make(chan string, 10)
	c4 := make(<-chan float64, 10)
	c5 := make(chan<- struct{}, 10)

	assert.True(t, Hash(c1) > 0)
	assert.True(t, Hash(&c1) > 0)
	assert.True(t, Hash(c2) > 0)
	assert.True(t, Hash(&c2) > 0)
	assert.True(t, Hash(c3) > 0)
	assert.True(t, Hash(&c3) > 0)
	assert.True(t, Hash(c4) > 0)
	assert.True(t, Hash(&c4) > 0)
	assert.True(t, Hash(c5) > 0)
	assert.True(t, Hash(&c5) > 0)
}

func TestHashingArrays(t *testing.T) {
	t.Parallel()

	p1 := 123
	a1 := [3]int{1, 2, 3}
	a2 := [4]string{"foo", "bar", "baz", "qux"}
	a3 := [3]*int{&p1, &p1, &p1}

	assert.True(t, Hash(a1) > 0)
	assert.True(t, Hash(a2) > 0)
	assert.True(t, Hash(a3) > 0)
}

type foo struct {
	foo string
	bar string
}

func TestHashingStructs(t *testing.T) {
	t.Parallel()

	assert.True(t, Hash(foo{}) > 0)
	assert.True(t, Hash(foo{foo: "foo", bar: "bar"}) > 0)
	assert.True(t, Hash(foo{foo: "baz", bar: "qux"}) > 0)
}

type baz struct {
	baz string
	qux string
}

func (h baz) Hash() uint32 {
	buf := new(bytes.Buffer)
	buf.Write([]byte(h.baz))
	buf.Write([]byte(h.qux))
	return HashBytes(buf.Bytes())
}

func TestHashingHashers(t *testing.T) {
	t.Parallel()

	assert.True(t, Hash(baz{}) > 0)
	assert.True(t, Hash(baz{baz: "foo", qux: "bar"}) > 0)
	assert.True(t, Hash(baz{baz: "baz", qux: "qux"}) > 0)
}

func TestHashBytes(t *testing.T) {
	t.Parallel()

	assert.True(t, HashBytes([]byte{1, 2, 3, 4, 5}) > 0)
	assert.True(t, HashBytes([]byte{6, 7, 8, 9, 10}) > 0)
}

func BenchmarkHashingStrings(b *testing.B) {
	b.ReportAllocs()

	for x := 0; x < b.N; x++ {
		Hash("foo bar baz qux")
	}
}

func BenchmarkHashingIntegers(b *testing.B) {
	b.ReportAllocs()

	for x := 0; x < b.N; x++ {
		Hash(x)
	}
}

func BenchmarkHashingStructs(b *testing.B) {
	b.ReportAllocs()

	for x := 0; x < b.N; x++ {
		Hash(foo{foo: "foo", bar: "bar"})
	}
}

func BenchmarkHashingHashers(b *testing.B) {
	b.ReportAllocs()

	for x := 0; x < b.N; x++ {
		Hash(baz{baz: "baz", qux: "qux"})
	}
}
