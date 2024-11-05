package exec

import (
	"errors"
	"io"
	"os"
)

// MockStub returns an implementation of [Exec] that returns errors or panics on each method call.
func MockStub() Exec {
	return mockStubExec{}
}

// MockStubCommand returns an implementation of [Command] that returns errors or panics on each method call.
func MockStubCommand() Command {
	return &mockStubCommand{}
}

// MockStubProcess returns an implementation of [Process] that returns errors or panics on each method call.
func MockStubProcess() Process {
	return &mockStubProcess{}
}

// MockStubProcessState returns an implementation of [ProcessState] that returns errors or panics on each method call.
func MockStubProcessState() ProcessState {
	return &mockStubProcessState{}
}

// ---

type mockStubExec struct{}

func (mockStubExec) Command(name string, arg ...string) Command {
	panic(unimplemented)
}

// ---

type mockStubCommand struct{}

func (c *mockStubCommand) SetEnv([]string) {
	panic(unimplemented)
}

func (c *mockStubCommand) SetExtraEnv(...string) {
	panic(unimplemented)
}

func (c *mockStubCommand) SetStdout(io.Writer) {
	panic(unimplemented)
}

func (c *mockStubCommand) SetStderr(io.Writer) {
	panic(unimplemented)
}

func (c *mockStubCommand) Run() error {
	return unimplemented
}

func (c *mockStubCommand) Start() error {
	return unimplemented
}

func (c *mockStubCommand) Wait() error {
	return unimplemented
}

func (c *mockStubCommand) Process() Process {
	panic(unimplemented)
}

func (c *mockStubCommand) String() string {
	panic(unimplemented)
}

// ---

type mockStubProcess struct{}

func (p *mockStubProcess) Signal(os.Signal) error {
	return unimplemented
}

func (p *mockStubProcess) Wait() (ProcessState, error) {
	return nil, unimplemented
}

// ---

type mockStubProcessState struct{}

func (s *mockStubProcessState) ExitCode() int {
	panic(unimplemented)
}

func (s *mockStubProcessState) Exited() bool {
	panic(unimplemented)
}

// ---

var unimplemented = errors.New("not implemented")
