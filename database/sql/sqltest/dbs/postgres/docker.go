package postgres

import (
	"math/rand/v2"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/instances/docker"
)

func Docker() DockerInstanceManagerBuilder {
	return DockerInstanceManagerBuilder{}
}

// ---

type DockerInstanceManagerBuilder struct {
	options []docker.Option
}

func (b DockerInstanceManagerBuilder) WithImage(image string) DockerInstanceManagerBuilder {
	b.options = append(b.options, docker.WithImage(image))

	return b
}

func (b DockerInstanceManagerBuilder) WithRandSource(randSource rand.Source) DockerInstanceManagerBuilder {
	b.options = append(b.options, docker.WithRandSource(randSource))

	return b
}

func (b DockerInstanceManagerBuilder) New() InstanceManager {
	return docker.New(b.options...)
}

func (b DockerInstanceManagerBuilder) Manager() InstanceManager {
	return b.New()
}

// ---

var (
	_ InstanceManagerProvider = DockerInstanceManagerBuilder{}
)
