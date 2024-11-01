package sqltest

import (
	"cmp"
	"context"
	"database/sql"
	"hash/fnv"
	"net/url"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs"
)

func New[T Base[T]](tt T, starter dbs.Starter) *Test[T] {
	return Build(tt, starter).Done()
}

func Build[T Base[T]](tt T, starter dbs.Starter) TestBuilder[T] {
	return TestBuilder[T]{
		tt,
		starter,
		context.Background(),
		nil,
	}
}

// ---

type TestBuilder[T Base[T]] struct {
	tt      T
	starter dbs.Starter
	ctx     context.Context
	ctxFunc func(T) context.Context
}

func (b TestBuilder[T]) WithContext(ctx context.Context) TestBuilder[T] {
	b.ctx = ctx

	return b
}

func (b TestBuilder[T]) WithContextFunc(ctx func(T) context.Context) TestBuilder[T] {
	b.ctxFunc = ctx

	return b
}

func (b TestBuilder[T]) Done() *Test[T] {
	var ctx context.Context
	if b.ctxFunc == nil {
		ctx = cmp.Or(b.ctx, context.Background())
	} else {
		ctx = b.ctxFunc(b.tt)
	}

	clone := setup(ctx, b.tt, b.starter)
	database := clone(ctx, b.tt, "", "init")

	return &Test[T]{
		b.tt,
		b.tt,
		ctx,
		b.ctxFunc,
		database,
		sync.Once{},
		nil,
		"init",
		clone,
	}
}

// ---

type Base[T testing.TB] interface {
	testing.TB
	Run(name string, f func(T)) bool
}

// ---

type Test[T Base[T]] struct {
	testing.TB
	base     T
	ctx      context.Context
	ctxFunc  func(T) context.Context
	database dbs.Database
	onceDB   sync.Once
	db       *sql.DB
	dbName   string
	clone    cloneFunc
}

func (t *Test[T]) Base() T {
	return t.base
}

func (t *Test[T]) Context() context.Context {
	return t.ctx
}

func (t *Test[T]) DatabaseURL() *url.URL {
	return t.database.URL()
}

func (t *Test[T]) DB() *sql.DB {
	t.onceDB.Do(func() {
		var err error
		t.db, err = t.database.Open(t.ctx)
		if err != nil {
			t.Fatal(err)
		}

		t.Cleanup(func() {
			err := t.db.Close()
			if err != nil {
				t.Logf("[sqltest] failed to close database: %v", err)
			}
		})
	})

	return t.db
}

func (t *Test[T]) Run(name string, f func(*Test[T])) {
	t.base.Run(name, func(tt T) {
		f(t.fork(tt))
	})
}

func (t *Test[T]) fork(tt T) *Test[T] {
	ctx := t.ctx
	if t.ctxFunc != nil {
		ctx = t.ctxFunc(tt)
	}

	dbName := "test_" + hashBase32(fnv.New64(), []byte(tt.Name()))
	database := t.clone(ctx, tt, t.dbName, dbName)

	return &Test[T]{
		tt,
		tt,
		ctx,
		t.ctxFunc,
		database,
		sync.Once{},
		nil,
		dbName,
		t.clone,
	}
}

// ---

func setup(ctx context.Context, tb testing.TB, starter dbs.Starter) cloneFunc {
	var debug atomic.Bool
	debug.Store(envFlag("SQL_TEST_DEBUG"))

	password := hashBase32(fnv.New64(), passwordSalt, []byte(tb.Name()))
	port, err := freePort()
	if err != nil {
		tb.Fatalf("[sqltest] failed to allocate free port: %v", err)
	}

	tb.Logf("[sqltest] start database server on port %d", port)
	server, err := starter.Start(ctx, port, password)
	if err != nil {
		tb.Fatalf("[sqltest] failed to start database server: %v", err)
	}

	tb.Cleanup(func() {
		tb.Logf("[sqltest] stop database server on port %d", port)
		err := server.Stop(ctx)
		if err != nil {
			tb.Logf("[sqltest] failed to stop database server: %v", err)
		}
	})

	return func(ctx context.Context, tb testing.TB, base string, name string) dbs.Database {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		database := server.Database(base)

		for {
			base, err = database.Name(ctx)
			if err != nil {
				tb.Logf("[sqltest] database is not ready: %v", err)

				err := sleep(ctx, time.Second)
				if err != nil {
					tb.Fatalf("[sqltest] failed to wait for database readiness: %v", err)
				}

				continue
			}

			database, err = database.Clone(ctx, name)
			if err != nil {
				tb.Fatalf("[sqltest] failed to clone database: %v", err)
			}

			break
		}

		if debug.Load() {
			tb.Cleanup(func() {
				if !tb.Failed() || !debug.Swap(false) {
					return
				}

				tb.Logf("[sqltest] debug database %s", name)
				err := server.Database(name).Debug(ctx)
				if err != nil {
					tb.Logf("[sqltest] failed to debug database %s: %v", name, err)
				}
			})
		}

		return database
	}
}

// ---

type cloneFunc func(ctx context.Context, t testing.TB, base, name string) dbs.Database

// ---

var (
	passwordSalt = []byte("\x10\x00\xe2\x64\x72\xea\x4f\x50\xb5\xd9\xe2\x6a\x33\xbf\xe9\xc2")
)

// ---

var (
	_ = Test[*testing.T]{}
	_ = Test[*testing.B]{}
)
