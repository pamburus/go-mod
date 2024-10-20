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

// queryBuilderInterface is a subset of the qb.Builder interface that is supported by this package.
type queryBuilderInterface interface {
	AppendByte(byte)
	AppendString(string)
	AppendArg(any) error
	AppendRawExpr(expr string, args ...any) error
}

// ---

func newQueryBuilder() queryBuilder {
	return queryBuilder{Builder: qb.BuilderStub()}
}

// ---

type queryBuilder struct {
	qb.Builder
	queryBuilderImpl
}

func (b *queryBuilder) AppendByte(val byte) {
	b.queryBuilderImpl.AppendByte(val)
}

func (b *queryBuilder) AppendString(val string) {
	b.queryBuilderImpl.AppendString(val)
}

func (b *queryBuilder) AppendArg(arg any) error {
	return b.queryBuilderImpl.AppendArg(arg)
}

func (b *queryBuilder) AppendRawExpr(expr string, args ...any) error {
	return b.queryBuilderImpl.AppendRawExpr(expr, args...)
}

// ---

type queryBuilderImpl struct {
	sql   strings.Builder
	args  []any
	named pgx.StrictNamedArgs
}

func (b *queryBuilderImpl) AppendByte(val byte) {
	_ = b.sql.WriteByte(val)
}

func (b *queryBuilderImpl) AppendString(val string) {
	_, _ = b.sql.WriteString(val)
}

func (b *queryBuilderImpl) AppendArg(arg any) error {
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

func (b *queryBuilderImpl) AppendRawExpr(expr string, args ...any) error {
	b.AppendString(expr)

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

func (b *queryBuilderImpl) result() (string, []any, error) {
	args := slices.Clip(b.args)
	if len(b.named) > 0 {
		args = []any{b.named}
	}

	return b.sql.String(), args, nil
}

func (b *queryBuilderImpl) appendNamedArg(name string, value any, placeholder bool) error {
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

func (b *queryBuilderImpl) appendPositionalArg(value any, placeholder bool) error {
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

var (
	_ queryBuilderInterface = qb.Builder(nil)
	_ queryBuilderInterface = &queryBuilderImpl{}
	_ qb.Builder            = &queryBuilder{}
)
