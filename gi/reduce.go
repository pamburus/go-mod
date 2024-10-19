package gi

import (
	"iter"
)

// Reduce reduces the values to a value which is the accumulated result of running accumulate function
// on each element where each successive invocation is supplied the return value of the previous one.
//
// It returns zero value and false if the sequence is empty.
func Reduce[V any, F ~func(V, V) V](values iter.Seq[V], accumulate F) (V, bool) {
	return ReduceWith(accumulate)(values)
}

// ReduceWith returns a function that can be used to reduce sequences of values in a single value.
func ReduceWith[V any, F ~func(V, V) V](accumulate F) func(iter.Seq[V]) (V, bool) {
	return func(values iter.Seq[V]) (V, bool) {
		type opt struct {
			value V
			valid bool
		}

		result := Fold(values, opt{}, func(result opt, value V) opt {
			if result.valid {
				return opt{
					value: accumulate(result.value, value),
					valid: true,
				}
			}

			return opt{value, true}
		})

		return result.value, result.valid
	}
}
