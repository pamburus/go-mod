package gi

import (
	"iter"
)

// Empty returns an iterator that yields nothing.
func Empty[V any]() iter.Seq[V] {
	return func(yield func(V) bool) {}
}

// EmptyPair returns an iterator over pairs that yields nothing.
func EmptyPair[V1, V2 any]() iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {}
}
