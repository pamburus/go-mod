package qxpgx

import (
	"context"
	"database/sql"
	"iter"

	"github.com/jackc/pgx/v5"

	"github.com/pamburus/go-mod/database/sql/qb"
	"github.com/pamburus/go-mod/database/sql/qb/qx"
	"github.com/pamburus/go-mod/database/sql/qb/qxpgx/backend"
)

// transactionInterface is a subset of qx.Transaction that is implemented in this package.
type transactionInterface interface {
	Exec(context.Context, qb.Statement) (sql.Result, error)
	Query(context.Context, qb.Statement) iter.Seq2[qx.Result, error]
	QueryRow(context.Context, qb.Statement) qx.Row
	Transact(context.Context, func(context.Context, qx.Transaction) error) error
}

// ---

func newTransaction(connection backend.Transaction) *transaction {
	return &transaction{qx.TransactionStub{}, transactionImpl{connection}}
}

// ---

type transaction struct {
	qx.TransactionStub
	transactionImpl
}

func (t *transaction) Exec(ctx context.Context, statement qb.Statement) (sql.Result, error) {
	return t.transactionImpl.Exec(ctx, statement)
}

func (t *transaction) Query(ctx context.Context, statement qb.Statement) iter.Seq2[qx.Result, error] {
	return t.transactionImpl.Query(ctx, statement)
}

func (t *transaction) QueryRow(ctx context.Context, statement qb.Statement) qx.Row {
	return t.transactionImpl.QueryRow(ctx, statement)
}

func (t *transaction) Transact(ctx context.Context, fn func(context.Context, qx.Transaction) error) error {
	return t.transactionImpl.Transact(ctx, fn)
}

// ---

type transactionImpl struct {
	connection backend.Transaction
}

func (t *transactionImpl) Exec(ctx context.Context, statement qb.Statement) (sql.Result, error) {
	sql, args, err := t.build(statement)
	if err != nil {
		return nil, err
	}

	commandTag, err := t.connection.Exec(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return sqlResult(commandTag), nil
}

func (t *transactionImpl) Query(ctx context.Context, statement qb.Statement) iter.Seq2[qx.Result, error] {
	fail := func(err error) iter.Seq2[qx.Result, error] {
		return func(yield func(qx.Result, error) bool) {
			yield(qx.ErrResult(err), err)
		}
	}

	sql, args, err := t.build(statement)
	if err != nil {
		return fail(err)
	}

	return func(yield func(qx.Result, error) bool) {
		err := func() error {
			rows, err := t.connection.Query(ctx, sql, args...)
			if err != nil {
				return err
			}
			defer rows.Close()

			fields := rows.FieldDescriptions()
			columns := make([]string, 0, len(fields))
			for _, field := range fields {
				columns = append(columns, string(field.Name))
			}

			yield(result{qx.ResultStub{}, rows, columns}, nil)

			return nil
		}()
		if err != nil {
			yield(qx.ErrResult(err), err)
		}
	}

}

func (t *transactionImpl) QueryRow(ctx context.Context, statement qb.Statement) qx.Row {
	sql, args, err := t.build(statement)
	if err != nil {
		return qx.ErrRow(err)
	}

	return t.connection.QueryRow(ctx, sql, args...)
}

func (t *transactionImpl) Transact(ctx context.Context, fn func(context.Context, qx.Transaction) error) error {
	return pgx.BeginFunc(ctx, t.connection, func(tx pgx.Tx) error {
		return fn(ctx, newTransaction(tx))
	})
}

func (t *transactionImpl) build(statement qb.Statement) (string, []any, error) {
	var b queryBuilder
	err := statement.BuildStatement(&b, qb.DefaultStatementOptions())
	if err != nil {
		return "", nil, err
	}

	return b.sql.String(), b.args, nil
}

// ---

var (
	_ transactionInterface = qx.Transaction(nil)
	_ transactionInterface = &transactionImpl{}
	_ qx.Transaction       = &transaction{}
)
