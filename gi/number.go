package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/gic"
	"github.com/pamburus/go-mod/gi/giop"
)

// Sum returns the sum of the given values.
func Sum[V gic.Number](values iter.Seq[V]) V {
	return Fold(values, 0, giop.Add[V])
}

// Product returns the product of the given values.
func Product[V gic.Number](values iter.Seq[V]) V {
	return Fold(values, 1, giop.Multiply[V])
}
