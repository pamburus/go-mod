package optpair

import (
	"cmp"
	"iter"

	"github.com/pamburus/go-mod/optional/internal/cmpx"
	"github.com/pamburus/go-mod/optional/optval"
)

// ---

// New constructs a new optional [Pair].
func New[V1, V2 any](v1 V1, v2 V2, valid bool) Pair[V1, V2] {
	if valid {
		return Some(v1, v2)
	}

	return None[V1, V2]()
}

// Some returns an optional [Pair] that has the given inner pair of values.
func Some[V1, V2 any](v1 V1, v2 V2) Pair[V1, V2] {
	return Pair[V1, V2]{
		v1, v2,
		true,
	}
}

// None returns an empty optional [Pair].
func None[V1, V2 any]() Pair[V1, V2] {
	return Pair[V1, V2]{}
}

// ByKey returns [Some] in case it is found in the provided map by the given key.
// Otherwise, it returns [None].
func ByKey[K comparable, V any, M ~map[K]V](key K, m M) Pair[K, V] {
	v, ok := m[key]

	return New(key, v, ok)
}

// Map transforms optional [Pair][V1, V2] to [Pair][R1, R1] using the given function f.
// If the provided [Pair] is [None], it returns [None].
func Map[V1, V2, R1, R2 any, F ~func(V1, V2) (R1, R2)](pair Pair[V1, V2], f F) Pair[R1, R2] {
	if pair.IsSome() {
		return Some(f(pair.values()))
	}

	return None[R1, R2]()
}

// FlatMap transforms optional [Pair][V1, V2] to [Pair][R1, R1] using the given function f.
// If the provided [Pair] is [None], it returns [None].
// If f returns [None], it returns [None].
func FlatMap[V1, V2, R1, R2 any, F ~func(V1, V2) Pair[R1, R2]](pair Pair[V1, V2], f F) Pair[R1, R2] {
	if pair.IsSome() {
		return f(pair.values())
	}

	return None[R1, R2]()
}

// Filter calls the given function f with the inner values of the given pair
// and returns the same pair if the function returns true.
// Otherwise, it returns [None].
func Filter[V1, V2 any, F ~func(V1, V2) bool](pair Pair[V1, V2], f F) Pair[V1, V2] {
	if pair.IsSome() && f(pair.values()) {
		return pair
	}

	return None[V1, V2]()
}

// SomeOnly returns a sequence of pairs that contain only [Some] values from the given pairs.
func SomeOnly[V1, V2 any](pairs iter.Seq[Pair[V1, V2]]) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		for pair := range pairs {
			if v1, v2, ok := pair.Unwrap(); ok {
				if !yield(v1, v2) {
					break
				}
			}
		}
	}
}

// Compare compares two optional pairs.
// [Some] is considered less than [None].
func Compare[V1, V2 cmp.Ordered](a, b Pair[V1, V2]) int {
	switch {
	case a.valid != b.valid:
		return cmpx.IfElse(a.valid, -1, 1)
	case !a.valid:
		return 0
	case a.v1 != b.v1:
		return cmp.Compare(a.v1, b.v1)
	default:
		return cmp.Compare(a.v2, b.v2)
	}
}

// Less returns true if a is less b.
// [Some] is considered less than [None].
func Less[V1, V2 cmp.Ordered](a, b Pair[V1, V2]) bool {
	return Compare(a, b) < 0
}

// IfBoth returns [Some] pair containing the inner values of value1 and value2 if both are [optval.Some].
// Otherwise, it returns [None].
func IfBoth[V1, V2 any](v1 optval.Value[V1], v2 optval.Value[V2]) Pair[V1, V2] {
	if v1, ok := v1.Unwrap(); ok {
		if v2, ok := v2.Unwrap(); ok {
			return Some(v1, v2)
		}
	}

	return None[V1, V2]()
}

// Left returns the left value of the given pair if it is [Some].
// Otherwise, it returns [None].
func Left[V1, V2 any](pair Pair[V1, V2]) optval.Value[V1] {
	if pair.IsSome() {
		return optval.Some(pair.v1)
	}

	return optval.None[V1]()
}

// Right returns the right value of the given pair if it is [Some].
// Otherwise, it returns [None].
func Right[V1, V2 any](pair Pair[V1, V2]) optval.Value[V2] {
	if pair.IsSome() {
		return optval.Some(pair.v2)
	}

	return optval.None[V2]()
}

// ---

// Pair represents an optional pair of values of types V1 and V2.
// It can be either [Some] or [None].
// If it is [Some], it contains the inner values of types V1 and V2.
type Pair[V1, V2 any] struct {
	v1    V1
	v2    V2
	valid bool
}

// Unwrap returns the inner values of type V1 and V2 and true if present.
// Otherwise, it returns zero values and false.
func (p Pair[V1, V2]) Unwrap() (V1, V2, bool) {
	return p.v1, p.v2, p.valid
}

// IsSome returns true if the optional pair p has inner values.
func (p Pair[V1, V2]) IsSome() bool {
	return p.valid
}

// IsNone returns true if the optional pair p has no inner values.
func (p Pair[V1, V2]) IsNone() bool {
	return !p.IsSome()
}

// Or returns p if it is [Some] or other given pair otherwise.
func (p Pair[V1, V2]) Or(other Pair[V1, V2]) Pair[V1, V2] {
	if p.IsSome() {
		return p
	}

	return other
}

// OrSome returns the inner values if present, otherwise it returns the given values.
func (p Pair[V1, V2]) OrSome(v1 V1, v2 V2) (V1, V2) {
	if p.IsSome() {
		return p.v1, p.v2
	}

	return v1, v2
}

// OrZero returns the inner values if present, otherwise it returns zero initialized values.
func (p Pair[V1, V2]) OrZero() (V1, V2) {
	return p.v1, p.v2
}

// OrElse returns p if the inner values are present, otherwise it calls f and returns its result.
func (p Pair[V1, V2]) OrElse(f func() Pair[V1, V2]) Pair[V1, V2] {
	if p.IsSome() {
		return p
	}

	return f()
}

// OrElseSome returns the inner values if present, otherwise it calls f and returns its result.
func (p Pair[V1, V2]) OrElseSome(f func() (V1, V2)) (V1, V2) {
	if p.IsSome() {
		return p.v1, p.v2
	}

	return f()
}

// Reset resets p to [None].
func (p *Pair[V1, V2]) Reset() {
	*p = None[V1, V2]()
}

// Take returns a copy of p and resets it to [None].
func (p *Pair[V1, V2]) Take() Pair[V1, V2] {
	result := *p
	p.Reset()

	return result
}

// Replace returns a copy of p and resets it to [Some] with v1 and v2.
func (p *Pair[V1, V2]) Replace(v1 V1, v2 V2) Pair[V1, V2] {
	result := *p
	*p = Some(v1, v2)

	return result
}

func (p Pair[V1, V2]) values() (V1, V2) {
	return p.v1, p.v2
}
