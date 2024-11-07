package docker_test

import (
	"context"
	"errors"
	"math/rand/v2"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/instances"
	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/instances/docker"
	"github.com/pamburus/go-mod/database/sql/sqltest/dbs/postgres/instances/docker/mocks"
)

func TestManager(t *testing.T) {
	ctx := context.Background()

	begin := func() (func(), func(mock.Arguments)) {
		ch := make(chan struct{})
		done := func() {
			close(ch)
		}
		wait := func(mock.Arguments) {
			<-ch
		}

		return done, wait
	}

	t.Run("PortAllocFailure", func(t *testing.T) {
		errAllocationFailure := errors.New("port allocation failed")
		manager := docker.New(
			docker.WithExec(mocks.NewExec(t)),
			docker.WithPortAllocator(func() (uint16, error) {
				return 8800, errAllocationFailure
			}),
		)

		_, _, err := manager.Start(ctx, instances.Options{})
		assert.ErrorIs(t, err, errAllocationFailure)
	})

	t.Run("Success", func(t *testing.T) {
		exec := mocks.NewExec(t)
		command := mocks.NewCommand(t)

		manager := docker.New(
			docker.WithExec(exec),
			docker.WithImage("postgres:latest"),
			docker.WithPortAllocator(func() (uint16, error) {
				return 8800, nil
			}),
			docker.WithRandSource(rand.NewPCG(0, 0)),
		)
		assert.Equal(t, manager, manager.Manager())

		process := mocks.NewProcess(t)
		processState := mocks.NewProcessState(t)

		command.On("SetExtraEnv", []string{"POSTGRES_PASSWORD=aiqjsgw3pccqa"}).Once().Return()
		command.On("SetStderr", mock.Anything).Return()
		exec.On("Command", "docker", []string{
			"run",
			"--rm",
			"--name",
			"sqltest-postgres-8800",
			"-e",
			"POSTGRES_PASSWORD",
			"-p",
			"8800:5432",
			"postgres:latest",
		}).Return(command)

		errStartFailure := errors.New("start failure")
		command.On("Start").Once().Return(errStartFailure)
		_, _, err := manager.Start(ctx, instances.Options{})
		assert.ErrorIs(t, err, errStartFailure)

		command.On("Start").Return(nil)
		command.On("SetExtraEnv", []string{"POSTGRES_PASSWORD=7mkifacgjntgq"}).Once().Return()
		command.On("Process").Return(process)

		processDone, processWait := begin()
		process.On("Wait").Return(processState, nil).Run(processWait)

		server, stop, err := manager.Start(ctx, instances.Options{})
		if err != nil {
			t.Fatalf("failed to start postgres container: %v", err)
		}

		errSignalFailure := errors.New("signal failure")
		process.On("Signal", os.Interrupt).Once().Return(errSignalFailure)
		err = stop(ctx)
		assert.ErrorIs(t, err, errSignalFailure)

		location := server.URL()
		assert.NotNil(t, location)
		assert.Equal(t, location.Scheme, "postgres")
		assert.Equal(t, location.Hostname(), "localhost")
		assert.NotNil(t, location.User)
		assert.NotNil(t, server.Context())
		assert.NoError(t, server.Context().Err())

		process.On("Signal", os.Interrupt).Return(nil)
		processState.On("ExitCode").Return(0)

		processDone()
		stop(ctx)

		assert.ErrorIs(t, context.Cause(server.Context()), context.Canceled)
	})

	t.Run("Failure", func(t *testing.T) {
		exec := mocks.NewExec(t)
		command := mocks.NewCommand(t)

		manager := docker.New(
			docker.WithExec(exec),
			docker.WithImage("postgres:latest"),
			docker.WithPortAllocator(func() (uint16, error) {
				return 8800, nil
			}),
			docker.WithRandSource(rand.NewPCG(0, 0)),
		)
		assert.Equal(t, manager, manager.Manager())

		process := mocks.NewProcess(t)
		processState := mocks.NewProcessState(t)

		command.On("SetExtraEnv", []string{"POSTGRES_PASSWORD=aiqjsgw3pccqa"}).Once().Return()
		command.On("SetStderr", mock.Anything).Return()
		command.On("Start").Return(nil)
		command.On("Process").Return(process)
		exec.On("Command", "docker", []string{
			"run",
			"--rm",
			"--name",
			"sqltest-postgres-8800",
			"-e",
			"POSTGRES_PASSWORD",
			"-p",
			"8800:5432",
			"postgres:latest",
		}).Return(command)

		process.On("Wait").Return(processState, nil)

		server, stop, err := manager.Start(ctx, instances.Options{})
		if err != nil {
			t.Fatalf("failed to start postgres container: %v", err)
		}

		location := server.URL()
		assert.NotNil(t, location)
		assert.Equal(t, location.Scheme, "postgres")
		assert.Equal(t, location.Hostname(), "localhost")
		assert.NotNil(t, location.User)
		assert.NotNil(t, server.Context())
		assert.NoError(t, server.Context().Err())

		process.On("Signal", os.Interrupt).Return(nil)
		processState.On("ExitCode").Return(1)
		processState.On("String").Return("exit code 1")

		stop(ctx)

		assert.Error(t, context.Cause(server.Context()), "postgres container exited with exit code 1")
	})
}
