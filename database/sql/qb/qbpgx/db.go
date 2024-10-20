package qbpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Connection interface {
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	SendBatch(context.Context, *pgx.Batch) pgx.BatchResults
}

type DB struct {
	inner Connection
}

// ---

var (
	_ Connection = &pgx.Conn{}
	_ Connection = &pgxpool.Pool{}
)
