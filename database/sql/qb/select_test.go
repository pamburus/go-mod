package qb_test

import (
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/pamburus/go-mod/database/sql/qb"
	"github.com/pamburus/go-mod/database/sql/qb/qbpgx"
	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	q := qb.
		Select(
			qb.Column("id"),
			qb.Column("name").As("alias"),
		).
		From(qb.Table("user")).
		Where(qb.Equal(
			qb.Column("id"),
			qb.Arg(1),
		)).
		Limit(3)

	builder := qbpgx.NewBuilder
	sql, args, err := qb.Build(builder, q)
	require.NoError(t, err)
	require.Equal(t, `SELECT id, name AS alias FROM user WHERE id = $1 LIMIT 3`, sql)
	require.Equal(t, []any{1}, args)
}

func TestSelectSubquery(t *testing.T) {
	q := qb.
		Select(
			qb.Column("id"),
			qb.Column("name").As("alias"),
		).
		From(
			qb.Select(qb.Star()).From(qb.Table("user")).As("_sq1"),
		).
		Where(qb.Equal(
			qb.Column("id"),
			qb.NamedArg("id1", 1),
		)).
		Limit(3)

	builder := qbpgx.NewBuilder
	sql, args, err := qb.Build(builder, q)
	require.NoError(t, err)
	require.Equal(t, `SELECT id, name AS alias FROM (SELECT * FROM user) AS _sq1 WHERE id = @id1 LIMIT 3`, sql)
	require.Equal(t, []any{pgx.StrictNamedArgs{"id1": 1}}, args)
}
