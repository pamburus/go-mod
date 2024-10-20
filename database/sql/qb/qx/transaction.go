package qx

import (
	"context"
	"database/sql"
	"iter"

	"github.com/pamburus/go-mod/database/sql/qb"
)

type Transaction interface {
	// Exec executes a statement that does not return rows.
	Exec(context.Context, qb.Statement) (sql.Result, error)

	// Query executes a statement that returns rows.
	// The implementation must return [ErrResult] along with the error if the error is not nil.
	Query(context.Context, qb.Statement) iter.Seq2[Result, error]

	// QueryRow executes a statement that returns a single row.
	QueryRow(context.Context, qb.Statement) Row

	// Transact executes a function in a transaction or nested transaction.
	Transact(context.Context, func(context.Context, Transaction) error) error

	sealedTransaction()
}

// ---
// TransactionStub is a stub implementation of Transaction that returns "not implemented" error in all methods.
type TransactionStub struct{}

// Exec returns "not implemented" error.
func (TransactionStub) Exec(context.Context, qb.Statement) (sql.Result, error) {
	return nil, errNotImplemented
}

// Query returns "not implemented" error.
func (TransactionStub) Query(context.Context, qb.Statement) iter.Seq2[Result, error] {
	return func(yield func(Result, error) bool) {
		yield(ResultStub{}, errNotImplemented)
	}
}

// QueryRow returns "not implemented" error.
func (TransactionStub) QueryRow(context.Context, qb.Statement) Row {
	return errRow{errNotImplemented}
}

// Transact returns "not implemented" error.
func (TransactionStub) Transact(context.Context, func(context.Context, Transaction) error) error {
	return errNotImplemented
}

func (TransactionStub) sealedTransaction() {}

// ---

var (
	_ Transaction = TransactionStub{}
)
