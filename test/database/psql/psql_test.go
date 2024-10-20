package psql_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/pamburus/go-mod/database/sql/qb"
	"github.com/pamburus/go-mod/database/sql/qb/qbpgx"
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
	db, err := pgx.Connect(ctx, "postgres://platform_cloudconnsvc_db_user:1b986da4e860a8a16e735dbc59bef075@10.235.158.60:5432/azureneo_platform_cloudconnsvc_pi_tmp_2?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close(ctx)

	query, args, err := qb.Build(
		qbpgx.NewBuilder,
		qb.Select(
			qb.Star(),
			qb.Arg("bb"),
		).
			From(qb.Table("feature")).
			Where(qb.And(
				qb.NotEqual(qb.Column("id"), qb.Arg("aa")),
				qb.NotEqual(qb.Column("type"), qb.Arg("connection")),
			)),
	)
	if err != nil {
		t.Fatal(err)
	}

	rows, err := db.Query(ctx, query, args...)
	if err != nil {
		t.Fatal(err)
	}

	for rows.Next() {
		var id string
		var version int
		var typ string
		var checksum sql.NullString
		var x string

		err = rows.Scan(&id, &version, &typ, &checksum, &x)
		if err != nil {
			t.Fatal(err)
		}

		if !checksum.Valid {
			checksum.String = "<NULL>"
		}

		t.Log(id, version, typ, checksum.String, x)
	}
}
