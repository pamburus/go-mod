package qx

import "errors"

// IsNotImplemented reports whether an error is an "not implemented" error.
func IsNotImplemented(err error) bool {
	return errors.Is(err, errNotImplemented)
}

// ---

var errNotImplemented = errors.New("not implemented")
