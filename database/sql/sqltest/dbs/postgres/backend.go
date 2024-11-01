package postgres

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

type Backend interface {
	Start(ctx context.Context, port int, password string) (func(context.Context) error, error)
}

// ---

func Docker(image string) Backend {
	return dockerBackend{image}
}

// ---

type dockerBackend struct {
	image string
}

func (b dockerBackend) Start(ctx context.Context, port int, password string) (func(context.Context) error, error) {
	cmd := exec.Command(
		"docker", "run",
		"--rm",
		"--name", fmt.Sprintf("sqltest-postgres-%d", port),
		"-e", "POSTGRES_PASSWORD",
		"-p", fmt.Sprintf("%d:5432", port),
		b.image,
	)
	cmd.Env = append(cmd.Env, fmt.Sprintf("POSTGRES_PASSWORD=%s", password))
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	stop := func(ctx context.Context) error {
		err := cmd.Process.Signal(os.Interrupt)
		if err != nil {
			return fmt.Errorf("failed to send interrupt signal: %w", err)
		}

		_, err = cmd.Process.Wait()
		if err != nil {
			return fmt.Errorf("failed to wait for postgres container: %w", err)
		}

		return nil
	}

	return stop, nil
}
