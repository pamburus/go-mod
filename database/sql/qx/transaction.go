package qx

import (
	"context"
	"database/sql"
	"iter"

	"github.com/pamburus/go-mod/database/sql/qb"
)

type Transaction interface {
	// Exec executes a query that does not return rows.
	Exec(context.Context, qb.Query) (sql.Result, error)

	// Query executes a query that returns rows.
	// The implementation must return [ErrResult] along with the error if the error is not nil.
	Query(context.Context, qb.Query) iter.Seq2[Result, error]

	// QueryRow executes a query that returns a single row.
	QueryRow(context.Context, qb.Query) Row

	// Transact executes a function in a transaction or nested transaction.
	Transact(context.Context, func(context.Context, Transaction) error) error

	sealedTransaction()
}

// TransactionStub returns a stub implementation of Transaction.
func TransactionStub() Transaction {
	return transactionStubInstance
}

// ---

var transactionStubInstance = &transactionStub{}

// ---

type transactionStub struct{}

func (transactionStub) Exec(context.Context, qb.Query) (sql.Result, error) {
	return nil, errNotImplemented
}

func (transactionStub) Query(context.Context, qb.Query) iter.Seq2[Result, error] {
	return func(yield func(Result, error) bool) {
		yield(ResultStub(), errNotImplemented)
	}
}

func (transactionStub) QueryRow(context.Context, qb.Query) Row {
	return errRow{errNotImplemented}
}

func (transactionStub) Transact(context.Context, func(context.Context, Transaction) error) error {
	return errNotImplemented
}

func (transactionStub) sealedTransaction() {}

// ---

var (
	_ Transaction = transactionStub{}
)
