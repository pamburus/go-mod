package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/gi/giop"
)

// Sum returns the sum of the given values.
func Sum[V constraints.Number](values iter.Seq[V]) V {
	return Fold(values, 0, giop.Add)
}

// Product returns the product of the given values.
func Product[V constraints.Number](values iter.Seq[V]) V {
	return Fold(values, 1, giop.Multiply)
}
