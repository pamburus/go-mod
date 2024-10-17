package gi

import (
	"iter"
)

// Empty returns an iterator that yields nothing.
func Empty[V any]() iter.Seq[V] {
	return func(func(V) bool) {}
}
