package lily

import (
	"sync/atomic"
	"unsafe"
)

type node[T comparable] struct {
	key     T
	payload []byte
	next    *node[T]
}

// LinkedList generic linked list with comparable types
type LinkedList[T comparable] struct {
	root *node[T]
	size atomic.Uintptr
}

// NewLinkedListInt constructor int type
func NewLinkedListInt() *LinkedList[int] {
	return &LinkedList[int]{
		root: nil,
		size: atomic.Uintptr{},
	}
}

// NewLinkedListInt64 constructor int64 type
func NewLinkedListInt64() *LinkedList[int64] {
	return &LinkedList[int64]{
		root: nil,
		size: atomic.Uintptr{},
	}
}

// NewLinkedListString constructor string type
func NewLinkedListString() *LinkedList[string] {
	return &LinkedList[string]{
		root: nil,
		size: atomic.Uintptr{},
	}
}

// Iterator
type Iterator[T comparable] struct {
	node *node[T]
}

// Insert new node into list
func (ll *LinkedList[T]) Insert(key T, payload []byte) {

	nn := &node[T]{
		key:     key,
		payload: payload,
		next:    nil,
	}

	if ll.root == nil {
		ll.root = nn
		s := unsafe.Sizeof(nn.payload)
		ll.size.Store(s)
		return
	}

	n := ll.root
	for {
		if n.next == nil {
			break
		}
		n = n.next
	}

	n.next = nn
	s := unsafe.Sizeof(nn.payload)
	ll.size.Add(s)
}

// Search list for key and return payload
func (ll *LinkedList[T]) Search(key T) ([]byte, error) {

	if ll.root == nil {
		return nil, ErrNilRoot
	}

	n := ll.root
	for n != nil {
		if equal(n.key, key) {
			return n.payload, nil
		}
		n = n.next
	}

	return nil, ErrNotFound
}

// Flush nil root and reset size
func (ll *LinkedList[T]) Flush() {
	ll.root = nil
	ll.size.Store(0)
}

// Size returns size of list
func (ll *LinkedList[T]) Size() uintptr {
	return ll.size.Load()
}

// HasNext returns conditional if more nodes
func (it *Iterator[T]) HasNext() bool {
	return it.node != nil || it.node.next != nil
}

// Next returns next node payload
func (it *Iterator[T]) Next() []byte {
	p := it.node.payload
	it.node = it.node.next
	return p
}

func equal[T comparable](a, b T) bool {
	return a == b
}
