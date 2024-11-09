package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs"
	"github.com/pamburus/go-mod/database/sql/sqltest/util/logging/logctx"
)

var _ dbs.Database = &database{}

// ---

type database struct {
	s   *server
	url *url.URL
}

func (d *database) URL() *url.URL {
	return d.url
}

func (d *database) Open(ctx context.Context) (*sql.DB, error) {
	return sql.Open("postgres", d.url.String())
}

func (d *database) Debug(ctx context.Context) error {
	name := d.name()

	host, port, err := net.SplitHostPort(d.url.Host)
	if err != nil {
		return fmt.Errorf("failed to split host and port: %w", err)
	}

	password, ok := d.url.User.Password()
	if !ok {
		return fmt.Errorf("missing password in url")
	}

	cmd := exec.Command("psql", "-U", "postgres", "-h", host, "-p", port, "-d", name)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+password)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	logctx.Get(ctx).LogAttrs(ctx, slog.LevelDebug, "run command", slog.Any("command", cmd.String()))
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run psql command: %w", err)
	}

	return nil
}

func (d *database) Name(ctx context.Context) (dbs.DatabaseName, error) {
	db, err := d.Open(ctx)
	if err != nil {
		return "", err
	}
	defer db.Close()

	var name string
	err = db.QueryRowContext(ctx, "SELECT current_database()").Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}

func (d *database) Clone(ctx context.Context, target dbs.DatabaseName) (dbs.Database, error) {
	fail := func(err error) (dbs.Database, error) {
		return nil, err
	}

	logctx.Get(ctx).LogAttrs(ctx, slog.LevelDebug, "clone database",
		slog.String("source", d.name()),
		slog.String("target", target),
	)

	db, err := d.Open(ctx)
	if err != nil {
		return fail(fmt.Errorf("failed to open database: %w", err))
	}
	defer db.Close()

	source := d.name()

	_, err = db.ExecContext(ctx, fmt.Sprintf("SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname='%s' AND pid <> pg_backend_pid();", source))
	if err != nil {
		return fail(fmt.Errorf("failed to terminate connections to database %s: %w", source, err))
	}

	_, err = db.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE %s WITH TEMPLATE %s", target, source))
	if err != nil {
		return fail(fmt.Errorf("failed to clone database %s to %s: %w", source, target, err))
	}

	return d.s.Database(target), nil
}

func (d *database) name() string {
	prefix := d.s.URL().Path
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	return strings.TrimPrefix(d.url.Path, prefix)
}
