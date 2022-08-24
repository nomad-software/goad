package binaryheap

import (
	"golang.org/x/exp/slices"
)

// BinaryHeap is the main heap type.
type BinaryHeap[T comparable] struct {
	data   []T
	pred   func(a T, b T) bool
	sorted bool
}

// New is used to create a new heap.
// The passed function is a predicate that returns true if the first parameter
// is greater than the second. This predicate defines the sorting order between
// the heap items and is called during insertion and extraction.
func New[T comparable](pred func(a T, b T) bool) BinaryHeap[T] {
	return BinaryHeap[T]{
		data: make([]T, 0, 16),
		pred: pred,
	}
}

// Count returns the amount of entries in the heap.
func (b BinaryHeap[T]) Count() int {
	return len(b.data)
}

// Empty returns true if the heap is empty, false if not.
func (b BinaryHeap[T]) Empty() bool {
	return b.Count() == 0
}

// SiftUp sifts the value at the passed index up through the heap until it finds
// its correct position.
func (b *BinaryHeap[T]) siftUp(childIndex int) {
	var parent T
	var child T
	var parentIndex int

	if childIndex > 0 {
		if childIndex > 2 {
			if childIndex%2 == 0 {
				parentIndex = (childIndex - 2) / 2
			} else {
				parentIndex = (childIndex - 1) / 2
			}
		} else {
			parentIndex = 0
		}

		parent = b.data[parentIndex]
		child = b.data[childIndex]

		if b.pred(child, parent) {
			b.data[parentIndex] = child
			b.data[childIndex] = parent

			if parentIndex > 0 {
				b.siftUp(parentIndex)
			}
		}
	}
}

// SiftDown sifts the value at the passed index down through the heap until it finds
// its correct position.
func (b *BinaryHeap[T]) siftDown(parentIndex int) {
	var parent T
	var child1 T
	var child2 T
	var child1Index int
	var child2Index int

	child1Index = (2 * parentIndex) + 1
	child2Index = (2 * parentIndex) + 2

	if b.Count() <= child1Index { // The parent has no children.
		return

	} else if b.Count() == child2Index { // The parent has one child.
		parent = b.data[parentIndex]
		child1 = b.data[child1Index]

		if b.pred(child1, parent) {
			b.data[parentIndex] = child1
			b.data[child1Index] = parent
			b.siftDown(child1Index)
		}

	} else { // The parent has two children.
		parent = b.data[parentIndex]
		child1 = b.data[child1Index]
		child2 = b.data[child2Index]

		// Compare the parent against the greater child.
		if b.pred(child1, child2) {
			if b.pred(child1, parent) {
				b.data[parentIndex] = child1
				b.data[child1Index] = parent
				b.siftDown(child1Index)
			}
		} else {
			if b.pred(child2, parent) {
				b.data[parentIndex] = child2
				b.data[child2Index] = parent
				b.siftDown(child2Index)
			}
		}
	}
}

// Insert inserts a new value into the heap.
func (b *BinaryHeap[T]) Insert(val T) {
	b.data = append(b.data, val)
	b.siftUp(b.Count() - 1)
	b.sorted = false
}

// Peek returns the first value at the top of the heap.
func (b BinaryHeap[T]) Peek() T {
	return b.data[0]
}

// Extract returns and removes the first value from the heap.
func (b *BinaryHeap[T]) Extract() T {
	if b.Empty() {
		panic("binary heap empty, extracting failed")
	}
	val := b.data[0]
	b.data[0] = b.data[b.Count()-1]
	b.data = b.data[0 : b.Count()-1]
	b.siftDown(0)
	b.sorted = false
	return val
}

// Contains returns true if the value exists in the heap, false if not.
func (b BinaryHeap[T]) Contains(val T) bool {
	for _, v := range b.data {
		if v == val {
			return true
		}
	}
	return false
}

// Clear empties the entire heap.
func (b *BinaryHeap[T]) Clear() {
	b.data = b.data[:0:0]
}

// Sort the heap ready for iterating.
// The heap needs to be sorted to be iterated correctly using a loop. This is to
// make sure values are delivered in the correct order. Sorted or unsorted, the
// heap data structure will always be correct because of the algorithms used
// when sifting.
func (b *BinaryHeap[T]) sort() {
	if !b.sorted {
		slices.SortFunc(b.data, b.pred)
		b.sorted = true
	}
}

// ForEach iterates over the dataset within the heap, calling the passed
// function for each value.
func (b BinaryHeap[T]) ForEach(f func(val T)) {
	b.sort()
	for _, v := range b.data {
		f(v)
	}
}
