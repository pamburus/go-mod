package postgres

import (
	"math/rand/v2"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/backend/docker"
)

func Docker() DockerBackendBuilder {
	return DockerBackendBuilder{}
}

// ---

type DockerBackendBuilder struct {
	options []docker.Option
}

func (b DockerBackendBuilder) WithImage(image string) DockerBackendBuilder {
	b.options = append(b.options, docker.WithImage(image))

	return b
}

func (b DockerBackendBuilder) WithRandSource(randSource rand.Source) DockerBackendBuilder {
	b.options = append(b.options, docker.WithRandSource(randSource))

	return b
}

func (b DockerBackendBuilder) New() Backend {
	return docker.New(b.options...)
}

func (b DockerBackendBuilder) Backend() Backend {
	return b.New()
}

// ---

var (
	_ BackendProvider = DockerBackendBuilder{}
)
