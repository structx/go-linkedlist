package linkedlist

import (
	"sync/atomic"
	"unsafe"
)

// Node
type Node[K comparable, T any] struct {
	Key     K
	Payload T
	next    *Node[K, T]
}

// LinkedList generic linked list with comparable types
type LinkedList[K comparable, T any] struct {
	root *Node[K, T]
	size atomic.Uintptr
}

// Iterator
type Iterator[K comparable, T any] struct {
	node *Node[K, T]
}

// Iterator
func (ll *LinkedList[K, T]) Iterator() *Iterator[K, T] {
	return &Iterator[K, T]{
		node: ll.root,
	}
}

// Insert new node into list
func (ll *LinkedList[K, T]) Insert(key K, payload T) *Node[K, T] {
	defer ll.size.Add(unsafe.Sizeof(payload))

	nn := &Node[K, T]{
		Key:     key,
		Payload: payload,
		next:    nil,
	}

	if ll.root == nil {
		ll.root = nn
		return nn
	}

	n := ll.root

	for {
		if n.next == nil {
			break
		}
		n = n.next
	}

	n.next = nn

	return nn
}

// Search list for key and return payload
func (ll *LinkedList[K, T]) Search(key K) (any, error) {

	if ll.root == nil {
		return nil, ErrNilRoot
	}

	n := ll.root
	for n != nil {
		if n.Key == key {
			return n.Payload, nil
		}
		n = n.next
	}

	return nil, ErrNotFound
}

// Flush nil root and reset size
func (ll *LinkedList[K, T]) Flush() {
	ll.root = nil
	ll.size.Store(0)
}

// Size returns size of list
func (ll *LinkedList[K, T]) Size() uintptr {
	return ll.size.Load()
}

// HasNext returns conditional if more nodes
func (it *Iterator[K, T]) HasNext() bool {
	return it.node != nil || it.node.next != nil
}

// Next returns next node payload
func (it *Iterator[K, T]) Next() any {
	p := it.node.Payload
	it.node = it.node.next
	return p
}
