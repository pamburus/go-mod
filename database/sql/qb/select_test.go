package qb_test

import (
	"testing"

	"github.com/pamburus/go-mod/database/sql/qb"
	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	q := qb.Select(
		qb.Column{Name: "id"},
		qb.Column{Name: "type"},
	).
		From(qb.Table{Name: "users"}).
		Where(qb.Equal(
			qb.Column{Name: "id"},
			qb.Value(1),
		)).
		Limit(3)

	sql, args, err := qb.Build(q)
	require.NoError(t, err)
	require.Equal(t, "SELECT id, type FROM users WHERE id = ? LIMIT ?", sql)
	require.Equal(t, []any{1, 3}, args)
}
