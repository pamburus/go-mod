package qb

import (
	"database/sql"
	"strings"
)

// Builder is an interface for building SQL queries.
type Builder interface {
	AppendByte(byte)
	AppendString(string)
	AppendArg(any)
}

// FromItem is an abstract SQL from item.
type FromItem interface {
	BuildFromItem(Builder) error
}

// Build builds the SQL query.
func Build(q Statement) (string, []any, error) {
	var b StandardBuilder

	err := q.Build(&b, DefaultStatementOptions())
	if err != nil {
		return "", nil, err
	}

	sql, args := b.Result()

	return sql, args, nil
}

// ---

// StandardBuilder is a standard builder for SQL query.
type StandardBuilder struct {
	sql  strings.Builder
	args []any
}

// AppendByte appends a byte to the SQL query.
func (b *StandardBuilder) AppendByte(val byte) {
	_ = b.sql.WriteByte(val)
}

// AppendString appends a string to the SQL query.
func (b *StandardBuilder) AppendString(val string) {
	_, _ = b.sql.WriteString(val)
}

// AppendArg appends an argument to the SQL query.
func (b *StandardBuilder) AppendArg(arg any) {
	b.args = append(b.args, arg)

	switch arg := arg.(type) {
	case sql.NamedArg:
		b.AppendByte('@')
		b.AppendString(arg.Name)
	default:
		b.AppendByte('?')
	}
}

// Result returns the SQL query and its arguments.
func (b *StandardBuilder) Result() (string, []any) {
	return b.sql.String(), b.args
}
