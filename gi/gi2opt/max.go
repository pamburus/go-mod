package gi2opt

import (
	"cmp"
	"iter"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/optional/optpair"
)

// Max returns the maximum pair of the given pair sequence or [optpair.None] if it is empty.
func Max[V1, V2 cmp.Ordered](pairs iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
	return optpair.New(gi2.Max(pairs))
}

// MaxBy returns the maximum pair of the given pair sequence or [optpair.None] if it is empty.
// It uses the provided key function for comparison.
func MaxBy[V1, V2 any, K cmp.Ordered, F ~func(V1, V2) K](pairs iter.Seq2[V1, V2], key F) optpair.Pair[V1, V2] {
	return optpair.New(gi2.MaxBy(pairs, key))
}

// MaxByLeft returns the pair with the maximum left value or [optpair.None] if it is empty.
func MaxByLeft[V1 cmp.Ordered, V2 any](pairs iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
	return optpair.New(gi2.MaxByLeft(pairs))
}

// MaxByRight returns the pair with the maximum right value or [optpair.None] if it is empty.
func MaxByRight[V1 any, V2 cmp.Ordered](pairs iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
	return optpair.New(gi2.MaxByRight(pairs))
}
