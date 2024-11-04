package docker

import (
	"bytes"
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
	"sync"
	"time"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/backend"
	"github.com/pamburus/go-mod/database/sql/sqltest/util/logging/logctx"
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

	var stderr bytes.Buffer
	container := fmt.Sprintf("sqltest-postgres-%d", port)
	logger := logctx.Get(ctx)

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

	logger.LogAttrs(ctx, slog.LevelDebug, "start postgres docker container", slog.Any("command", cmd.String()))
	err := cmd.Start()
	if err != nil {
		return fail(fmt.Errorf("failed to start postgres docker container: %w", err))
	}

	var wg sync.WaitGroup
	processContext, cancel := context.WithCancelCause(ctx)

	stop := func(ctx context.Context) error {
		defer wg.Wait()

		if processContext.Err() == nil {
			logger.LogAttrs(ctx, slog.LevelDebug, "stop postgres docker container", slog.String("container", container))
			err := cmd.Process.Signal(os.Interrupt)
			if err != nil {
				logger.LogAttrs(ctx, slog.LevelDebug, "failed to send interrupt signal", slog.Any("error", err))
			}
		}

		return nil
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		state, err := cmd.Process.Wait()
		if err == nil && state.ExitCode() != 0 {
			err = fmt.Errorf("postgres container exited with %s", state)
		}
		if err != nil {
			logger.LogAttrs(ctx, slog.LevelError, "docker container exited with error",
				slog.String("container", container),
				slog.Any("error", err),
			)
			for _, line := range strings.Split(stderr.String(), "\n") {
				logger.LogAttrs(ctx, slog.LevelError, "stderr", slog.String("line", line))
			}
		} else {
			logger.LogAttrs(ctx, slog.LevelDebug, "docker container exited", slog.String("container", container))
		}
		cancel(err)
	}()

	location := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("postgres", password),
		Host:   net.JoinHostPort("localhost", strconv.FormatUint(uint64(port), 10)),
		RawQuery: url.Values{
			"sslmode": []string{"disable"},
		}.Encode(),
	}

	return &server{processContext, location}, stop, nil
}

func (b *dockerBackend) Backend() backend.Backend {
	return b
}

// ---

type server struct {
	ctx      context.Context
	location *url.URL
}

func (s *server) Context() context.Context {
	return s.ctx
}

func (s *server) URL() *url.URL {
	return s.location
}
