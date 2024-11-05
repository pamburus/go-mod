package docker_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres"
	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/instances"
)

func TestManager(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	manager := postgres.Docker().New()
	assert.Equal(t, manager, manager.Manager())

	server, stop, err := postgres.Docker().New().Start(ctx, instances.Options{
		Password: "password",
	})
	if err != nil {
		t.Fatalf("failed to start postgres container: %v", err)
	}
	defer stop(ctx)

	location := server.URL()
	assert.NotNil(t, location)
	assert.Equal(t, location.Scheme, "postgres")
	assert.Equal(t, location.Hostname(), "localhost")
	assert.NotNil(t, location.User)
	assert.NotNil(t, server.Context())
	assert.NoError(t, server.Context().Err())
}
