package gi

import (
	"iter"
)

// Single returns an iterator over a single value.
func Single[V any](value V) iter.Seq[V] {
	return func(yield func(V) bool) {
		yield(value)
	}
}

// SinglePair returns an iterator over a single pair of values.
func SinglePair[V1, V2 any](v1 V1, v2 V2) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		yield(v1, v2)
	}
}
