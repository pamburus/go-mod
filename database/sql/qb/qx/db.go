package qx

import (
	"context"
	"database/sql"

	"github.com/pamburus/go-mod/database/sql/qb"
)

type Database interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

// ---

func New[D Database](inner D) DB[D] {
	return DB[D]{inner}
}

// ---

type DB[D Database] struct {
	inner D
}

func (r DB[D]) Inner() D {
	return r.inner
}

func (r DB[D]) Exec(ctx context.Context, statement qb.Statement) (sql.Result, error) {
	sql, args, err := qb.Build(statement)
	if err != nil {
		return nil, err
	}

	return r.inner.ExecContext(ctx, sql, args...)
}

func (r DB[D]) Query(ctx context.Context, statement qb.Statement) (*sql.Rows, error) {
	sql, args, err := qb.Build(statement)
	if err != nil {
		return nil, err
	}

	return r.inner.QueryContext(ctx, sql, args...)
}

func (r *DB[D]) QueryRow(ctx context.Context, statement qb.Statement) *sql.Row {
	sql, args, err := qb.Build(statement)
	if err != nil {
		return nil
	}

	return r.inner.QueryRowContext(ctx, sql, args...)
}

func (r *DB[D]) Prepare(ctx context.Context, statement qb.Statement) (*sql.Stmt, []any, error) {
	sql, args, err := qb.Build(statement)
	if err != nil {
		return nil, nil, err
	}

	stmt, err := r.inner.PrepareContext(ctx, sql)
	if err != nil {
		return nil, nil, err
	}

	return stmt, args, nil
}
