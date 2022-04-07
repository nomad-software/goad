package hash

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
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

	default:
		binary.Write(buf, binary.LittleEndian, v)
		hash.Write(buf.Bytes())
	}

	return hash.Sum32()
}
