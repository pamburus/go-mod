package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/gi/gi2op"
)

// Sum returns the sum of the given values.
func Sum[V1, V2 constraints.Number](pairs iter.Seq2[V1, V2]) (V1, V2) {
	return Fold(pairs, 0, 0, gi2op.Add)
}

// Product returns the product of the given values.
func Product[V1, V2 constraints.Number](values iter.Seq2[V1, V2]) (V1, V2) {
	return Fold(values, 1, 1, gi2op.Multiply)
}
