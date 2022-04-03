package linkedlist

// Node is the node type used within the linked list.
type node[T comparable] struct {
	prev *node[T]
	next *node[T]
	val  T
}

// LinkedList is the main linked list type.
type LinkedList[T comparable] struct {
	first *node[T]
	last  *node[T]
	count int
}

// New is used to create a new linked list.
func New[T comparable]() LinkedList[T] {
	return LinkedList[T]{}
}

// Count returns the amount of entries in the linked list.
func (l *LinkedList[T]) Count() int {
	return l.count
}

// Empty returns true if the linked list is empty, false if not.
func (l *LinkedList[T]) Empty() bool {
	return l.Count() == 0
}

// InsertFirst inserts a value at the beginning of the linked list.
func (l *LinkedList[T]) InsertFirst(val T) {
	n := &node[T]{val: val}

	if l.first == nil {
		l.first = n
		l.last = n
	} else {
		l.first.prev = n
		n.next = l.first
		l.first = n
	}

	l.count++
}

// First returns the value at the beginning of the linked list.
func (l *LinkedList[T]) First() T {
	if l.first == nil {
		panic("linked list empty, getting first failed")
	}

	return l.first.val
}

// RemoveFirst removes the first value in the linked list.
func (l *LinkedList[T]) RemoveFirst() {
	if l.first != nil {
		if l.first.next == nil {
			l.first = nil
			l.last = nil
		} else {
			l.first = l.first.next
			l.first.prev = nil
		}
	}

	l.count--
}

// InsertLast inserts a value at the end of the linked list.
func (l *LinkedList[T]) InsertLast(val T) {
	n := &node[T]{val: val}

	if l.last == nil {
		l.first = n
		l.last = n
	} else {
		l.last.next = n
		n.prev = l.last
		l.last = n
	}

	l.count++
}

// Last returns the value at the end of the linked list.
func (l *LinkedList[T]) Last() T {
	if l.last == nil {
		panic("linked list empty, getting last failed")
	}

	return l.last.val
}

// RemoveLast removes the last value in the linked list.
func (l *LinkedList[T]) RemoveLast() {
	if l.last != nil {
		if l.last.prev == nil {
			l.first = nil
			l.last = nil
		} else {
			l.last = l.last.prev
			l.last.next = nil
		}
	}

	l.count--
}

// Insert inserts a value at the specified index.
func (l *LinkedList[T]) Insert(val T, index int) {
	if index > l.Count() {
		panic("Insertion index invalid")
	}

	if index == 0 {
		l.InsertFirst(val)

	} else if index == l.Count() {
		l.InsertLast(val)

	} else {
		n := &node[T]{val: val}

		var listIndex int = 0
		for ln := l.first; ln != nil; ln = ln.next {
			if listIndex == index {
				ln.prev.next = n
				n.prev = ln.prev
				n.next = ln
				ln.prev = n
				break
			}
			listIndex++
		}

		l.count++
	}
}

// Get gets a value at the specified index.
func (l *LinkedList[T]) Get(index int) T {
	if index >= l.Count() {
		panic("index outside of linked list bounds")
	}

	if index == 0 {
		return l.first.val

	} else if index == l.Count()-1 {
		return l.last.val

	} else {
		var listIndex int = 0
		for ln := l.first; ln != nil; ln = ln.next {
			if listIndex == index {
				return ln.val
			}
			listIndex++
		}
	}

	panic("unreachable")
}

// Update updates a value at the specified index.
func (l *LinkedList[T]) Update(val T, index int) {
	if index >= l.Count() {
		panic("index outside of linked list bounds")
	}

	var listIndex int = 0
	for ln := l.first; ln != nil; ln = ln.next {
		if listIndex == index {
			ln.val = val
			break
		}
		listIndex++
	}
}

// Remove removes a value at the specified index.
func (l *LinkedList[T]) Remove(index int) {
	if l.Count() == 0 {
		panic("linked list is empty")
	}

	if index >= l.Count() {
		panic("index outside of linked list bounds")
	}

	if index == 0 {
		l.RemoveFirst()

	} else if index == l.Count()-1 {
		l.RemoveLast()

	} else {
		var listIndex int = 0
		for ln := l.first; ln != nil; ln = ln.next {
			if listIndex == index {
				ln.prev.next = ln.next
				ln.next.prev = ln.prev
				break
			}
			listIndex++
		}

		l.count--
	}
}

// Contains returns true if the value exists in the linked list, false if not.
func (l *LinkedList[T]) Contains(val T) bool {
	for ln := l.first; ln != nil; ln = ln.next {
		if ln.val == val {
			return true
		}
	}

	return false
}

// Clear empties the entire linked list.
func (l *LinkedList[T]) Clear() {
	l.first = nil
	l.last = nil
	l.count = 0
}

// ForEach iterates over the dataset within the linked list, calling the passed
// function for each value.
func (l *LinkedList[T]) ForEach(f func(val T)) {
	for ln := l.first; ln != nil; ln = ln.next {
		f(ln.val)
	}
}
