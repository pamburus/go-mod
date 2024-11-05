package postgres

import (
	"cmp"
	"context"
	"net/url"
	"path"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs"
	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/instances"
)

var _ dbs.Server = &server{}

// ---

type server struct {
	opts     instances.Options
	instance instances.Instance
	stop     func(context.Context) error
}

func (s *server) Context() context.Context {
	return s.instance.Context()
}

func (s *server) URL() *url.URL {
	return s.instance.URL()
}

func (s *server) Database(name dbs.DatabaseName) dbs.Database {
	if name == "" {
		name = "postgres"
	}

	u := *s.instance.URL()
	u.Path = path.Join(cmp.Or(u.Path, "/"), name)

	return &database{s, &u}
}

func (s *server) Stop(ctx context.Context) error {
	return s.stop(ctx)
}
