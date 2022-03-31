package hash

import (
	"bytes"
	"encoding/binary"
	"hash/crc64"

	"github.com/nomad-software/goad/constraint"
)

var isoTable = crc64.MakeTable(crc64.ISO)

// Hash returns a 64bit unsigned integer hash for any value passed in.
func Hash[T constraint.BuiltinTypes](val T) uint64 {
	hash := crc64.New(isoTable)
	buf := new(bytes.Buffer)

	switch v := any(val).(type) {
	case int:
		binary.Write(buf, binary.LittleEndian, int64(v))
		hash.Write(buf.Bytes())

	case uint:
		binary.Write(buf, binary.LittleEndian, uint64(v))
		hash.Write(buf.Bytes())

	case uintptr:
		binary.Write(buf, binary.LittleEndian, uint64(v))
		hash.Write(buf.Bytes())

	case string:
		hash.Write([]byte(v))

	default:
		binary.Write(buf, binary.LittleEndian, v)
		hash.Write(buf.Bytes())

	}

	return hash.Sum64()
}
