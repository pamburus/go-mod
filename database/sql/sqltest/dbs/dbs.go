package dbs

import (
	"context"
	"database/sql"
	"net/url"
)

type Starter interface {
	Start(ctx context.Context) (Server, error)
	WithPassword(string) Starter
	WithPort(uint16) Starter
}

type Server interface {
	Context() context.Context
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
