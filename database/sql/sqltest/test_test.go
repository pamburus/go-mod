package sqltest_test

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"

	"github.com/pamburus/go-mod/database/sql/sqltest"
	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres"
)

type BaseTest = *sqltest.Test[*testing.T]

// ---

type Test struct {
	BaseTest
}

func (t Test) ExpectNoError(err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func (t Test) Run(name string, f func(Test)) {
	t.BaseTest.Run(name, func(base BaseTest) {
		f(Test{base})
	})
}

// ---

func TestTest(tt *testing.T) {
	t := Test{sqltest.New(tt, postgres.NewStarter(postgres.Docker()))}

	countRows := func(rows *sql.Rows) int {
		var count int
		for rows.Next() {
			count++
		}

		return count
	}

	t.Run("T1", func(t Test) {
		_, err := t.DB().ExecContext(t.Context(), "CREATE TABLE test (id int)")
		t.ExpectNoError(err)

		_, err = t.DB().ExecContext(t.Context(), "INSERT INTO test (id) VALUES (1)")
		t.ExpectNoError(err)

		t.Run("T2", func(t Test) {
			_, err := t.DB().ExecContext(t.Context(), "INSERT INTO test (id) VALUES (2)")
			t.ExpectNoError(err)

			rows, err := t.DB().QueryContext(t.Context(), "SELECT * FROM test")
			t.ExpectNoError(err)
			defer rows.Close()

			expectedRows := 2
			if count := countRows(rows); count != expectedRows {
				t.Errorf("expected %d rows, got %d", expectedRows, count)
			}
		})

		rows, err := t.DB().QueryContext(t.Context(), "SELECT * FROM test")
		t.ExpectNoError(err)
		defer rows.Close()

		if count := countRows(rows); count != 1 {
			t.Errorf("expected 1 row, got %d", count)
		}
	})
}
