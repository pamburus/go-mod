package gi2

import (
	"cmp"
	"iter"

	"github.com/pamburus/go-mod/gi/giop2"
	"github.com/pamburus/go-mod/optional/optpair"
)

// Max returns maximum pair of the given pair or [optpair.None] if pairs is empty.
func Max[V1, V2 cmp.Ordered](values iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
	return Reduce(values, giop2.Max)
}

// MaxBy returns maximum pair of the given pair or [optpair.None] if pairs is empty.
// It uses provided key function for comparison.
func MaxBy[V1, V2 any, K cmp.Ordered, F ~func(V1, V2) K](pairs iter.Seq2[V1, V2], key F) optpair.Pair[V1, V2] {
	return Reduce(pairs, giop2.MaxBy(key))
}
