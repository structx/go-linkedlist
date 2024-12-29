## Lily

Generic Linked List with `comparable` types to be used as keys.

```go
func main() {
	ll := lily.NewLinkedListInt()
	ll.Insert(0, []byte("rick"))
	ll.Insert(1, []byte("morty"))
	ll.Insert(2, []byte("summer"))
	ll.Insert(3, []byte("beth"))
	ll.Insert(4, []byte("jerry"))

	value, err := ll.Search(0)
	if err != nil {
		if errors.Is(err, lily.ErrNotFound) {
			// record not found
		} else if errors.Is(err, lily.ErrNilHead) {
			// uninitialized linked list
		}
	}

 	fmt.Println(string(value)) // "rick"
}
```

### Benchmarks

builtin golang benchmark results

```
goos: linux
goarch: amd64
pkg: github.com/structx/lily
cpu: 13th Gen Intel(R) Core(TM) i5-1340P
BenchmarkInsert-16    	 6866752	       152.2 ns/op	      64 B/op	       2 allocs/op
PASS
ok  	github.com/structx/lily	1.239s
```
