package exec

import (
	"io"
	"os"
)

// Stub returns an implementation of [Exec] that does nothing.
func Stub() Exec {
	return stubExec{}
}

// StubCommand returns an implementation of [Command] that does nothing.
func StubCommand() Command {
	return &stubCommand{}
}

// StubProcess returns an implementation of [Process] that does nothing.
func StubProcess() Process {
	return &stubProcess{}
}

// StubProcessState returns an implementation of [ProcessState] that does nothing.
func StubProcessState() ProcessState {
	return &stubProcessState{}
}

// ---

type stubExec struct{}

func (stubExec) Command(name string, arg ...string) Command {
	return &stubCommand{}
}

// ---

type stubCommand struct{}

func (c *stubCommand) SetEnv([]string) {}

func (c *stubCommand) SetExtraEnv(...string) {}

func (c *stubCommand) SetStdout(io.Writer) {}

func (c *stubCommand) SetStderr(io.Writer) {}

func (c *stubCommand) Run() error {
	return nil
}

func (c *stubCommand) Start() error {
	return nil
}

func (c *stubCommand) Wait() error {
	return nil
}

func (c *stubCommand) Process() Process {
	return &stubProcess{}
}

func (c *stubCommand) String() string {
	return ""
}

// ---

type stubProcess struct{}

func (p *stubProcess) Signal(os.Signal) error {
	return nil
}

func (p *stubProcess) Wait() (ProcessState, error) {
	return &stubProcessState{}, nil
}

// ---

type stubProcessState struct{}

func (s *stubProcessState) ExitCode() int {
	return 0
}

func (s *stubProcessState) Exited() bool {
	return true
}
