package dbs

import (
	"context"
)

type Starter interface {
	Start(ctx context.Context, port int, password string) (Server, error)
}

type Server interface {
	Database(DatabaseName) Database
	Stop(context.Context) error
}
