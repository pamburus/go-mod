package sqltest

import (
	"cmp"
	"context"
	"database/sql"
	"hash/fnv"
	"log/slog"
	"net/url"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs"
	"github.com/pamburus/go-mod/database/sql/sqltest/util/envflag"
	"github.com/pamburus/go-mod/database/sql/sqltest/util/hashstr"
	"github.com/pamburus/go-mod/database/sql/sqltest/util/logging"
	"github.com/pamburus/go-mod/database/sql/sqltest/util/logging/logctx"
	"github.com/pamburus/go-mod/database/sql/sqltest/util/sleep"
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

	dbName := "init"
	ctx, logger, clone := setup(ctx, b.tt, b.starter)
	database := clone(ctx, b.tt, "", dbName)

	return &Test[T]{
		b.tt,
		b.tt,
		ctx,
		b.ctxFunc,
		logger,
		database,
		sync.Once{},
		nil,
		dbName,
		clone,
	}
}

// ---

type Test[T Base[T]] struct {
	testing.TB
	base     T
	ctx      context.Context
	ctxFunc  func(T) context.Context
	logger   loggerFunc
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
	ctx = logctx.New(ctx, t.logger(tt))

	dbName := "test_" + hashstr.New(fnv.New64(), []byte(tt.Name()))
	database := t.clone(ctx, tt, t.dbName, dbName)

	return &Test[T]{
		tt,
		tt,
		ctx,
		t.ctxFunc,
		t.logger,
		database,
		sync.Once{},
		nil,
		dbName,
		t.clone,
	}
}

// ---

func setup(ctx context.Context, t testing.TB, starter dbs.Starter) (context.Context, loggerFunc, cloneFunc) {
	var debug atomic.Bool
	debug.Store(envflag.Get(envVarDebug))

	level := slog.LevelInfo
	if v, ok := os.LookupEnv(envVarLogLevel); ok {
		err := level.UnmarshalText([]byte(v))
		if err != nil {
			t.Fatalf("failed to parse %s: %v", envVarLogLevel, err)
		}
	}

	newLogger := func(t testing.TB) *slog.Logger {
		return logging.RawToStructured(logging.RawPrefix("[sqltest]", t), level)
	}

	logger := newLogger(t)
	ctx = logctx.New(ctx, logger)

	server, err := starter.Start(ctx)
	if err != nil {
		logger.LogAttrs(ctx, slog.LevelError, "failed to start database server", slog.Any("error", err))
		t.FailNow()
	}

	t.Cleanup(func() {
		err := server.Stop(ctx)
		if err != nil {
			logger.LogAttrs(ctx, slog.LevelError, "failed to stop database server",
				slog.Any("location", server.URL()),
				slog.Any("error", err),
			)
			t.FailNow()
		}
	})

	clone := func(ctx context.Context, t testing.TB, base string, name string) dbs.Database {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		database := server.Database(base)

		for {
			base, err = database.Name(ctx)
			if err != nil {
				logger.DebugContext(ctx, "database is not ready", slog.Any("reason", err))
				err := sleep.Sleep(ctx, time.Second/4)
				if err != nil {
					logger.LogAttrs(ctx, slog.LevelError, "failed to wait for database readiness",
						slog.Any("database", database.URL()),
						slog.Any("error", err),
					)
					t.FailNow()
				}

				continue
			}

			database, err = database.Clone(ctx, name)
			if err != nil {
				logger.LogAttrs(ctx, slog.LevelError, "failed to clone database",
					slog.Any("database", database.URL()),
					slog.String("target", name),
					slog.Any("error", err),
				)
				t.FailNow()
			}

			break
		}

		if debug.Load() {
			t.Cleanup(func() {
				if !t.Failed() || !debug.Swap(false) {
					return
				}

				logger.InfoContext(ctx, "debug database", slog.Any("database", database.URL()))
				err := database.Debug(ctx)
				if err != nil {
					logger.LogAttrs(ctx, slog.LevelError, "failed to debug database",
						slog.Any("database", database.URL()),
						slog.Any("error", err),
					)
					t.FailNow()
				}
			})
		}

		return database
	}

	return server.Context(), newLogger, clone
}

// ---

type (
	cloneFunc  func(ctx context.Context, t testing.TB, base, name string) dbs.Database
	loggerFunc func(t testing.TB) *slog.Logger
)

// ---

var (
	_ = Test[*testing.T]{}
	_ = Test[*testing.B]{}
)
