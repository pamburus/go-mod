package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Cloned returns a new iterator that consumes original iterator and clones each value.
func Cloned[V constraints.Cloneable[V]](values iter.Seq[V]) iter.Seq[V] {
	return Map(values, V.Clone)
}
