package docker_test

import (
	"context"
	"testing"
	"time"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres"
	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/backend"
)

func TestDocker(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srv, stop, err := postgres.Docker().New().Start(ctx, backend.Options{
		Password: "password",
	})
	if err != nil {
		t.Fatalf("failed to start postgres container: %v", err)
	}
	defer stop(ctx)

	location := srv.URL()
	if location == nil {
		t.Fatalf("got nil server url")
	}
	if location.Scheme != "postgres" {
		t.Fatalf("got scheme %q, want %q", location.Scheme, "postgres")
	}
	if location.Host != "localhost" {
		t.Fatalf("got host %q, want %q", location.Host, "localhost")
	}
	if location.User == nil {
		t.Fatalf("got nil user")
	}
}
