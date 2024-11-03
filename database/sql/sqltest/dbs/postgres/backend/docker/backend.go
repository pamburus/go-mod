package docker

import (
	"bytes"
	"cmp"
	"context"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"net"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/backend"
	"github.com/pamburus/go-mod/database/sql/sqltest/util/logging"
	"github.com/pamburus/go-mod/database/sql/sqltest/util/portalloc"
	"github.com/pamburus/go-mod/database/sql/sqltest/util/random"
)

// ---

func New(options ...Option) backend.Backend {
	opts := jointOptions{
		image: "postgres:alpine",
	}
	for _, o := range options {
		o(&opts)
	}
	if opts.randSource == nil {
		opts.randSource = rand.NewPCG(0, uint64(time.Now().UnixNano()))
	}

	return &dockerBackend{opts}
}

func WithImage(image string) Option {
	return func(o *jointOptions) {
		o.image = image
	}
}

func WithRandSource(randSource rand.Source) Option {
	return func(o *jointOptions) {
		o.randSource = randSource
	}
}

type Option func(*jointOptions)

// ---

type jointOptions struct {
	image      string
	randSource rand.Source
}

// ---

type dockerBackend struct {
	jointOptions
}

func (b *dockerBackend) Start(ctx context.Context, options backend.Options) (backend.Server, backend.StopFunc, error) {
	fail := func(err error) (backend.Server, backend.StopFunc, error) {
		return nil, nil, err
	}

	port, password := options.Port, options.Password
	if password == "" {
		password = random.Password(b.randSource)
	}
	if port == 0 {
		var err error
		port, err = portalloc.New()
		if err != nil {
			return fail(fmt.Errorf("failed to allocate port: %w", err))
		}
	}
	logger := cmp.Or(options.Logger, logging.DiscardLogger())

	var stderr bytes.Buffer
	container := fmt.Sprintf("sqltest-postgres-%d", port)

	logger.InfoContext(ctx, "start postgres docker container", slog.String("container", container))
	cmd := exec.Command(
		"docker", "run",
		"--rm",
		"--name", container,
		"-e", "POSTGRES_PASSWORD",
		"-p", fmt.Sprintf("%d:5432", port),
		b.image,
	)
	cmd.Env = append(cmd.Env, fmt.Sprintf("POSTGRES_PASSWORD=%s", password))
	cmd.Stderr = &stderr

	logger.DebugContext(ctx, "command", slog.Any("command", cmd.String()))
	err := cmd.Start()
	if err != nil {
		return fail(fmt.Errorf("failed to start postgres docker container: %w", err))
	}

	stop := func(ctx context.Context) error {
		logger.InfoContext(ctx, "stop postgres docker container", slog.String("container", container))
		err := cmd.Process.Signal(os.Interrupt)
		if err != nil {
			return fmt.Errorf("failed to send interrupt signal: %w", err)
		}

		_, err = cmd.Process.Wait()
		if err != nil {
			for _, line := range strings.Split(stderr.String(), "\n") {
				logger.Error(line)
			}

			return fmt.Errorf("failed to wait for postgres container: %w", err)
		}

		return nil
	}

	location := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("postgres", password),
		Host:   net.JoinHostPort("localhost", strconv.FormatUint(uint64(port), 10)),
		RawQuery: url.Values{
			"sslmode": []string{"disable"},
		}.Encode(),
	}

	return &server{location}, stop, nil
}

// ---

type server struct {
	location *url.URL
}

func (s *server) URL() *url.URL {
	return s.location
}
