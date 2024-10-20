package qxpgx

import (
	"database/sql"
	"errors"
	"fmt"
)

func IsErrUnsupportedIsolationLevel(err error) bool {
	var e errUnsupportedIsolationLevel

	return errors.As(err, &e)
}

var (
	ErrPositionalArgsNotAllowed               = errors.New("qxpgx: positional arguments are not allowed in raw expressions")
	ErrMixingNamedAndPositionalArgsNotAllowed = errors.New("qxpgx: mixing named and positional arguments is not allowed")
	ErrLastInsertIdNotSupported               = errors.New("qxpgx: LastInsertId is not supported")
	ErrRowHasBeenInvalidated                  = errors.New("qxpgx: row has been invalidated")
)

// ---

type errUnsupportedIsolationLevel sql.IsolationLevel

func (e errUnsupportedIsolationLevel) Error() string {
	return fmt.Sprintf("qxpgx: unsupported isolation level %s", sql.IsolationLevel(e))
}
