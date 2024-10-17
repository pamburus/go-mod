package gi2

import "iter"

// Map returns an iterator adapter over pairs that uses a transform function to transform pairs.
func Map[V1, V2, R1, R2 any, F ~func(V1, V2) (R1, R2)](pairs iter.Seq2[V1, V2], transform F) iter.Seq2[R1, R2] {
	return func(yield func(R1, R2) bool) {
		for v1, v2 := range pairs {
			if !yield(transform(v1, v2)) {
				return
			}
		}
	}
}

// MapLeft returns an iterator adapter over pairs that uses a transform function to transform first values of pairs.
func MapLeft[V1, V2, R any, F ~func(V1) R](pairs iter.Seq2[V1, V2], transform F) iter.Seq2[R, V2] {
	return func(yield func(R, V2) bool) {
		for v1, v2 := range pairs {
			if !yield(transform(v1), v2) {
				return
			}
		}
	}
}

// MapRight returns an iterator adapter over pairs that uses a transform function to transform second values of pairs.
func MapRight[V1, V2, R any, F ~func(V2) R](pairs iter.Seq2[V1, V2], transform F) iter.Seq2[V1, R] {
	return func(yield func(V1, R) bool) {
		for v1, v2 := range pairs {
			if !yield(v1, transform(v2)) {
				return
			}
		}
	}
}

// MapKeys is an alias for [MapLeft] that assumes the given iterator is an iterator over key/value pairs.
func MapKeys[K, V, R any, F ~func(K) R](pairs iter.Seq2[K, V], transform F) iter.Seq2[R, V] {
	return MapLeft(pairs, transform)
}

// MapValues is an alias for [MapRight] that assumes the given iterator is an iterator over key/value pairs.
func MapValues[K, V, R any, F ~func(V) R](pairs iter.Seq2[K, V], transform F) iter.Seq2[K, R] {
	return MapRight(pairs, transform)
}
