package gi2

import (
	"iter"
	"slices"
)

// Concat returns an iterator that concatenates the values of the given pair iterators.
func Concat[V1, V2 any](values ...iter.Seq2[V1, V2]) iter.Seq2[V1, V2] {
	return Flatten(slices.Values(values))
}
