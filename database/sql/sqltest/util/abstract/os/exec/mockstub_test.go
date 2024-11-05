package exec_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/database/sql/sqltest/util/abstract/os/exec"
)

func TestMockStub(t *testing.T) {
	t.Run("Exec", func(t *testing.T) {
		exec := exec.MockStub()

		t.Run("Command", func(t *testing.T) {
			assert.Panics(t, func() {
				exec.Command("go", "version")
			})
		})
	})

	t.Run("Command", func(t *testing.T) {
		cmd := exec.MockStubCommand()

		t.Run("Start", func(t *testing.T) {
			err := cmd.Start()
			require.Error(t, err)
		})

		t.Run("Wait", func(t *testing.T) {
			err := cmd.Wait()
			require.Error(t, err)
		})

		t.Run("Run", func(t *testing.T) {
			err := cmd.Run()
			require.Error(t, err)
		})

		t.Run("String", func(t *testing.T) {
			assert.Panics(t, func() {
				_ = cmd.String()
			})
		})

		t.Run("SetEnv", func(t *testing.T) {
			assert.Panics(t, func() {
				cmd.SetEnv([]string{"GODEBUG=some-gibberish"})
			})
		})

		t.Run("SetExtraEnv", func(t *testing.T) {
			assert.Panics(t, func() {
				cmd.SetExtraEnv("GODEBUG=some-gibberish")
			})
		})

		t.Run("SetStdout", func(t *testing.T) {
			assert.Panics(t, func() {
				cmd.SetStdout(nil)
			})
		})

		t.Run("SetStderr", func(t *testing.T) {
			assert.Panics(t, func() {
				cmd.SetStderr(nil)
			})
		})

		t.Run("Process", func(t *testing.T) {
			assert.Panics(t, func() {
				_ = cmd.Process()
			})
		})
	})

	t.Run("Process", func(t *testing.T) {
		process := exec.MockStubProcess()

		t.Run("Signal", func(t *testing.T) {
			err := process.Signal(os.Interrupt)
			require.Error(t, err)
		})

		t.Run("Wait", func(t *testing.T) {
			state, err := process.Wait()
			require.Error(t, err)
			assert.Nil(t, state)
		})
	})

	t.Run("ProcessState", func(t *testing.T) {
		state := exec.MockStubProcessState()

		t.Run("ExitCode", func(t *testing.T) {
			assert.Panics(t, func() {
				_ = state.ExitCode()
			})
		})

		t.Run("Exited", func(t *testing.T) {
			assert.Panics(t, func() {
				_ = state.Exited()
			})
		})
	})
}
