package qxpgx

import (
	"context"
	"database/sql"
	"iter"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/pamburus/go-mod/database/sql/qb"
	"github.com/pamburus/go-mod/database/sql/qb/qx"
	"github.com/pamburus/go-mod/database/sql/qb/qxpgx/backend"
)

func New(connection backend.Connection) qx.Database {
	return &database{qx.DatabaseStub{}, connection}
}

// ---

type database struct {
	qx.DatabaseStub
	connection backend.Connection
}

func (d *database) Exec(ctx context.Context, statement qb.Statement) (sql.Result, error) {
	return d.tx().Exec(ctx, statement)
}

func (d *database) Query(ctx context.Context, statement qb.Statement) iter.Seq2[qx.Result, error] {
	return d.tx().Query(ctx, statement)
}

func (d *database) QueryRow(ctx context.Context, statement qb.Statement) qx.Row {
	return d.tx().QueryRow(ctx, statement)
}

func (d *database) Transact(ctx context.Context, fn func(context.Context, qx.Transaction) error) error {
	return d.tx().Transact(ctx, fn)
}

func (d *database) TransactWithOptions(ctx context.Context, opts sql.TxOptions, fn func(context.Context, qx.Transaction) error) error {
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

	return pgx.BeginTxFunc(ctx, d.connection, innerOpts, func(tx pgx.Tx) error {
		return fn(ctx, &transaction{qx.TransactionStub{}, tx})
	})
}

func (d *database) tx() *transaction {
	return &transaction{qx.TransactionStub{}, d.connection}
}

// ---

type transaction struct {
	qx.TransactionStub
	connection backend.Transaction
}

func (d *transaction) Exec(ctx context.Context, statement qb.Statement) (sql.Result, error) {
	sql, args, err := d.build(statement)
	if err != nil {
		return nil, err
	}

	commandTag, err := d.connection.Exec(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return sqlResult(commandTag), nil
}

func (d *transaction) Query(ctx context.Context, statement qb.Statement) iter.Seq2[qx.Result, error] {
	fail := func(err error) iter.Seq2[qx.Result, error] {
		return func(yield func(qx.Result, error) bool) {
			yield(qx.ErrResult(err), err)
		}
	}

	sql, args, err := d.build(statement)
	if err != nil {
		return fail(err)
	}

	return func(yield func(qx.Result, error) bool) {
		err := func() error {
			rows, err := d.connection.Query(ctx, sql, args...)
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

func (d *transaction) QueryRow(ctx context.Context, statement qb.Statement) qx.Row {
	sql, args, err := d.build(statement)
	if err != nil {
		return qx.ErrRow(err)
	}

	return d.connection.QueryRow(ctx, sql, args...)
}

func (d *transaction) Transact(ctx context.Context, fn func(context.Context, qx.Transaction) error) error {
	return pgx.BeginFunc(ctx, d.connection, func(tx pgx.Tx) error {
		return fn(ctx, &transaction{qx.TransactionStub{}, tx})
	})
}

func (d *transaction) build(statement qb.Statement) (string, []any, error) {
	var b queryBuilder
	err := statement.BuildStatement(&b, qb.DefaultStatementOptions())
	if err != nil {
		return "", nil, err
	}

	return b.sql.String(), b.args, nil
}

// ---

type sqlResult pgconn.CommandTag

func (r sqlResult) RowsAffected() (int64, error) {
	return pgconn.CommandTag(r).RowsAffected(), nil
}

func (r sqlResult) LastInsertId() (int64, error) {
	return 0, ErrLastInsertIdNotSupported
}

// ---

type errRow struct {
	err error
}

func (r errRow) Scan(...any) error {
	return r.err
}

// ---

type errResult struct {
	err error
}

func (r errResult) LastInsertId() (int64, error) {
	return 0, r.err
}

func (r errResult) RowsAffected() (int64, error) {
	return 0, r.err
}

func (errResult) Columns() []string {
	return nil
}

func (r errResult) Rows() iter.Seq2[qx.Row, error] {
	return func(yield func(qx.Row, error) bool) {
		yield(errRow{r.err}, r.err)
	}
}

// ---

type result struct {
	qx.ResultStub
	rows    pgx.Rows
	columns []string
}

func (r result) LastInsertId() (int64, error) {
	return 0, ErrLastInsertIdNotSupported
}

func (r result) RowsAffected() (int64, error) {
	return r.rows.CommandTag().RowsAffected(), nil
}

func (r result) Columns() []string {
	return r.columns
}

func (r result) Rows() iter.Seq2[qx.Row, error] {
	return func(yield func(qx.Row, error) bool) {
		defer r.rows.Close()

		for r.rows.Next() {
			row := oneShotRow{row: r.rows}
			if !yield(&row, nil) {
				return
			}
			row.done = true
		}
	}
}

// ---

type oneShotRow struct {
	row  qx.Row
	done bool
}

func (r *oneShotRow) Scan(dest ...any) error {
	if r.done {
		return ErrRowHasBeenInvalidated
	}

	return r.row.Scan(dest...)
}

// ---

var _ qx.Database = &database{}
