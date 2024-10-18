package gi

import (
	"iter"
	"slices"
)

// Concat returns an iterator that concatenates the values of the given iterators.
func Concat[V any](values ...iter.Seq[V]) iter.Seq[V] {
	return Flatten(slices.Values(values))
}

// ConcatSlices returns an iterator that over concatenation of the given slices.
func ConcatSlices[V any, VV ~[]V](values ...VV) iter.Seq[V] {
	return FlattenSlices(slices.Values(values))
}
