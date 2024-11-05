package exec_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/database/sql/sqltest/util/abstract/os/exec"
)

func TestStub(t *testing.T) {
	testProcessState := func(state exec.ProcessState) func(*testing.T) {
		return func(t *testing.T) {
			t.Run("ExitCode", func(t *testing.T) {
				assert.Equal(t, 0, state.ExitCode())
			})

			t.Run("Exited", func(t *testing.T) {
				assert.True(t, state.Exited())
			})
		}
	}

	testProcess := func(process exec.Process) func(*testing.T) {
		return func(t *testing.T) {
			t.Run("Signal", func(t *testing.T) {
				err := process.Signal(os.Interrupt)
				require.NoError(t, err)
			})

			t.Run("Wait", func(t *testing.T) {
				state, err := process.Wait()
				require.NoError(t, err)
				assert.NotNil(t, state)
				t.Run("ProcessState", testProcessState(state))
			})
		}
	}

	testCommand := func(command exec.Command) func(*testing.T) {
		return func(t *testing.T) {
			t.Run("Start", func(t *testing.T) {
				err := command.Start()
				require.NoError(t, err)
			})

			t.Run("Wait", func(t *testing.T) {
				err := command.Start()
				require.NoError(t, err)

				err = command.Wait()
				require.NoError(t, err)
			})

			t.Run("Run", func(t *testing.T) {
				err := command.Run()
				require.NoError(t, err)
			})

			t.Run("String", func(t *testing.T) {
				assert.Empty(t, command.String())
			})

			t.Run("SetEnv", func(t *testing.T) {
				assert.NotPanics(t, func() {
					command.SetEnv([]string{"GODEBUG=some-gibberish"})
				})
			})

			t.Run("SetExtraEnv", func(t *testing.T) {
				assert.NotPanics(t, func() {
					command.SetExtraEnv("GODEBUG=some-gibberish")
				})
			})

			t.Run("SetStdout", func(t *testing.T) {
				assert.NotPanics(t, func() {
					command.SetStdout(nil)
				})
			})

			t.Run("SetStderr", func(t *testing.T) {
				assert.NotPanics(t, func() {
					command.SetStderr(nil)
				})
			})

			t.Run("Process", testProcess(command.Process()))
		}
	}

	t.Run("Exec", func(t *testing.T) {
		exec := exec.Stub()

		t.Run("Command", testCommand(exec.Command("test")))
	})

	t.Run("Command", testCommand(exec.StubCommand()))
	t.Run("Process", testProcess(exec.StubProcess()))
	t.Run("ProcessState", testProcessState(exec.StubProcessState()))
}
