package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs"
	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/backend"
	"github.com/pamburus/go-mod/database/sql/sqltest/util/logging"
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
	srv, stop, err := s.backend.Start(ctx, s.options)
	if err != nil {
		return nil, fmt.Errorf("failed to start postgres container: %w", err)
	}

	return &server{
		srv,
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

func (s starter) WithLogger(logger *slog.Logger) dbs.Starter {
	s.options.Logger = logger

	return &s
}

func (s starter) WithRawLogger(logger logging.RawLogger, level slog.Level) dbs.Starter {
	s.options.Logger = logging.RawToStructured(logger, level)

	return &s
}
