package qxpgx

import (
	"context"
	"database/sql"
	"iter"

	"github.com/jackc/pgx/v5"

	"github.com/pamburus/go-mod/database/sql/qb"
	"github.com/pamburus/go-mod/database/sql/qx"
	"github.com/pamburus/go-mod/database/sql/qxpgx/backend"
)

// transactionInterface is a subset of qx.Transaction that is implemented in this package.
type transactionInterface interface {
	Exec(context.Context, qb.Query) (sql.Result, error)
	Query(context.Context, qb.Query) iter.Seq2[qx.Result, error]
	QueryRow(context.Context, qb.Query) qx.Row
	Transact(context.Context, func(context.Context, qx.Transaction) error) error
}

// ---

func newTransaction(backend backend.Transaction) *transaction {
	return &transaction{qx.TransactionStub(), transactionImpl{backend}}
}

// ---

type transaction struct {
	qx.Transaction
	transactionImpl
}

func (t *transaction) Exec(ctx context.Context, query qb.Query) (sql.Result, error) {
	return t.transactionImpl.Exec(ctx, query)
}

func (t *transaction) Query(ctx context.Context, query qb.Query) iter.Seq2[qx.Result, error] {
	return t.transactionImpl.Query(ctx, query)
}

func (t *transaction) QueryRow(ctx context.Context, query qb.Query) qx.Row {
	return t.transactionImpl.QueryRow(ctx, query)
}

func (t *transaction) Transact(ctx context.Context, fn func(context.Context, qx.Transaction) error) error {
	return t.transactionImpl.Transact(ctx, fn)
}

// ---

type transactionImpl struct {
	backend backend.Transaction
}

func (t *transactionImpl) Exec(ctx context.Context, query qb.Query) (sql.Result, error) {
	sql, args, err := t.build(query)
	if err != nil {
		return nil, err
	}

	commandTag, err := t.backend.Exec(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return sqlResult(commandTag), nil
}

func (t *transactionImpl) Query(ctx context.Context, query qb.Query) iter.Seq2[qx.Result, error] {
	fail := func(err error) iter.Seq2[qx.Result, error] {
		return func(yield func(qx.Result, error) bool) {
			yield(qx.ErrResult(err), err)
		}
	}

	sql, args, err := t.build(query)
	if err != nil {
		return fail(err)
	}

	return func(yield func(qx.Result, error) bool) {
		err := func() error {
			rows, err := t.backend.Query(ctx, sql, args...)
			if err != nil {
				return err
			}
			defer rows.Close()

			fields := rows.FieldDescriptions()
			columns := make([]string, 0, len(fields))
			for _, field := range fields {
				columns = append(columns, string(field.Name))
			}

			yield(result{qx.ResultStub(), rows, columns}, nil)

			return nil
		}()
		if err != nil {
			yield(qx.ErrResult(err), err)
		}
	}

}

func (t *transactionImpl) QueryRow(ctx context.Context, query qb.Query) qx.Row {
	sql, args, err := t.build(query)
	if err != nil {
		return qx.ErrRow(err)
	}

	return t.backend.QueryRow(ctx, sql, args...)
}

func (t *transactionImpl) Transact(ctx context.Context, fn func(context.Context, qx.Transaction) error) error {
	return pgx.BeginFunc(ctx, t.backend, func(tx pgx.Tx) error {
		return fn(ctx, newTransaction(tx))
	})
}

func (t *transactionImpl) build(query qb.Query) (string, []any, error) {
	b := newQueryBuilder()
	err := query.BuildQuery(&b, qb.DefaultQueryOptions())
	if err != nil {
		return "", nil, err
	}

	return b.result()
}

// ---

var (
	_ transactionInterface = qx.Transaction(nil)
	_ transactionInterface = &transactionImpl{}
	_ qx.Transaction       = &transaction{}
)
