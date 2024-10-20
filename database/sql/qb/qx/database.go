package qx

import (
	"context"
	"database/sql"
)

// Database is a database accessor.
type Database interface {
	Transaction

	// TransactWithOptions executes a function in a transaction or nested transaction with the provided options.
	TransactWithOptions(context.Context, sql.TxOptions, func(context.Context, Transaction) error) error

	sealedDatabase()
}

// ---

// DatabaseStub is a stub implementation of Database that returns "not implemented" error in all methods.
type DatabaseStub struct{ TransactionStub }

// TransactWithOptions returns "not implemented" error.
func (DatabaseStub) TransactWithOptions(context.Context, sql.TxOptions, func(context.Context, Transaction) error) error {
	return errNotImplemented
}

func (DatabaseStub) sealedDatabase() {}

// ---

var (
	_ Database = DatabaseStub{}
)
