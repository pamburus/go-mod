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

func New(backend backend.Database) qx.Database {
	return &database{
		qx.DatabaseStub(),
		databaseImpl{
			backend,
			transactionImpl{backend},
		},
	}
}

// ---

// databaseInterface is a subset of qx.Database that is implemented in this package.
type databaseInterface interface {
	Exec(context.Context, qb.Query) (sql.Result, error)
	Query(context.Context, qb.Query) iter.Seq2[qx.Result, error]
	QueryRow(context.Context, qb.Query) qx.Row
	Transact(context.Context, func(context.Context, qx.Transaction) error) error
	TransactWithOptions(context.Context, sql.TxOptions, func(context.Context, qx.Transaction) error) error
}

// ---

type database struct {
	qx.Database
	databaseImpl
}

func (d *database) Exec(ctx context.Context, query qb.Query) (sql.Result, error) {
	return d.databaseImpl.Exec(ctx, query)
}

func (d *database) Query(ctx context.Context, query qb.Query) iter.Seq2[qx.Result, error] {
	return d.databaseImpl.Query(ctx, query)
}

func (d *database) QueryRow(ctx context.Context, query qb.Query) qx.Row {
	return d.databaseImpl.QueryRow(ctx, query)
}

func (d *database) Transact(ctx context.Context, fn func(context.Context, qx.Transaction) error) error {
	return d.databaseImpl.Transact(ctx, fn)
}

func (d *database) TransactWithOptions(ctx context.Context, opts sql.TxOptions, fn func(context.Context, qx.Transaction) error) error {
	return d.databaseImpl.TransactWithOptions(ctx, opts, fn)
}

// ---

type databaseImpl struct {
	backend backend.Database
	transactionImpl
}

func (d *databaseImpl) TransactWithOptions(ctx context.Context, opts sql.TxOptions, fn func(context.Context, qx.Transaction) error) error {
	var innerOpts pgx.TxOptions

	switch opts.Isolation {
	case sql.LevelDefault:
		break
	case sql.LevelSerializable:
		innerOpts.IsoLevel = pgx.Serializable
	case sql.LevelRepeatableRead:
		innerOpts.IsoLevel = pgx.RepeatableRead
	case sql.LevelReadCommitted:
		innerOpts.IsoLevel = pgx.ReadCommitted
	case sql.LevelReadUncommitted:
		innerOpts.IsoLevel = pgx.ReadUncommitted
	default:
		return errUnsupportedIsolationLevel(opts.Isolation)
	}

	if opts.ReadOnly {
		innerOpts.AccessMode = pgx.ReadOnly
	} else {
		innerOpts.AccessMode = pgx.ReadWrite
	}

	return pgx.BeginTxFunc(ctx, d.backend, innerOpts, func(tx pgx.Tx) error {
		return fn(ctx, newTransaction(tx))
	})
}

// ---

var (
	_ databaseInterface = qx.Database(nil)
	_ databaseInterface = &databaseImpl{}
	_ qx.Database       = &database{}
)
