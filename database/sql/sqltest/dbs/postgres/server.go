package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs"
)

func NewStarter(backend Backend) dbs.Starter {
	return starter{backend}
}

// ---

type starter struct {
	backend Backend
}

func (s starter) Start(ctx context.Context, port int, password string) (dbs.Server, error) {
	stop, err := s.backend.Start(ctx, port, password)
	if err != nil {
		return nil, fmt.Errorf("failed to start postgres container: %w", err)
	}

	return &server{
		stop,
		port,
		password,
	}, nil
}

// ---

type server struct {
	stop     func(context.Context) error
	port     int
	password string
}

func (s *server) Database(name dbs.DatabaseName) dbs.Database {
	if name == "" {
		name = "postgres"
	}

	return &database{s, &url.URL{
		Scheme: "postgres",
		Host:   net.JoinHostPort("localhost", strconv.Itoa(s.port)),
		User:   url.UserPassword("postgres", s.password),
		Path:   path.Join("/", name),
		RawQuery: url.Values{
			"sslmode": []string{"disable"},
		}.Encode(),
	}}
}

func (s *server) Stop(ctx context.Context) error {
	return s.stop(ctx)
}

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

	cmd := exec.Command("psql", "-U", "postgres", "-h", "localhost", "-p", strconv.Itoa(d.s.port), "-d", name)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+d.s.password)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
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
	return strings.TrimPrefix(d.url.Path, "/")
}
