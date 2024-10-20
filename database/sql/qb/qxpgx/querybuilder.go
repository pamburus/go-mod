package qxpgx

import (
	"database/sql"
	"fmt"
	"iter"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/pamburus/go-mod/database/sql/qb"
)

// queryBuilder is a  queryBuilder for SQL query.
type queryBuilder struct {
	sql   strings.Builder
	args  []any
	named pgx.StrictNamedArgs
}

// AppendByte appends a byte to the SQL query.
func (b *queryBuilder) AppendByte(val byte) {
	_ = b.sql.WriteByte(val)
}

// AppendString appends a string to the SQL query.
func (b *queryBuilder) AppendString(val string) {
	_, _ = b.sql.WriteString(val)
}

// AppendArg appends an argument to the SQL query.
func (b *queryBuilder) AppendArg(arg any) error {
	var named iter.Seq2[string, any]

	switch arg := arg.(type) {
	case sql.NamedArg:
		named = func(yield func(string, any) bool) {
			yield(arg.Name, arg.Value)
		}
	case pgx.NamedArgs:
		named = maps.All(arg)
	case pgx.StrictNamedArgs:
		named = maps.All(arg)
	default:
		return b.appendPositionalArg(arg, true)
	}

	for name, value := range named {
		err := b.appendNamedArg(name, value, true)
		if err != nil {
			return err
		}
	}

	return nil
}

// AppendArg appends an argument to the SQL query.
func (b *queryBuilder) AppendRawArgs(args ...any) error {
	for _, arg := range args {
		var named iter.Seq2[string, any]

		switch arg := arg.(type) {
		case sql.NamedArg:
			named = func(yield func(string, any) bool) {
				yield(arg.Name, arg.Value)
			}
		case pgx.NamedArgs:
			named = maps.All(arg)
		case pgx.StrictNamedArgs:
			named = maps.All(arg)
		default:
			return ErrPositionalArgsNotAllowed
		}

		for name, value := range named {
			err := b.appendNamedArg(name, value, false)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Result returns the SQL query and its arguments.
func (b *queryBuilder) Result() (string, []any) {
	args := slices.Clip(b.args)
	if len(b.named) > 0 {
		args = []any{b.named}
	}

	return b.sql.String(), args
}

func (b *queryBuilder) appendNamedArg(name string, value any, placeholder bool) error {
	if len(b.args) != 0 {
		return ErrMixingNamedAndPositionalArgsNotAllowed
	}

	if _, found := b.named[name]; found {
		return fmt.Errorf("%w %q", qb.ErrDuplicateNamedArg, name)
	}

	b.named[name] = value

	if placeholder {
		b.AppendByte('@')
		b.AppendString(name)
	}

	return nil
}

func (b *queryBuilder) appendPositionalArg(value any, placeholder bool) error {
	if len(b.named) != 0 {
		return ErrMixingNamedAndPositionalArgsNotAllowed
	}

	b.args = append(b.args, value)
	if placeholder {
		b.AppendByte('$')
		b.AppendString(strconv.Itoa(len(b.args)))
	}

	return nil
}

// ---

var _ qb.Builder = &queryBuilder{}
