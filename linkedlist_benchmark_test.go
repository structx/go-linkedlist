package lily_test

import (
	"testing"

	"github.com/structx/lily"
)

var (
	ll *lily.LinkedList[int]
)

func init() {
	ll = lily.NewLinkedListInt()
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ll.Insert(i, []byte("helloworld"))
	}
}
