package postgres

import (
	"context"
	"fmt"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs"
	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/backend"
)

func NewStarter(backend BackendProvider) dbs.Starter {
	return &starter{backend: backend.Backend()}
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
