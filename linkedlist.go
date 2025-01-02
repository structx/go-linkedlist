package lili

import (
	"sync/atomic"
	"unsafe"
)

type node[K comparable] struct {
	key     K
	payload []byte
	next    *node[K]
}

// LinkedList generic linked list with comparable types
type LinkedList[K comparable] struct {
	root *node[K]
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
type Iterator[K comparable] struct {
	node *node[K]
}

// Iterator
func (ll *LinkedList[K]) Iterator() *Iterator[K] {
	return &Iterator[K]{
		node: ll.root,
	}
}

// Insert new node into list
func (ll *LinkedList[K]) Insert(key K, payload []byte) {

	nn := &node[K]{
		key:     key,
		payload: payload,
		next:    nil,
	}
	s := unsafe.Sizeof(nn.payload)

	if ll.root == nil {
		ll.root = nn
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
	ll.size.Add(s)
}

// Search list for key and return payload
func (ll *LinkedList[K]) Search(key K) (any, error) {

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
func (ll *LinkedList[K]) Flush() {
	ll.root = nil
	ll.size.Store(0)
}

// Size returns size of list
func (ll *LinkedList[K]) Size() uintptr {
	return ll.size.Load()
}

// HasNext returns conditional if more nodes
func (it *Iterator[K]) HasNext() bool {
	return it.node != nil || it.node.next != nil
}

// Next returns next node payload
func (it *Iterator[K]) Next() any {
	p := it.node.payload
	it.node = it.node.next
	return p
}
