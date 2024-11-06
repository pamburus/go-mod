package postgres

import (
	"context"
	"fmt"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs"
	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/instances"
)

func NewStarter(instances InstanceManagerProvider) dbs.Starter {
	return &starter{instances: instances.Manager()}
}

// ---

type starter struct {
	instances InstanceManager
	options   instances.Options
}

func (s *starter) Start(ctx context.Context) (dbs.Server, error) {
	instance, stop, err := s.instances.Start(ctx, s.options)
	if err != nil {
		return nil, fmt.Errorf("failed to start postgres container: %w", err)
	}

	return &server{
		s.options,
		instance,
		stop,
	}, nil
}
