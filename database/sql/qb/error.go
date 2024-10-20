package qb

import "errors"

var (
	ErrDuplicateNamedArg = errors.New("duplicate named argument")
	ErrNotImplemented    = errors.New("not implemented")
)
