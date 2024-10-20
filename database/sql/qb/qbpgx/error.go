package qbpgx

import "errors"

var (
	ErrPositionalArgsNotAllowed               = errors.New("qbpgx: positional arguments are not allowed in raw expressions")
	ErrMixingNamedAndPositionalArgsNotAllowed = errors.New("qbpgx: mixing named and positional arguments is not allowed")
)
