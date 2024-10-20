// Package result provides a type [Result][T] that represents a value of type T or an error.
package result

import (
	"errors"
	"iter"
	"slices"
)

// ---

// New constructs a new [Result].
func New[T any](value T, err error) Result[T] {
	return Result[T]{value, err}
}

// Ok returns a [Result] with the provided value and nil error.
func Ok[T any](value T) Result[T] {
	return New(value, nil)
}

// NewErr returns a [Result] with the provided error and zero value.
func NewErr[T any](err error) Result[T] {
	return Result[T]{err: err}
}

// NewPanic returns a [Result] with the provided panic value wrapped into an error with [WrapPanic].
func NewPanic[T any](value any) Result[T] {
	return NewErr[T](WrapPanic(value))
}

// Get returns a [Result] with the value and error returned by the provided function.
// If the function panics, the panic value is wrapped into an error with [WrapPanic].
func Get[T any](f func() (T, error)) (result Result[T]) {
	defer func() {
		if pv := recover(); pv != nil {
			result = NewPanic[T](pv)
		}
	}()

	return New(f())
}

// Unwrap returns the inner value of type T and the error.
// If the result is a wrapped panic error, it panics with the original panic value.
// If the result is any other error, the value may still be returned
// if it was originally stored along with the error using [New].
func Unwrap[T any](v Result[T]) (T, error) {
	return v.Unwrap()
}

// UnwrapNoPanic returns the inner value of type T and the error.
// If the result is a wrapped panic error, it is returned as an error.
func UnwrapNoPanic[T any](v Result[T]) (T, error) {
	return v.UnwrapNoPanic()
}

// Err returns the error if the result is an error or a panic.
// If the result is a success, it returns nil.
func Err[T any](result Result[T]) error {
	return result.Err()
}

// Value returns the value and a boolean indicating whether the result is a success.
func Value[T any](result Result[T]) (T, bool) {
	return result.Value()
}

// ValueOrZero returns the value if the result is a success or the zero value of type T.
func ValueOrZero[T any](result Result[T]) T {
	return result.ValueOrZero()
}

// IsOk returns true if the result is a success.
func IsOk[T any](result Result[T]) bool {
	return result.IsOk()
}

// IsErr returns true if the result is an error, including wrapped panic error.
func IsErr[T any](result Result[T]) bool {
	return result.IsErr()
}

// IsPanic returns true if the result is a wrapped panic error.
func IsPanic[T any](result Result[T]) bool {
	return result.IsPanic()
}

// Map transforms [Result][T] to [Result][U] using the given function.
// Map applies the function only if the result is a success.
// If the result is an error, the function is not called and the returned result contains no value.
// So, the original value is discarded in this case.
func Map[T, U any, F ~func(T) U](result Result[T], f F) Result[U] {
	value, err := result.UnwrapNoPanic()
	if err != nil {
		return NewErr[U](err)
	}

	return Ok(f(value))
}

// MapErr converts the error of the given result using the provided function.
func MapErr[T any, F ~func(error) error](result Result[T], f F) Result[T] {
	value, err := result.UnwrapNoPanic()
	if err != nil {
		return New(value, f(err))
	}

	return Ok(value)
}

// Flatten flattens optional [Result][[Result][T]] to [Result][T].
func Flatten[T any](result Result[Result[T]]) Result[T] {
	value, err := result.UnwrapNoPanic()
	if err != nil {
		return New(ValueOrZero(value), err)
	}

	return value
}

// FlatMap transforms [Result][T] to [Result][U] using the given function that returns a value and an error.
// FlatMap applies the function only if the result is a success.
// If the result is an error, the function is not called and the returned result contains no value.
// So, the original value is discarded in this case.
func FlatMap[T, U any, F ~func(T) (U, error)](result Result[T], f F) Result[U] {
	return FlatMapResult(result, func(value T) Result[U] {
		return New(f(value))
	})
}

// FlatMapResult transforms [Result][T] to [Result][U] using the given function that returns a result.
// FlatMapResult applies the function only if the result is a success.
// If the result is an error, the function is not called and the returned result contains no value.
// So, the original value is discarded in this case.
func FlatMapResult[T, U any, F ~func(T) Result[U]](result Result[T], f F) Result[U] {
	value, err := result.UnwrapNoPanic()
	if err != nil {
		return NewErr[U](err)
	}

	return f(value)
}

// Join collects all values from the given results and returns them as a result over a slice.
func Join[T any](results ...Result[T]) Result[[]T] {
	return join(make([]T, 0, len(results)), slices.Values(results))
}

// JoinSeq collects all values from the given sequence and returns them as a result over a slice.
func JoinSeq[T any](seq iter.Seq[Result[T]]) Result[[]T] {
	return join(nil, seq)
}

// FromSeq2 converts a sequence of value and error pairs to a sequence of results.
func FromSeq2[T any](seq iter.Seq2[T, error]) iter.Seq[Result[T]] {
	return func(yield func(Result[T]) bool) {
		for value, err := range seq {
			yield(New(value, err))
		}
	}
}

// UnwrapCollect unwraps each result in the given sequence and collects all values
// by appending them to the provided slice.
// If an error is encountered, it is returned immediately.
// If a wrapped panic error is encountered, the function panics with the original panic value.
func UnwrapCollect[Slice ~[]T, T any](s Slice, seq iter.Seq[Result[T]]) ([]T, error) {
	for result := range seq {
		val, err := result.Unwrap()
		if err != nil {
			return nil, err
		}

		s = append(s, val)
	}

	return s, nil
}

// ---

// Result is either a success with a value of type T or an error.
//
// In some cases, the value may be stored along with the error.
//
// The error may also be a wrapped panic value created by [WrapPanic] or [NewPanic].
// This will force [Unwrap] to panic.
// In this case it is considered that the result is a panic.
type Result[T any] struct {
	inner T
	err   error
}

// Unwrap returns the inner value of type T and the error.
// If the result is a panic, it panics with the original panic value.
// If the result is an error, the value may still be returned
// if it was originally stored along with the error.
func (v Result[T]) Unwrap() (T, error) {
	return v.inner, RecallPanic(v.err)
}

// UnwrapNoPanic returns the inner value of type T and the error.
// If the result is a panic, it is returned as an error.
func (v Result[T]) UnwrapNoPanic() (T, error) {
	return v.inner, v.err
}

// IsOk returns true if the result is a success.
func (v Result[T]) IsOk() bool {
	return v.err == nil
}

// IsErr returns true if the result is an error or panic.
func (v Result[T]) IsErr() bool {
	return !v.IsOk()
}

// IsPanic returns true if the result is a wrapped panic error.
func (v Result[T]) IsPanic() bool {
	_, ok := UnwrapPanic(v.err)

	return ok
}

// Err returns the error.
// If the result is [Ok], it returns nil.
// If the result is a panic, it is returned as an error.
// To panic with the original panic value, use RecallPanic.
func (v Result[T]) Err() error {
	return v.err
}

// Value returns the value and a boolean indicating whether the result is a success.
func (v Result[T]) Value() (T, bool) {
	value, err := v.UnwrapNoPanic()

	return value, err == nil
}

// ValueOrZero returns the value if the result is a success or the zero value of type T.
func (v Result[T]) ValueOrZero() T {
	value, _ := v.Value()

	return value
}

// ---

func join[T any](values []T, results iter.Seq[Result[T]]) Result[[]T] {
	var errs []error

	for result := range results {
		value, err := result.UnwrapNoPanic()
		if err != nil {
			errs = append(errs, err)
		} else {
			values = append(values, value)
		}
	}

	return New(values, errors.Join(errs...))
}
