package qx

import (
	"context"
	"database/sql"
	"iter"

	"github.com/pamburus/go-mod/database/sql/qb"
)

type Database interface {
	Exec(ctx context.Context, statement qb.Statement) (sql.Result, error)
	Query(ctx context.Context, statement qb.Statement) iter.Seq2[Result, error]
	QueryRow(ctx context.Context, statement qb.Statement) Row
	sealed()
}

type Result interface {
	sql.Result
	Columns() []string
	Rows() iter.Seq2[Row, error]
	sealed()
}

type Row interface {
	Scan(dest ...any) error
}

func ErrResult(err error) Result {
	return errResult{err}
}

func ErrRow(err error) Row {
	return errRow{err}
}

// ---

type DatabaseStub struct{}

func (DatabaseStub) Exec(ctx context.Context, statement qb.Statement) (sql.Result, error) {
	return nil, errNotImplemented
}

func (DatabaseStub) Query(ctx context.Context, statement qb.Statement) iter.Seq2[Result, error] {
	return func(yield func(Result, error) bool) {
		yield(ResultStub{}, errNotImplemented)
	}
}

func (DatabaseStub) QueryRow(ctx context.Context, statement qb.Statement) Row {
	return errRow{errNotImplemented}
}

func (DatabaseStub) sealed() {}

// ---

type ResultStub struct{}

func (ResultStub) LastInsertId() (int64, error) {
	return 0, errNotImplemented
}

func (ResultStub) RowsAffected() (int64, error) {
	return 0, errNotImplemented
}

func (ResultStub) Columns() []string {
	return nil
}

func (ResultStub) Rows() iter.Seq2[Row, error] {
	return errResult{errNotImplemented}.Rows()
}

func (ResultStub) sealed() {}

// ---

type errResult struct {
	err error
}

func (s errResult) LastInsertId() (int64, error) {
	return 0, s.err
}

func (s errResult) RowsAffected() (int64, error) {
	return 0, s.err
}

func (errResult) Columns() []string {
	return nil
}

func (s errResult) Rows() iter.Seq2[Row, error] {
	return func(yield func(Row, error) bool) {
		yield(errRow{s.err}, s.err)
	}
}

func (errResult) sealed() {}

// ---

type errRow struct {
	err error
}

func (s errRow) Scan(dest ...any) error {
	return s.err
}
