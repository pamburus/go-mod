package gi2

import "iter"

// Pack returns an iterator that yields the results of applying the provided pack function to pairs.
func Pack[V1, V2, R any, F ~func(V1, V2) R](pairs iter.Seq2[V1, V2], wrap F) iter.Seq[R] {
	return PackWith(wrap)(pairs)
}

// PackWith returns a function that transforms an iterator sequence over pairs to an iterator sequence over values.
// It uses the provided pack function to transform pairs to values.
func PackWith[V1, V2, R any, F ~func(V1, V2) R](pack F) func(iter.Seq2[V1, V2]) iter.Seq[R] {
	return func(pairs iter.Seq2[V1, V2]) iter.Seq[R] {
		return func(yield func(R) bool) {
			for v1, v2 := range pairs {
				if !yield(pack(v1, v2)) {
					return
				}
			}
		}
	}
}

// Unpack returns an iterator that yields the results of applying the provided unpack function to values.
func Unpack[V, R1, R2 any, F ~func(V) (R1, R2)](values iter.Seq[V], unpack F) iter.Seq2[R1, R2] {
	return UnpackWith(unpack)(values)
}

// UnpackWith returns a function that transforms an iterator sequence over values to an iterator sequence over pairs.
// It uses the provided unpack function to transform values to pairs.
func UnpackWith[V, R1, R2 any, F ~func(V) (R1, R2)](unpack F) func(iter.Seq[V]) iter.Seq2[R1, R2] {
	return func(values iter.Seq[V]) iter.Seq2[R1, R2] {
		return func(yield func(R1, R2) bool) {
			for v := range values {
				if !yield(unpack(v)) {
					return
				}
			}
		}
	}
}
