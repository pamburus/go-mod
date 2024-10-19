package giopt

import (
	"iter"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/optional/optval"
)

// Find returns the first value among the values matching the predicate.
func Find[V any, P constraints.Predicate[V]](values iter.Seq[V], predicate P) optval.Value[V] {
	return FindWith(predicate)(values)
}

// FindWith returns a function that finds the first value among the values matching the predicate.
func FindWith[V any, P constraints.Predicate[V]](predicate P) func(iter.Seq[V]) optval.Value[V] {
	return func(values iter.Seq[V]) optval.Value[V] {
		return optval.New(gi.Find(values, predicate))
	}
}
