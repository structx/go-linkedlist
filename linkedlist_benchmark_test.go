package linkedlist_test

import (
	"testing"

	"github.com/structx/go-linkedlist"
)

var (
	ll *linkedlist.LinkedList[int, []byte]
)

func init() {
	ll = &linkedlist.LinkedList[int, []byte]{}
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ll.Insert(i, []byte("helloworld"))
	}
}
