package docker_test

import (
	"context"
	"testing"
	"time"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres"
	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/instances"
)

func TestManager(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server, stop, err := postgres.Docker().New().Start(ctx, instances.Options{
		Password: "password",
	})
	if err != nil {
		t.Fatalf("failed to start postgres container: %v", err)
	}
	defer stop(ctx)

	location := server.URL()
	if location == nil {
		t.Fatalf("got nil server url")
	}
	if location.Scheme != "postgres" {
		t.Fatalf("got scheme %q, want %q", location.Scheme, "postgres")
	}
	if location.Hostname() != "localhost" {
		t.Fatalf("got host %q, want %q", location.Hostname(), "localhost")
	}
	if location.User == nil {
		t.Fatalf("got nil user")
	}
}
