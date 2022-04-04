package hash

import (
	"testing"

	"github.com/nomad-software/goad/assert"
)

func TestHashingBooleans(t *testing.T) {
	t.Parallel()

	assert.Eq(t, Hash(true), 7935342543426813952)
	assert.Eq(t, Hash(false), 8038925334856335360)
}

func TestHashingNumbers(t *testing.T) {
	t.Parallel()

	assert.Eq(t, Hash(123), 15852670688344145841)
	assert.Eq(t, Hash(int(123)), 15852670688344145841)
	assert.Eq(t, Hash(int8(123)), 2431943798780067840)
	assert.Eq(t, Hash(int16(123)), 8048425115320320000)
	assert.Eq(t, Hash(int32(123)), 8070450196166737920)
	assert.Eq(t, Hash(int64(123)), 15852670688344145841)
	assert.Eq(t, Hash(uint(123)), 15852670688344145841)
	assert.Eq(t, Hash(uint8(123)), 2431943798780067840)
	assert.Eq(t, Hash(uint16(123)), 8048425115320320000)
	assert.Eq(t, Hash(uint32(123)), 8070450196166737920)
	assert.Eq(t, Hash(uint64(123)), 15852670688344145841)
	assert.Eq(t, Hash(3.1415927), 15849285128134488456)
	assert.Eq(t, Hash(float32(3.1415927)), 261145898992533504)
	assert.Eq(t, Hash(float64(3.1415927)), 15849285128134488456)
	assert.Eq(t, Hash(complex(3, -5)), 8032196323694346239)
	assert.Eq(t, Hash(complex64(complex(3, -5))), 1734448804672045055)
	assert.Eq(t, Hash(complex128(complex(3, -5))), 8032196323694346239)
}

func TestHashingStrings(t *testing.T) {
	t.Parallel()

	assert.Eq(t, Hash("foo"), 4340397123594878976)
	assert.Eq(t, Hash("bar"), 3297786047190007808)
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

	assert.Eq(t, Hash(a1), 7561610947637318540)
	assert.Eq(t, Hash(a2), 7348474929389778593)
	assert.True(t, Hash(a3) > 0)
}

func TestHashingStructs(t *testing.T) {
	t.Parallel()

	type Foo struct {
		Foo string
		Bar string
	}

	assert.Eq(t, Hash(Foo{}), 17410835185711973985)
	assert.Eq(t, Hash(Foo{Foo: "foo", Bar: "bar"}), 4221156122319870358)
	assert.Eq(t, Hash(Foo{Foo: "baz", Bar: "qux"}), 10275651914583187888)
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
	type Foo struct {
		Foo string
		Bar string
	}

	b.ResetTimer()
	b.ReportAllocs()

	for x := 0; x < b.N; x++ {
		Hash(Foo{Foo: "foo", Bar: "bar"})
	}
}
