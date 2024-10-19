package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Contains returns true if the given values contain a value matching the predicate.
func Contains[V any, P constraints.Predicate[V]](values iter.Seq[V], predicate P) bool {
	return IsSome(Find(values, predicate))
}
