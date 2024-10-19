package gi2

import (
	"cmp"
	"iter"

	"github.com/pamburus/go-mod/gi/gi2op"
)

// Max returns the maximum pair of the given pair sequence and true if it is not empty.
// Otherwise, it returns zero values and false.
func Max[V1, V2 cmp.Ordered](values iter.Seq2[V1, V2]) (V1, V2, bool) {
	return Reduce(values, gi2op.Max)
}

// MaxBy returns the maximum pair of the given pair sequence and true if it is not empty.
// Otherwise, it returns zero values and false.
// It uses the provided key function for comparison.
func MaxBy[V1, V2 any, K cmp.Ordered, F ~func(V1, V2) K](pairs iter.Seq2[V1, V2], key F) (V1, V2, bool) {
	return Reduce(pairs, gi2op.MaxBy(key))
}

// MaxByLeft returns the pair with the maximum left value and true if it is not empty.
// Otherwise, it returns zero values and false.
func MaxByLeft[V1 cmp.Ordered, V2 any](values iter.Seq2[V1, V2]) (V1, V2, bool) {
	return Reduce(values, gi2op.MaxByLeft)
}

// MaxByRight returns the pair with the maximum right value and true if it is not empty.
// Otherwise, it returns zero values and false.
func MaxByRight[V1 any, V2 cmp.Ordered](values iter.Seq2[V1, V2]) (V1, V2, bool) {
	return Reduce(values, gi2op.MaxByRight)
}
