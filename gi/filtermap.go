package gi

import (
	"iter"

	"github.com/pamburus/go-mod/optional/optval"
)

// FilterMap returns an iterator that yields values by applying the provided filter-map function to values.
// The filter-map function returns an optional value.
// If the optional value is empty, the value is skipped.
func FilterMap[V, R any, F func(V) optval.Value[R]](values iter.Seq[V], f F) iter.Seq[R] {
	return FilterMapWith(f)(values)
}

// FilterMapWith returns a function that can be used to filter-map sequences of values.
// It uses the provided filter-map function to transform and filter values.
// If the optional value is empty, the value is skipped.
func FilterMapWith[V, R any, F func(V) optval.Value[R]](f F) func(values iter.Seq[V]) iter.Seq[R] {
	return func(values iter.Seq[V]) iter.Seq[R] {
		return func(yield func(R) bool) {
			for v := range values {
				if r, ok := f(v).Unwrap(); ok {
					if !yield(r) {
						return
					}
				}
			}
		}
	}
}
