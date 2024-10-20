package result

import (
	"errors"
	"fmt"
)

// WrapPanic wraps a panic value into an error.
func WrapPanic(value any) error {
	return &errPanic{value}
}

// UnwrapPanic unwraps a panic value from an error previously wrapped with [WrapPanic].
func UnwrapPanic(err error) (any, bool) {
	var e *errPanic
	if errors.As(err, &e) {
		return e.value, true
	}

	return nil, false
}

// RecallPanic panics with the original panic value if the error is a wrapped panic returned by [WrapPanic].
// Otherwise it returns the same error.
func RecallPanic(err error) error {
	if val, ok := UnwrapPanic(err); ok {
		panic(val)
	}

	return err
}

// ---

type errPanic struct {
	value any
}

func (e *errPanic) Error() string {
	return fmt.Sprintf("panic: %v", e.value)
}

func (e *errPanic) Unwrap() error {
	if err, ok := e.value.(error); ok {
		return err
	}

	return nil
}
