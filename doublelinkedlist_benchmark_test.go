package linkedlist_test

import (
	"testing"

	"github.com/structx/go-linkedlist"
)

var (
	dll *linkedlist.DoubleLinkedList[int, []byte]
)

func init() {
	dll = &linkedlist.DoubleLinkedList[int, []byte]{}
}

func BenchmarkDoubleLinkedListInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dll.Insert(i, []byte("helloworld"))
	}
}
