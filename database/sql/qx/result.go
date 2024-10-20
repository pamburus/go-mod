package qx

import (
	"database/sql"
	"iter"
)

// Result is a database query result.
type Result interface {
	sql.Result

	// Columns returns the column names.
	Columns() []string

	// Rows returns a sequence of rows.
	// The implementation must return [ErrRow] along with the error if the error is not nil.
	Rows() iter.Seq2[Row, error]

	sealed()
}

// Row is a row in a database query result.
type Row interface {
	Scan(dest ...any) error
}

// ErrResult returns a database query result stub that returns the provided error.
func ErrResult(err error) Result {
	return errResult{err}
}

// ErrRow returns a database query row stub that returns the provided error.
func ErrRow(err error) Row {
	return errRow{err}
}

// ResultStub returns a stub implementation of Result.
func ResultStub() Result {
	return resultStubInstance
}

// ---

var resultStubInstance = &resultStub{}

// ---

type resultStub struct{}

func (resultStub) LastInsertId() (int64, error) {
	return 0, errNotImplemented
}

func (resultStub) RowsAffected() (int64, error) {
	return 0, errNotImplemented
}

func (resultStub) Columns() []string {
	panic(errNotImplemented)
}

func (resultStub) Rows() iter.Seq2[Row, error] {
	return errResult{errNotImplemented}.Rows()
}

func (resultStub) sealed() {}

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

func (s errRow) Scan(...any) error {
	return s.err
}

// ---

var (
	_ Result = resultStub{}
	_ Result = errResult{}
	_ Row    = errRow{}
)
