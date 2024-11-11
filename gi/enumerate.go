package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Enumerate returns an iterator over the given values, yielding pairs of the index and the value.
func Enumerate[V any](values iter.Seq[V]) iter.Seq2[int, V] {
	return EnumerateFrom(0, values)
}

// EnumerateFrom returns an iterator over the given values, yielding pairs of the index and the value.
// The index starts from the given begin value.
func EnumerateFrom[V any, I constraints.Integer](begin I, values iter.Seq[V]) iter.Seq2[I, V] {
	return func(yield func(I, V) bool) {
		i := begin
		for value := range values {
			if !yield(i, value) {
				return
			}

			i++
		}
	}
}
