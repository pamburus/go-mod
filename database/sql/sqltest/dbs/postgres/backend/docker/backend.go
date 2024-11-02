package docker

import (
	"bytes"
	"context"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"math/rand/v2"
	"net"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/backend"
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

	if options.Password == "" {
		bytes := fnv.New64().Sum(binary.AppendUvarint(nil, b.randSource.Uint64()))
		options.Password = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(bytes)
	}

	var suffix string
	if options.Port != 0 {
		suffix = fmt.Sprintf("%d", options.Port)
	} else {
		suffix = fmt.Sprintf("x%d", rand.Uint32()&0xffff)
	}
	container := fmt.Sprintf("sqltest-postgres-%s", suffix)

	cmd := exec.Command(
		"docker", "run",
		"--rm",
		"--name", container,
		"-e", "POSTGRES_PASSWORD",
		"-p", fmt.Sprintf("%d:5432", options.Port),
		b.image,
	)
	cmd.Env = append(cmd.Env, fmt.Sprintf("POSTGRES_PASSWORD=%s", options.Password))
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return fail(fmt.Errorf("failed to start postgres docker container: %w", err))
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

	fail0 := fail
	fail = func(err error) (backend.Server, backend.StopFunc, error) {
		stopErr := stop(ctx)
		if stopErr != nil {
			fmt.Fprintf(os.Stderr, "failed to stop postgres container: %v\n", stopErr)
		}

		return fail0(err)
	}

	if options.Port == 0 {
		var stdout bytes.Buffer

		cmd = exec.Command("docker", "inspect", "-f", `{{range (index .NetworkSettings.Ports "5432/tcp")}}{{.HostPort}}{{end}}`, container)
		cmd.Stdout = &stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			return fail(fmt.Errorf("failed to inspect postgres container: %w", err))
		}
		err = cmd.Wait()
		if err != nil {
			return fail(fmt.Errorf("failed to wait for inspect command: %w", err))
		}

		port, err := strconv.ParseUint(stdout.String(), 10, 16)
		if err != nil {
			return fail(fmt.Errorf("failed to parse port: %w", err))
		}

		options.Port = uint16(port)
	}

	location := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("postgres", options.Password),
		Host:   net.JoinHostPort("localhost", strconv.FormatUint(uint64(options.Port), 10)),
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
