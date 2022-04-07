package hash

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
	"unsafe"
)

// Hasher is an interface primarily for structs to provide a hash of their
// contents.
type Hasher interface {
	Hash() uint32
}

// HashBytes returns a 64bit unsigned integer hash of the passed byte slice.
func HashBytes(b []byte) uint32 {
	hash := fnv.New32a()
	hash.Write(b)
	return hash.Sum32()
}

// Hash returns a 64bit unsigned integer hash for any value passed in.
func Hash[T comparable](val T) uint32 {
	hash := fnv.New32a()
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

	return hash.Sum32()
}
