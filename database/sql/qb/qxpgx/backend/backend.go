package backend

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Connection interface {
	Transaction
	BeginTx(context.Context, pgx.TxOptions) (pgx.Tx, error)
}

type Transaction interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	SendBatch(context.Context, *pgx.Batch) pgx.BatchResults
	Begin(context.Context) (pgx.Tx, error)
}

// ---

var (
	_ Connection  = &pgx.Conn{}
	_ Connection  = &pgxpool.Pool{}
	_ Transaction = pgx.Tx(nil)
)
