package linkedlist

import (
	"sync/atomic"
	"unsafe"
)

func zeroGN[T any]() T {
	var t T
	return t
}

// DNode
type DNode[K comparable, T any] struct {
	key      K
	payload  T
	next     *DNode[K, T]
	previous *DNode[K, T]
}

// DoubleLinkedList
type DoubleLinkedList[K comparable, T any] struct {
	head *DNode[K, T]
	size atomic.Uintptr
}

// Insert
func (dll *DoubleLinkedList[K, T]) Insert(key K, payload T) *DNode[K, T] {
	defer dll.size.Add(unsafe.Sizeof(payload))

	nn := &DNode[K, T]{
		key:      key,
		payload:  payload,
		next:     nil,
		previous: nil,
	}

	if dll.head == nil {
		dll.head = nn
		return nn
	}

	n := dll.head

	for {

		if n.key == key {
			// TODO
			// update size
			n.payload = payload

			return nn
		}

		if n.next == nil {
			break
		}
		n = n.next
	}

	n.next = nn
	nn.previous = n

	return nn
}

// Search
func (dll *DoubleLinkedList[K, T]) Search(key K) (T, error) {

	if dll.head == nil {
		return zeroGN[T](), ErrNilRoot
	}

	n := dll.head
	for {

		if n.key == key {
			return n.payload, nil
		}

		if n.next == nil {
			break
		}

		n = n.next
	}

	return zeroGN[T](), ErrNotFound
}

// Flush
func (dll *DoubleLinkedList[K, T]) Flush() {
	dll.head = nil
	dll.size.Store(0)
}

// Size
func (dll *DoubleLinkedList[K, T]) Size() uintptr {
	return dll.size.Load()
}
