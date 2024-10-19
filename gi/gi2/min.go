package gi2

import (
	"cmp"
	"iter"

	"github.com/pamburus/go-mod/gi/gi2op"
)

// Min returns the minimum pair of the given pair sequence and true if it is not empty.
// Otherwise, it returns zero values and false.
func Min[V1, V2 cmp.Ordered](pairs iter.Seq2[V1, V2]) (V1, V2, bool) {
	return Reduce(pairs, gi2op.Min)
}

// MinBy returns the minimum pair of the given pair sequence and true if it is not empty.
// Otherwise, it returns zero values and false.
// It uses the provided key function for comparison.
func MinBy[V1, V2 any, K cmp.Ordered, F ~func(V1, V2) K](pairs iter.Seq2[V1, V2], key F) (V1, V2, bool) {
	return Reduce(pairs, gi2op.MinBy(key))
}

// MinByLeft returns the pair with the minimum left value and true if it is not empty.
// Otherwise, it returns zero values and false.
func MinByLeft[V1 cmp.Ordered, V2 any](pairs iter.Seq2[V1, V2]) (V1, V2, bool) {
	return Reduce(pairs, gi2op.MinByLeft)
}

// MinByRight returns the pair with the minimum right value and true if it is not empty.
// Otherwise, it returns zero values and false.
func MinByRight[V1 any, V2 cmp.Ordered](pairs iter.Seq2[V1, V2]) (V1, V2, bool) {
	return Reduce(pairs, gi2op.MinByRight)
}
