package qxpgx

import (
	"iter"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/pamburus/go-mod/database/sql/qb/qx"
)

// ---

type sqlResult pgconn.CommandTag

func (r sqlResult) RowsAffected() (int64, error) {
	return pgconn.CommandTag(r).RowsAffected(), nil
}

func (r sqlResult) LastInsertId() (int64, error) {
	return 0, ErrLastInsertIdNotSupported
}

// ---

type result struct {
	qx.Result
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
