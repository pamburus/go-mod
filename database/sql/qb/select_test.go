package qb_test

import (
	"testing"

	"github.com/pamburus/go-mod/database/sql/qb"
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
			qb.Value(1),
		)).
		Limit(3)

	sql, args, err := qb.Build(q)
	require.NoError(t, err)
	require.Equal(t, `SELECT id, name AS alias FROM user WHERE id = ? LIMIT ?`, sql)
	require.Equal(t, []any{1, 3}, args)
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
			qb.Value(1),
		)).
		Limit(3)

	sql, args, err := qb.Build(q)
	require.NoError(t, err)
	require.Equal(t, `SELECT id, name AS alias FROM (SELECT * FROM user) AS _sq1 WHERE id = ? LIMIT ?`, sql)
	require.Equal(t, []any{1, 3}, args)
}
