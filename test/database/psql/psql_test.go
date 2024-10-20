package psql_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/pamburus/go-mod/database/sql/qb"
	"github.com/pamburus/go-mod/database/sql/qb/qxpgx"
	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/result"
)

func TestPGXStd(t *testing.T) {
	db, err := sql.Open("pgx", "postgres://platform_cloudconnsvc_db_user:1b986da4e860a8a16e735dbc59bef075@10.235.158.60:5432/azureneo_platform_cloudconnsvc_pi_tmp_2?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM feature WHERE "type" = $1;`, "connection")
	if err != nil {
		t.Fatal(err)
	}

	for rows.Next() {
		var id string
		var version int
		var typ string
		var checksum *string

		err = rows.Scan(&id, &version, &typ, &checksum)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(id, version, typ, checksum)
	}
}

func TestPGX(t *testing.T) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://platform_cloudconnsvc_db_user:1b986da4e860a8a16e735dbc59bef075@10.235.158.60:5432/azureneo_platform_cloudconnsvc_pi_tmp_2?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	db := qxpgx.New(pool)

	query := qb.Select(qb.Star()).
		From(qb.Table("feature")).
		Where(qb.And(
			qb.NotEqual(qb.Column("id"), qb.Arg("aa")),
			qb.NotEqual(qb.Column("type"), qb.Arg("connection")),
		))

	for resultSet, err := range db.Query(ctx, query) {
		if err != nil {
			t.Fatal(err)
		}

		t.Log(resultSet.Columns())

		var id string
		var version int
		var typ string
		var checksum sql.NullString

		for i, row := range gi.Enumerate(result.FromSeq2(resultSet.Rows())) {
			row, err := row.Unwrap()
			if err != nil {
				t.Fatal(err)
			}

			err = row.Scan(&id, &version, &typ, &checksum)
			if err != nil {
				t.Fatal(err)
			}

			if !checksum.Valid {
				checksum.String = "<NULL>"
			}

			t.Log("#", i, id, version, typ, checksum.String)
		}
	}

	query = qb.Select(qb.Raw("now()"))

	for result, err := range db.Query(ctx, query) {
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result.Columns())
		fields := make([]any, len(result.Columns()))
		ptr := make([]any, len(result.Columns()))
		for i := range fields {
			ptr[i] = &fields[i]
		}

		for row, err := range result.Rows() {
			err = row.Scan(ptr...)
			if err != nil {
				t.Fatal(err)
			}

			t.Log(fields...)
		}
	}
}

// ---

type jointStatement struct {
	statements []qb.Statement
}

func (j jointStatement) BuildStatement(b qb.Builder, options qb.StatementOptions) error {
	for _, statement := range j.statements {
		err := statement.BuildStatement(b, options)
		if err != nil {
			return err
		}

		b.AppendString(";\n")
	}

	return nil
}

// ---

var _ qb.Statement = jointStatement{}
