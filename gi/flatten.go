package gi

import (
	"iter"
	"slices"
)

// Flatten returns an iterator that flattens the values of the given iterator.
func Flatten[V any](values iter.Seq[iter.Seq[V]]) iter.Seq[V] {
	return FlattenBy(values, func(vs iter.Seq[V]) iter.Seq[V] {
		return vs
	})
}

// FlattenPairs returns an iterator that flattens the values of the given pair iterator.
func FlattenPairs[V1, V2 any](values iter.Seq[iter.Seq2[V1, V2]]) iter.Seq2[V1, V2] {
	return FlattenPairsBy(values, func(vs iter.Seq2[V1, V2]) iter.Seq2[V1, V2] {
		return vs
	})
}

// FlattenBy returns an iterator that flattens the values of the given iterator using the provided function.
func FlattenBy[VV any, V any, F ~func(VV) iter.Seq[V]](values iter.Seq[VV], f F) iter.Seq[V] {
	return func(yield func(V) bool) {
		for vs := range values {
			for v := range f(vs) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// FlattenPairsBy returns an iterator that flattens the values of the given iterator using the provided function.
func FlattenPairsBy[VV any, V1, V2 any, F ~func(VV) iter.Seq2[V1, V2]](values iter.Seq[VV], f F) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		for vs := range values {
			for v1, v2 := range f(vs) {
				if !yield(v1, v2) {
					return
				}
			}
		}
	}
}

// FlattenSlices returns an iterator that flattens the slices of the given iterator.
func FlattenSlices[V any, VV ~[]V](values iter.Seq[VV]) iter.Seq[V] {
	return FlattenBy(values, slices.Values)
}
