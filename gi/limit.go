package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Limit returns an iterator sequence that yields at most n values.
func Limit[V any, I constraints.Integer](n I, values iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		i := I(0)
		for v := range values {
			if i >= n {
				return
			}

			if !yield(v) {
				return
			}

			i++
		}
	}
}
