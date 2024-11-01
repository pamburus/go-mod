package dbs

import (
	"context"
	"database/sql"
	"net/url"
)

type Database interface {
	URL() *url.URL
	Open(context.Context) (*sql.DB, error)
	Name(context.Context) (DatabaseName, error)
	Debug(context.Context) error
	Clone(context.Context, DatabaseName) (Database, error)
}

type DatabaseName = string
