package gi2opt

import (
	"cmp"
	"iter"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/gi2op"
	"github.com/pamburus/go-mod/optional/optpair"
)

// Min returns minimum pair of the given pair or [optpair.None] if pairs is empty.
func Min[V1, V2 cmp.Ordered](pairs iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
	return optpair.New(gi2.Min(pairs))
}

// MinBy returns minimum pair of the given pair or [optpair.None] if pairs is empty.
// It uses provided key function for comparison.
func MinBy[V1, V2 any, K cmp.Ordered, F ~func(V1, V2) K](pairs iter.Seq2[V1, V2], key F) optpair.Pair[V1, V2] {
	return Reduce(pairs, gi2op.MinBy(key))
}

// MinByLeft returns the pair with the minimum left value or [optpair.None] if pairs is empty.
func MinByLeft[V1 cmp.Ordered, V2 any](pairs iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
	return optpair.New(gi2.MinByLeft(pairs))
}

// MinByRight returns the pair with the minimum right value or [optpair.None] if pairs is empty.
func MinByRight[V1 any, V2 cmp.Ordered](pairs iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
	return optpair.New(gi2.MinByRight(pairs))
}
