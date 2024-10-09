package optval

import (
	"cmp"
	"iter"

	"github.com/pamburus/go-mod/optional/internal/cmpx"
)

// ---

// New constructs a new optional [Value].
func New[T any](value T, valid bool) Value[T] {
	if valid {
		return Some(value)
	}

	return None[T]()
}

// Some returns an optional [Value] that has provided inner value.
func Some[T any](value T) Value[T] {
	return Value[T]{
		value,
		true,
	}
}

// None returns an optional [Value] that has no inner value.
func None[T any]() Value[T] {
	return Value[T]{}
}

// ByKey returns [Some] value in case it is found in the provided map by the provided key.
func ByKey[K comparable, V any, M ~map[K]V](key K, m M) Value[V] {
	v, ok := m[key]

	return New(v, ok)
}

// Key returns [Some] key in case it is found in the provided map.
func Key[K comparable, V any, M ~map[K]V](key K, m M) Value[K] {
	_, ok := m[key]

	return New(key, ok)
}

// FromPtr returns [Some] with a copy of the given value if the provided pointer is not nil.
// Otherwise, it returns [None].
func FromPtr[T any](value *T) Value[T] {
	if value != nil {
		return Some(*value)
	}

	return None[T]()
}

// Map transforms optional Value[T] to optional Value[U] using the given function.
func Map[T, U any, F ~func(T) U](v Value[T], f F) Value[U] {
	if v.IsSome() {
		return Some(f(v.inner))
	}

	return None[U]()
}

// MapFromPtr transforms v to optional [Value] using the given function.
// If v is nil, it returns [None].
// Otherwise, it returns [Some] with the result of the function call.
func MapFromPtr[T, U any, F ~func(T) U](v *T, f F) Value[U] {
	return Map(FromPtr(v), f)
}

// FlatMap transforms optional [Value][T] to optional [Value][U] using the given function.
func FlatMap[T, U any, F ~func(T) Value[U]](v Value[T], f F) Value[U] {
	if v.IsSome() {
		return f(v.inner)
	}

	return None[U]()
}

// Flatten flattens optional [Value][[Value][T]] to [Value][T].
func Flatten[T any](v Value[Value[T]]) Value[T] {
	if v.IsSome() {
		return v.inner
	}

	return None[T]()
}

// Filter filters optional [Value] by the given predicate.
// If the value is [Some] and the predicate returns true, it returns the same value.
// Otherwise, it returns [None].
func Filter[T any, F ~func(T) bool](v Value[T], f F) Value[T] {
	if v.IsSome() && f(v.inner) {
		return v
	}

	return None[T]()
}

// Collect collects all values from the given optional values and returns them as a slice.
func Collect[T any](values ...Value[T]) []T {
	result := make([]T, 0, len(values))

	for _, value := range values {
		if v, ok := value.Unwrap(); ok {
			result = append(result, v)
		}
	}

	return result
}

// UnwrapFilter returns a sequence of values that contain inner values from [Some] values.
func UnwrapFilter[T any](values iter.Seq[Value[T]]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range values {
			if v, ok := value.Unwrap(); ok {
				if !yield(v) {
					break
				}
			}
		}
	}
}

// Compare compares two optional values.
// [Some] is considered less than [None].
func Compare[T cmp.Ordered](a, b Value[T]) int {
	switch {
	case a.valid != b.valid:
		return cmpx.IfElse(a.valid, -1, 1)
	case !a.valid:
		return 0
	default:
		return cmp.Compare(a.inner, b.inner)
	}
}

// Less returns true if the first value is less than the second one.
// [Some] is considered less than [None].
func Less[T cmp.Ordered](a, b Value[T]) bool {
	return Compare(a, b) < 0
}

// ---

// Value represents an optional value of type T.
// It can be either [Some] with inner value or [None].
type Value[T any] struct {
	inner T
	valid bool
}

// Unwrap returns the inner value of type T and true if it is [Some].
// Otherwise it returns zero value and false.
func (v Value[T]) Unwrap() (T, bool) {
	return v.inner, v.valid
}

// IsSome returns true if the optional value v has inner value.
func (v Value[T]) IsSome() bool {
	return v.valid
}

// IsNone returns true if the optional value v has no inner value.
func (v Value[T]) IsNone() bool {
	return !v.IsSome()
}

// Or returns the optional value v if it is [Some], otherwise it returns provided value.
func (v Value[T]) Or(other Value[T]) Value[T] {
	if v.IsSome() {
		return v
	}

	return other
}

// OrSome returns the inner value T if present, otherwise it returns provided value.
func (v Value[T]) OrSome(value T) T {
	if v.IsSome() {
		return v.inner
	}

	return value
}

// OrZero returns the inner value T if present, otherwise it returns zero value.
func (v Value[T]) OrZero() T {
	return v.inner
}

// OrElse returns the optional value v if it is [Some], otherwise it calls provided function and returns its result.
func (v Value[T]) OrElse(value func() Value[T]) Value[T] {
	if v.IsSome() {
		return v
	}

	return value()
}

// OrElseSome returns the inner value T if it is [Some], otherwise it calls provided function and returns its result.
func (v Value[T]) OrElseSome(value func() T) T {
	if v.IsSome() {
		return v.inner
	}

	return value()
}

// Reset resets the optional value v to [None].
func (v *Value[T]) Reset() {
	*v = None[T]()
}

// Take returns a copy of optional value v and resets it to [None].
func (v *Value[T]) Take() Value[T] {
	result := *v
	v.Reset()

	return result
}

// Replace returns a copy of optional value v and sets it to [Some] with the given value.
func (v *Value[T]) Replace(value T) Value[T] {
	result := *v
	*v = Some(value)

	return result
}

// CopyPtr returns a pointer to a copy of the inner value T if present, otherwise it returns nil.
func (v Value[T]) CopyPtr() *T {
	if v.IsNone() {
		return nil
	}

	return &v.inner
}
