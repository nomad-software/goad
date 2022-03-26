package linkedlist

import (
	"fmt"

	"github.com/nomad-software/goad/constraint"
)

type node[T constraint.BuiltinTypes] struct {
	prev *node[T]
	next *node[T]
	val  T
}

type LinkedList[T constraint.BuiltinTypes] struct {
	first *node[T]
	last  *node[T]
	count int
}

func New[T constraint.BuiltinTypes]() LinkedList[T] {
	return LinkedList[T]{}
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) Empty() bool {
	return l.Count() == 0
}

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

func (l *LinkedList[T]) First() T {
	if l.first == nil {
		panic("linked list empty, getting first failed")
	}

	return l.first.val
}

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

func (l *LinkedList[T]) Last() T {
	if l.last == nil {
		panic("linked list empty, getting last failed")
	}

	return l.last.val
}

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

func (l *LinkedList[T]) Insert(val T, index int) {
	if index < 0 || index > l.Count() {
		panic("Insertion index invalid")
	}

	if index == 0 {
		l.InsertFirst(val)

	} else if index == l.Count() {
		l.InsertLast(val)

	} else {
		n := &node[T]{val: val}

		listIndex := 0
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

func (l *LinkedList[T]) Get(index int) T {
	if index < 0 || index >= l.Count() {
		panic("Index outside of linked list bounds")
	}

	if index == 0 {
		return l.first.val

	} else if index == l.Count()-1 {
		return l.last.val

	} else {
		listIndex := 0
		for ln := l.first; ln != nil; ln = ln.next {
			if listIndex == index {
				return ln.val
			}
			listIndex++
		}
	}

	panic("Index outside of linked list bounds")
}

func (l *LinkedList[T]) Update(val T, index int) {
	if index < 0 || index >= l.Count() {
		panic("Index outside of linked list bounds")
	}

	listIndex := 0
	for ln := l.first; ln != nil; ln = ln.next {
		if listIndex == index {
			ln.val = val
			break
		}
		listIndex++
	}
}

func (l *LinkedList[T]) Remove(index int) {
	if l.Count() == 0 {
		panic("Linked list is empty")
	}

	if index < 0 || index >= l.Count() {
		panic("Index outside of linked list bounds")
	}

	if index == 0 {
		l.RemoveFirst()

	} else if index == l.Count()-1 {
		l.RemoveLast()

	} else {
		listIndex := 0
		for ln := l.first; ln != nil; ln = ln.next {
			if listIndex == index {
				fmt.Printf("index: %v\n", listIndex)
				ln.prev.next = ln.next
				ln.next.prev = ln.prev
				break
			}
			listIndex++
		}

		l.count--
	}
}

func (l *LinkedList[T]) Contains(val T) bool {
	for ln := l.first; ln != nil; ln = ln.next {
		if ln.val == val {
			return true
		}
	}

	return false
}

func (l *LinkedList[T]) Clear() {
	l.first = nil
	l.last = nil
	l.count = 0
}

func (l *LinkedList[T]) Range(f func(val T)) {
	for ln := l.first; ln != nil; ln = ln.next {
		f(ln.val)
	}
}
