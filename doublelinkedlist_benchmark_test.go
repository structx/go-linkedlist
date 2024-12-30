package lili_test

import (
	"testing"

	"github.com/structx/lili"
)

var (
	dll *lili.DoubleLinkedList[int]
)

func init() {
	dll = lili.NewDoubleLinkedListInt()
}

func BenchmarkDoubleLinkedListInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dll.Insert(i, []byte("helloworld"))
	}
}
