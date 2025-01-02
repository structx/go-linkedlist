package lili

import (
	"sync/atomic"
	"unsafe"
)

type dnode[K comparable] struct {
	key      K
	payload  []byte
	next     *dnode[K]
	previous *dnode[K]
}

// DoubleLinkedList
type DoubleLinkedList[K comparable] struct {
	head *dnode[K]
	size atomic.Uintptr
}

// NewDoubleLinkedListInt
func NewDoubleLinkedListInt() *DoubleLinkedList[int] {
	return &DoubleLinkedList[int]{
		head: nil,
		size: atomic.Uintptr{},
	}
}

// Insert
func (dll *DoubleLinkedList[K]) Insert(key K, payload []byte) {

	nn := &dnode[K]{
		key:      key,
		payload:  payload,
		next:     nil,
		previous: nil,
	}
	s := unsafe.Sizeof(payload)

	if dll.head == nil {
		dll.head = nn
		dll.size.Store(s)
		return
	}

	n := dll.head

	for {

		if n.key == key {
			// TODO
			// update size
			n.payload = payload

			return
		}

		if n.next == nil {
			break
		}
		n = n.next
	}

	n.next = nn
	nn.previous = n

	dll.size.Add(s)
}

// Search
func (dll *DoubleLinkedList[K]) Search(key K) ([]byte, error) {

	if dll.head == nil {
		return nil, ErrNilRoot
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

	return nil, ErrNotFound
}

// Flush
func (dll *DoubleLinkedList[K]) Flush() {
	dll.head = nil
	dll.size.Store(0)
}

// Size
func (dll *DoubleLinkedList[K]) Size() uintptr {
	return dll.size.Load()
}
