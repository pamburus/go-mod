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

// FlattenSlices returns an iterator that flattens the slices of the given iterator.
func FlattenSlices[V any, VV ~[]V](values iter.Seq[VV]) iter.Seq[V] {
	return FlattenBy(values, slices.Values)
}
