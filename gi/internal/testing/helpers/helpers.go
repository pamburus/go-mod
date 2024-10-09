// Package helpers provides helper functions for testing.
package helpers

import "iter"

// ToPairs returns an iterator that yields pairs by applying the provided function to values.
func ToPairs[V, R1, R2 any, F ~func(V) (R1, R2)](values iter.Seq[V], f F) iter.Seq2[R1, R2] {
	return func(yield func(R1, R2) bool) {
		for v := range values {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// FlattenPairs returns an iterator that yields values by flattening pairs.
func FlattenPairs[V any](pairs iter.Seq2[V, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v1, v2 := range pairs {
			if !yield(v1) {
				return
			}
			if !yield(v2) {
				return
			}
		}
	}
}

// Swap returns an iterator that yields pairs with swapped values.
func Swap[V1, V2 any](pairs iter.Seq2[V1, V2]) iter.Seq2[V2, V1] {
	return func(yield func(V2, V1) bool) {
		for v1, v2 := range pairs {
			if !yield(v2, v1) {
				return
			}
		}
	}
}

// Limit returns an iterator sequence that yields at most n values.
func Limit[V any](values iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		i := n
		for v := range values {
			if i == 0 {
				return
			}
			i--
			if !yield(v) {
				return
			}
		}
	}
}

// LimitPairs returns an iterator sequence that yields at most n pairs.
func LimitPairs[V1, V2 any](values iter.Seq2[V1, V2], n int) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		i := n
		for v1, v2 := range values {
			if i == 0 {
				return
			}
			i--
			if !yield(v1, v2) {
				return
			}
		}
	}
}
