package giopt

import (
	"iter"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/optional/optval"
)

// Find returns the first value among the values matching the predicate.
func Find[V any, P constraints.Predicate[V]](values iter.Seq[V], predicate P) optval.Value[V] {
	return optval.New(gi.Find(values, predicate))
}
