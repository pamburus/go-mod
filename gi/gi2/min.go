package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/gi/giop2"
	"github.com/pamburus/go-mod/optional/optpair"
)

// Min returns minimum pair of the given pair or [optpair.None] if pairs is empty.
func Min[V1, V2 constraints.Ordered](pairs iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
	return Reduce(pairs, giop2.Min[V1, V2])
}

// MinBy returns minimum pair of the given pair or [optpair.None] if pairs is empty.
// It uses provided key function for comparison.
func MinBy[V1, V2 any, K constraints.Ordered, F ~func(V1, V2) K](pairs iter.Seq2[V1, V2], key F) optpair.Pair[V1, V2] {
	return Reduce(pairs, giop2.MinBy(key))
}
