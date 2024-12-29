package lily

import "errors"

var (
	// ErrNilRoot
	ErrNilRoot = errors.New("nil root node")
	// ErrNotFound
	ErrNotFound = errors.New("key not found")
)
