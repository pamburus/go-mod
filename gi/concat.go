package gi

import (
	"iter"
	"slices"
)

// Concat returns an iterator that concatenates the values of the given iterators.
func Concat[V any](values ...iter.Seq[V]) iter.Seq[V] {
	return Flatten(slices.Values(values))
}

// ConcatPairs returns an iterator that concatenates the values of the given pair iterators.
func ConcatPairs[V1, V2 any](values ...iter.Seq2[V1, V2]) iter.Seq2[V1, V2] {
	return FlattenPairs(slices.Values(values))
}

// ConcatSlices returns an iterator that over concatenation of the given slices.
func ConcatSlices[V any, VV ~[]V](values ...VV) iter.Seq[V] {
	return FlattenSlices(slices.Values(values))
}
