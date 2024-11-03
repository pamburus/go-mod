package postgres

import (
	"context"
	"fmt"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs"
	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/backend"
)

func NewStarter(backend Backend) dbs.Starter {
	return &starter{backend: backend}
}

// ---

type starter struct {
	backend Backend
	options backend.Options
}

func (s *starter) Start(ctx context.Context) (dbs.Server, error) {
	bs, stop, err := s.backend.Start(ctx, s.options)
	if err != nil {
		return nil, fmt.Errorf("failed to start postgres container: %w", err)
	}

	return &server{
		s.options,
		bs,
		stop,
	}, nil
}

func (s starter) WithPassword(password string) dbs.Starter {
	s.options.Password = password

	return &s
}

func (s starter) WithPort(port uint16) dbs.Starter {
	s.options.Port = port

	return &s
}
