package hash

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc64"
	"unsafe"
)

var isoTable = crc64.MakeTable(crc64.ISO)

// Hasher is an interface primarily for structs to provide a hash of their
// contents.
type Hasher interface {
	Hash() uint64
}

// HashBytes returns a 64bit unsigned integer hash of the passed byte slice.
func HashBytes(b []byte) uint64 {
	hash := crc64.New(isoTable)
	hash.Write(b)
	return hash.Sum64()
}

// Hash returns a 64bit unsigned integer hash for any value passed in.
func Hash[T comparable](val T) uint64 {
	hash := crc64.New(isoTable)
	buf := new(bytes.Buffer)

	switch v := any(val).(type) {
	case Hasher:
		return v.Hash()

	case int:
		binary.Write(buf, binary.LittleEndian, int64(v))
		hash.Write(buf.Bytes())

	case *int:
		binary.Write(buf, binary.LittleEndian, uint64(uintptr(unsafe.Pointer(v))))
		hash.Write(buf.Bytes())

	case uint:
		binary.Write(buf, binary.LittleEndian, uint64(v))
		hash.Write(buf.Bytes())

	case *uint:
		binary.Write(buf, binary.LittleEndian, uint64(uintptr(unsafe.Pointer(v))))
		hash.Write(buf.Bytes())

	case uintptr:
		binary.Write(buf, binary.LittleEndian, uint64(v))
		hash.Write(buf.Bytes())

	case string:
		hash.Write([]byte(v))

	case *string:
		binary.Write(buf, binary.LittleEndian, uint64(uintptr(unsafe.Pointer(v))))
		hash.Write(buf.Bytes())

	default:
		binary.Write(buf, binary.LittleEndian, v)
		hash.Write(buf.Bytes())
	}

	n := hash.Sum64()

	if n == 0 {
		hash.Write([]byte(fmt.Sprintf("%T:%v", val, val)))
		n = hash.Sum64()
	}

	return n
}
