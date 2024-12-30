package lili_test

import (
	"testing"

	"github.com/structx/lili"
)

var (
	bll *lili.LinkedList[int]
)

func init() {
	bll = lili.NewLinkedListInt()
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bll.Insert(i, []byte("helloworld"))
	}
}
