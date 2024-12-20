package optpair

import (
	"cmp"
	"iter"
	"slices"

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

// FromValue converts [optval.Value] to [Pair] using the given function f.
func FromValue[V, R1, R2 any, F ~func(V) (R1, R2)](value optval.Value[V], f F) Pair[R1, R2] {
	if v, ok := value.Unwrap(); ok {
		return Some(f(v))
	}

	return None[R1, R2]()
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

// UnwrapFilter returns a sequence of pairs that contain only [Some] values from the given pairs.
func UnwrapFilter[V1, V2 any](pairs iter.Seq[Pair[V1, V2]]) iter.Seq2[V1, V2] {
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

// JoinAnd returns a [Some] pair containing the inner values of v1 and v2 if they are both [optval.Some].
// Otherwise, it returns [None].
func JoinAnd[V1, V2 any](v1 optval.Value[V1], v2 optval.Value[V2]) Pair[V1, V2] {
	if v1, ok := v1.Unwrap(); ok {
		if v2, ok := v2.Unwrap(); ok {
			return Some(v1, v2)
		}
	}

	return None[V1, V2]()
}

// JoinOr returns a [Some] pair containing the inner values of v1 and v2 if at least one of them is [optval.Some].
// Otherwise, it returns [None].
// If some of the values are [None], they are replaced with zero values.
func JoinOr[V1, V2 any](v1 optval.Value[V1], v2 optval.Value[V2]) Pair[V1, V2] {
	if v1, ok := v1.Unwrap(); ok {
		return Some(v1, v2.OrZero())
	}

	if v2, ok := v2.Unwrap(); ok {
		return Some(v1.OrZero(), v2)
	}

	return None[V1, V2]()
}

// Split returns the inner values of the given pair as [optval.Value]s.
func Split[V1, V2 any](pair Pair[V1, V2]) (optval.Value[V1], optval.Value[V2]) {
	return pair.Split()
}

// Left returns the left value of the given pair if it is [Some].
// Otherwise, it returns [None].
func Left[V1, V2 any](pair Pair[V1, V2]) optval.Value[V1] {
	return pair.Left()
}

// Right returns the right value of the given pair if it is [Some].
// Otherwise, it returns [None].
func Right[V1, V2 any](pair Pair[V1, V2]) optval.Value[V2] {
	return pair.Right()
}

// Swap returns a new pair with swapped inner values.
func Swap[V1, V2 any](pair Pair[V1, V2]) Pair[V2, V1] {
	return pair.Swap()
}

// IsSome returns true if the given pair is [Some].
func IsSome[V1, V2 any](pair Pair[V1, V2]) bool {
	return pair.IsSome()
}

// IsNone returns true if the given pair is [None].
func IsNone[V1, V2 any](pair Pair[V1, V2]) bool {
	return pair.IsNone()
}

// Or returns the first [Some] pair from the given pairs.
// If all pairs are [None], it returns [None].
func Or[V1, V2 any](pairs ...Pair[V1, V2]) Pair[V1, V2] {
	return FindSome(slices.Values(pairs))
}

// OrZero returns the inner values of the given pair if it is [Some].
// Otherwise, it returns zero initialized values.
func OrZero[V1, V2 any](pair Pair[V1, V2]) (V1, V2) {
	return pair.OrZero()
}

// Unwrap returns the inner values of the given pair if it is [Some].
// Otherwise, it returns zero values and false.
func Unwrap[V1, V2 any](pair Pair[V1, V2]) (V1, V2, bool) {
	return pair.Unwrap()
}

// FindSome returns the first optional pair that is [Some] or [None] if all pairs are [None].
// It returns [None] if the sequence is empty.
func FindSome[V1, V2 any](pairs iter.Seq[Pair[V1, V2]]) Pair[V1, V2] {
	for pair := range pairs {
		if pair.IsSome() {
			return pair
		}
	}

	return None[V1, V2]()
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

// Left returns the left value of the pair if it is [Some].
func (p Pair[V1, V2]) Left() optval.Value[V1] {
	if p.IsSome() {
		return optval.Some(p.v1)
	}

	return optval.None[V1]()
}

// Right returns the right value of the pair if it is [Some].
func (p Pair[V1, V2]) Right() optval.Value[V2] {
	if p.IsSome() {
		return optval.Some(p.v2)
	}

	return optval.None[V2]()
}

// Swap returns a new pair with swapped inner values.
func (p Pair[V1, V2]) Swap() Pair[V2, V1] {
	if p.IsSome() {
		return Some(p.v2, p.v1)
	}

	return None[V2, V1]()
}

// Split returns the inner values of the pair as [optval.Value]s.
func (p Pair[V1, V2]) Split() (optval.Value[V1], optval.Value[V2]) {
	return p.Left(), p.Right()
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
