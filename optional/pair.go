package optional

import "iter"

// ---

// NewPair constructs a new optional Pair.
func NewPair[V1, V2 any](v1 V1, v2 V2, valid bool) Pair[V1, V2] {
	if valid {
		return SomePair(v1, v2)
	}

	return NonePair[V1, V2]()
}

// SomePair returns an optional Pair that has provided inner pair of values.
func SomePair[V1, V2 any](v1 V1, v2 V2) Pair[V1, V2] {
	return Pair[V1, V2]{
		v1, v2,
		true,
	}
}

// NonePair returns an empty optional Pair.
func NonePair[V1, V2 any]() Pair[V1, V2] {
	return Pair[V1, V2]{}
}

// PairByKey returns SomePair in case it is found in the provided map by the provided key.
// Otherwise, it returns NonePair.
func PairByKey[K comparable, V any, M ~map[K]V](key K, m M) Pair[K, V] {
	v, ok := m[key]

	return NewPair(key, v, ok)
}

// MapPair transforms optional Pair[V1, V2] to Pair[R1, R1] using the given function f.
// If the provided Pair is NonePair, it returns NonePair.
func MapPair[V1, V2, R1, R2 any, F ~func(V1, V2) (R1, R2)](pair Pair[V1, V2], f F) Pair[R1, R2] {
	if pair.IsSome() {
		return SomePair(f(pair.values()))
	}

	return NonePair[R1, R2]()
}

// FlatMapPair transforms optional Pair[V1, V2] to Pair[R1, R1] using the given function f.
// If the provided Pair is NonePair, it returns NonePair.
// If f returns NonePair, it returns NonePair.
func FlatMapPair[V1, V2, R1, R2 any, F ~func(V1, V2) Pair[R1, R2]](pair Pair[V1, V2], f F) Pair[R1, R2] {
	if pair.IsSome() {
		return f(pair.values())
	}

	return NonePair[R1, R2]()
}

// FilterPair calls the provided function f with the inner values of the provided Pair and returns the Pair if the function returns true.
// Otherwise, it returns NonePair.
func FilterPair[V1, V2 any, F ~func(V1, V2) bool](pair Pair[V1, V2], f F) Pair[V1, V2] {
	if pair.IsSome() && f(pair.values()) {
		return pair
	}

	return NonePair[V1, V2]()
}

// SomePairOnly returns a sequence of pairs that contain only Some values.
func SomePairOnly[V1, V2 any](pairs iter.Seq[Pair[V1, V2]]) iter.Seq2[V1, V2] {
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

// Both returns a SomePair containing the inner values of value1 and value2 if both are Some.
// Otherwise, it returns NonePair.
func Both[V1, V2 any](value1 Value[V1], value2 Value[V2]) Pair[V1, V2] {
	if v1, ok := value1.Unwrap(); ok {
		if v2, ok := value2.Unwrap(); ok {
			return SomePair(v1, v2)
		}
	}

	return NonePair[V1, V2]()
}

// Left returns the left value of the provided pair if it is Some.
// Otherwise, it returns None.
func Left[V1, V2 any](pair Pair[V1, V2]) Value[V1] {
	if pair.IsSome() {
		return Some(pair.v1)
	}

	return None[V1]()
}

// Right returns the right value of the provided pair if it is Some.
// Otherwise, it returns None.
func Right[V1, V2 any](pair Pair[V1, V2]) Value[V2] {
	if pair.IsSome() {
		return Some(pair.v2)
	}

	return None[V2]()
}

// ---

// Pair represents an optional pair of values of types V1 and V2.
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

// Or returns p if it is Some or other pair.
func (p Pair[V1, V2]) Or(other Pair[V1, V2]) Pair[V1, V2] {
	if p.IsSome() {
		return p
	}

	return other
}

// OrSome returns the inner values if present, otherwise it returns provided values.
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

// OrElse returns p if the inner values are present, otherwise it calls provided function and returns its result.
func (p Pair[V1, V2]) OrElse(value func() Pair[V1, V2]) Pair[V1, V2] {
	if p.IsSome() {
		return p
	}

	return value()
}

// OrElseSome returns the inner values if present, otherwise it calls provided function and returns its result.
func (p Pair[V1, V2]) OrElseSome(f func() (V1, V2)) (V1, V2) {
	if p.IsSome() {
		return p.v1, p.v2
	}

	return f()
}

// Reset resets p to NonePair.
func (p *Pair[V1, V2]) Reset() {
	*p = NonePair[V1, V2]()
}

// Take returns a copy of p and resets it to NonePair.
func (p *Pair[V1, V2]) Take() Pair[V1, V2] {
	result := *p
	p.Reset()

	return result
}

// Replace returns a copy of p and resets it to SomePair.
func (p *Pair[V1, V2]) Replace(v1 V1, v2 V2) Pair[V1, V2] {
	result := *p
	*p = SomePair(v1, v2)

	return result
}

func (p Pair[V1, V2]) values() (V1, V2) {
	return p.v1, p.v2
}
