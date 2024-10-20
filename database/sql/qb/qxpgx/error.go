package qxpgx

import "errors"

var (
	ErrPositionalArgsNotAllowed               = errors.New("qxpgx: positional arguments are not allowed in raw expressions")
	ErrMixingNamedAndPositionalArgsNotAllowed = errors.New("qxpgx: mixing named and positional arguments is not allowed")
	ErrLastInsertIdNotSupported               = errors.New("qxpgx: LastInsertId is not supported")
	ErrRowHasBeenInvalidated                  = errors.New("qxpgx: row has been invalidated")
)
