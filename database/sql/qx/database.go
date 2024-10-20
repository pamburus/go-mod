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

// DatabaseStub returns a stub implementation of Database.
func DatabaseStub() Database {
	return databaseStubInstance
}

// ---

var databaseStubInstance = &databaseStub{}

// ---

type databaseStub struct{ transactionStub }

func (databaseStub) TransactWithOptions(context.Context, sql.TxOptions, func(context.Context, Transaction) error) error {
	return errNotImplemented
}

func (databaseStub) sealedDatabase() {}

// ---

var (
	_ Database = databaseStub{}
)
