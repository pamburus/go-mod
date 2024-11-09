package exec_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/testing/mocks/os/exec"
)

func TestDefault(t *testing.T) {
	exec := exec.Default()

	t.Run("Command", func(t *testing.T) {
		t.Run("Process", func(t *testing.T) {
			cmd := exec.Command("go", "version")

			err := cmd.Start()
			require.NoError(t, err)

			state, err := cmd.Process().Wait()
			require.NoError(t, err)
			require.Equal(t, 0, state.ExitCode())
			require.True(t, state.Exited())
		})

		t.Run("Wait", func(t *testing.T) {
			cmd := exec.Command("go", "version")

			err := cmd.Start()
			require.NoError(t, err)

			err = cmd.Wait()
			require.NoError(t, err)
		})

		t.Run("Run", func(t *testing.T) {
			cmd := exec.Command("go", "version")

			err := cmd.Run()
			require.NoError(t, err)
		})

		t.Run("String", func(t *testing.T) {
			cmd := exec.Command("go", "version")

			assert.Contains(t, cmd.String(), "go version")
		})

		t.Run("SetEnv", func(t *testing.T) {
			var stdout bytes.Buffer

			cmd := exec.Command("go", "env")
			cmd.SetStdout(&stdout)
			cmd.SetStderr(nil)
			cmd.SetEnv(append(os.Environ(), "GODEBUG=some-gibberish"))

			err := cmd.Run()
			require.NoError(t, err)
			assert.Contains(t, strings.Split(stdout.String(), "\n"), "GODEBUG='some-gibberish'")
		})

		t.Run("SetExtraEnv", func(t *testing.T) {
			var stdout bytes.Buffer

			cmd := exec.Command("go", "env")
			cmd.SetStdout(&stdout)
			cmd.SetStderr(nil)
			cmd.SetExtraEnv("CGO_ENABLED=1", "GODEBUG=some-gibberish")
			cmd.SetExtraEnv("CGO_ENABLED=0")

			err := cmd.Run()
			require.NoError(t, err)
			assert.Contains(t, strings.Split(stdout.String(), "\n"), "GODEBUG='some-gibberish'")
			assert.Contains(t, strings.Split(stdout.String(), "\n"), "CGO_ENABLED='0'")
		})

		t.Run("Signal", func(t *testing.T) {
			cmd := exec.Command("sleep", "1")

			err := cmd.Start()
			require.NoError(t, err)

			err = cmd.Process().Signal(os.Interrupt)
			require.NoError(t, err)

			state, err := cmd.Process().Wait()
			require.NoError(t, err)
			require.NotEqual(t, 0, state.ExitCode())

			state, err = cmd.Process().Wait()
			require.Error(t, err)
		})
	})
}
