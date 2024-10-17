package gi2

import (
	"iter"
)

// Flatten returns an iterator that flattens the values of the given pair iterator.
func Flatten[V1, V2 any](values iter.Seq[iter.Seq2[V1, V2]]) iter.Seq2[V1, V2] {
	return FlattenBy(values, func(vs iter.Seq2[V1, V2]) iter.Seq2[V1, V2] {
		return vs
	})
}

// FlattenBy returns an iterator that flattens the values of the given iterator using the provided function.
func FlattenBy[VV any, V1, V2 any, F ~func(VV) iter.Seq2[V1, V2]](values iter.Seq[VV], f F) iter.Seq2[V1, V2] {
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
