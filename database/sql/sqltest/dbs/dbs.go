package dbs

import (
	"context"
	"database/sql"
	"log/slog"
	"net/url"

	"github.com/pamburus/go-mod/database/sql/sqltest/util/logging"
)

type Starter interface {
	Start(ctx context.Context) (Server, error)
	WithPassword(string) Starter
	WithPort(uint16) Starter
	WithLogger(*slog.Logger) Starter
	WithRawLogger(logging.RawLogger, slog.Level) Starter
}

type Server interface {
	URL() *url.URL
	Database(DatabaseName) Database
	Stop(context.Context) error
}

type Database interface {
	URL() *url.URL
	Open(context.Context) (*sql.DB, error)
	Name(context.Context) (DatabaseName, error)
	Debug(context.Context) error
	Clone(context.Context, DatabaseName) (Database, error)
}

type DatabaseName = string
