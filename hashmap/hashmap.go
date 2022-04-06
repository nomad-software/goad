package hashmap

import (
	"math"

	"github.com/nomad-software/goad/hash"
	"github.com/nomad-software/goad/linkedlist"
)

const (
	minBuckets = 16
	loadFactor = 0.75
)

// Payload is the main payload of the map.
type payload[K comparable, V comparable] struct {
	key K
	val V
}

// HashMap is the main map type.
type HashMap[K comparable, V comparable] struct {
	data     []linkedlist.LinkedList[payload[K, V]]
	capacity int
	count    int
}

// New is used to create a new map.
func New[K comparable, V comparable]() HashMap[K, V] {
	return HashMap[K, V]{
		data:     make([]linkedlist.LinkedList[payload[K, V]], minBuckets),
		capacity: minBuckets,
		count:    0,
	}
}

// Count returns the amount of entries in the map.
func (m HashMap[K, V]) Count() int {
	return m.count
}

// Empty returns true if the map is empty, false if not.
func (m HashMap[K, V]) Empty() bool {
	return m.Count() == 0
}

// Resize reallocates and resizes the underlying array to the passed number of
// buckets.
func (m *HashMap[K, V]) resize(cap int) {
	if cap < minBuckets {
		panic("can only resize greater or equal to the minimum bucket size")
	}

	data := m.data

	m.capacity = cap
	m.data = make([]linkedlist.LinkedList[payload[K, V]], m.capacity)
	m.count = 0

	for _, ln := range data {
		ln.ForEach(func(i int, p payload[K, V]) {
			m.Put(p.key, p.val)
		})
	}
}

// Put adds a value to the hash map relating to the passed key.
func (m *HashMap[K, V]) Put(key K, val V) {
	if m.count+1 >= int(float64(m.capacity)*loadFactor) {
		m.resize(m.capacity * 2)
	}

	hash := hash.Hash(key)
	bucket := int(math.Mod(float64(hash), float64(m.capacity)))

	var keyExists bool
	m.data[bucket].ForEach(func(i int, p payload[K, V]) {
		if p.key == key {
			keyExists = true
			m.data[bucket].Update(i, payload[K, V]{key: key, val: val})
			return
		}
	})

	if !keyExists {
		m.data[bucket].InsertLast(payload[K, V]{key: key, val: val})
		m.count++
	}
}

// Get gets a value from the hash map relating to the passed key.
func (m *HashMap[K, V]) Get(key K) (val V, ok bool) {
	hash := hash.Hash(key)
	bucket := int(math.Mod(float64(hash), float64(m.capacity)))

	m.data[bucket].ForEach(func(i int, p payload[K, V]) {
		if p.key == key {
			val = p.val
			ok = true
			return
		}
	})

	return val, ok
}

// Remove deletes a value from the hash map relating to the passed key.
func (m *HashMap[K, V]) Remove(key K) {
	hash := hash.Hash(key)
	bucket := int(math.Mod(float64(hash), float64(m.capacity)))

	m.data[bucket].ForEach(func(i int, p payload[K, V]) {
		if p.key == key {
			m.data[bucket].Remove(i)
			m.count--
			return
		}
	})

	if (m.capacity/2) >= minBuckets && m.count < int((float64(m.capacity)/2)*loadFactor) {
		m.resize(m.capacity / 2)
	}
}
