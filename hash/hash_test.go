package hash

import (
	"testing"

	"github.com/nomad-software/goad/assert"
)

func TestHashSignedInteger(t *testing.T) {
	t.Parallel()

	assert.Eq(t, Hash(123), 15852670688344145841)
	assert.Eq(t, Hash(int(123)), 15852670688344145841)
	assert.Eq(t, Hash(int8(123)), 2431943798780067840)
	assert.Eq(t, Hash(int16(123)), 8048425115320320000)
	assert.Eq(t, Hash(int32(123)), 8070450196166737920)
	assert.Eq(t, Hash(int64(123)), 15852670688344145841)
}

func TestHashUnsignedInteger(t *testing.T) {
	t.Parallel()

	assert.Eq(t, Hash(uint(123)), 15852670688344145841)
	assert.Eq(t, Hash(uint8(123)), 2431943798780067840)
	assert.Eq(t, Hash(uint16(123)), 8048425115320320000)
	assert.Eq(t, Hash(uint32(123)), 8070450196166737920)
	assert.Eq(t, Hash(uint64(123)), 15852670688344145841)
}

func TestHashUintptr(t *testing.T) {
	t.Parallel()

	assert.Eq(t, Hash(uintptr(0xDEADBEEF)), 8070450529433285204)
}

func TestHashFloat(t *testing.T) {
	t.Parallel()

	assert.Eq(t, Hash(3.1415927), 15849285128134488456)
	assert.Eq(t, Hash(float32(3.1415927)), 261145898992533504)
	assert.Eq(t, Hash(float64(3.1415927)), 15849285128134488456)
}

func TestHashComplex(t *testing.T) {
	t.Parallel()

	assert.Eq(t, Hash(complex(3, -5)), 8032196323694346239)
	assert.Eq(t, Hash(complex64(complex(3, -5))), 1734448804672045055)
	assert.Eq(t, Hash(complex128(complex(3, -5))), 8032196323694346239)
}

func TestHashString(t *testing.T) {
	t.Parallel()

	assert.Eq(t, Hash("foo"), 4340397123594878976)
	assert.Eq(t, Hash("bar"), 3297786047190007808)
}

func TestHashBool(t *testing.T) {
	t.Parallel()

	assert.Eq(t, Hash(true), 7935342543426813952)
	assert.Eq(t, Hash(false), 8038925334856335360)
}

func BenchmarkHash(b *testing.B) {
	b.ReportAllocs()

	for x := 0; x < b.N; x++ {
		Hash(x)
	}
}
