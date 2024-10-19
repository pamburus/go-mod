package gi2

import (
	"iter"
)

// Empty returns an iterator over pairs that yields nothing.
func Empty[V1, V2 any]() iter.Seq2[V1, V2] {
	return func(func(V1, V2) bool) {}
}
