package qx

import (
	"context"
	"database/sql"
	"iter"

	"github.com/pamburus/go-mod/database/sql/qb"
)

type Database interface {
	Exec(ctx context.Context, statement qb.Statement) (sql.Result, error)
	Query(ctx context.Context, statement qb.Statement) iter.Seq2[Row, error]
	QueryRow(ctx context.Context, statement qb.Statement) Row
}

type Row interface {
	Scan(dest ...any) error
}

type CloseFunc func()
