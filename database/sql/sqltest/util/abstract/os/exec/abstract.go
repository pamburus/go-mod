// Package exec provides an abstraction over os/exec package.
//
// Note that there are no strong backward compatibility guarantees for this package.
// New methods may be added in the future to any interface in this package.
// If you are using this package to automatically generate mocks,
// be prepared to update your code if necessary.
// If you are writing mocks by hand, you may want to embed corresponding stubs
// to you implementations for each of the interface provided by this package
// to ensure build is not broken after update.
package exec

import (
	"io"
	"os"
)

// Exec is an abstraction over [os/exec](https://pkg.go.dev/os/exec) package.
type Exec interface {
	// Command creates a new command.
	Command(name string, arg ...string) Command
}

// Command is an abstraction over [os/exec.Cmd](https://pkg.go.dev/os/exec#Cmd) type.
type Command interface {
	// SetEnv sets the environment variables to the command.
	// Each entry is of the form "key=value".
	// If env is nil, the new process uses the current process's environment.
	// If env contains duplicate environment keys, only the last value in the slice for each duplicate key is used.
	// By default env is nil.
	SetEnv(env []string)
	// SetExtraEnv appends the environment variables to the command.
	// If env is not initialized yet, it is initialized with the current process's environment.
	SetExtraEnv(env ...string)
	// SetStdout sets target for the command's stdout.
	// If set to nil, the stdout file descriptor is connected to the null device (os.DevNull).
	// If set to an *os.File, the stdout from the process is connected directly to that file.
	// By default stdout target is nil.
	SetStdout(io.Writer)
	// SetStderr sets target for the command's stderr.
	// If set to nil, the stderr file descriptor is connected to the null device (os.DevNull).
	// If set to an *os.File, the stderr from the process is connected directly to that file.
	// By default stderr target is nil.
	SetStderr(io.Writer)
	// Run executes the command and waits for it to complete.
	Run() error
	// Start starts the command but does not wait for it to complete.
	Start() error
	// Wait waits for the command to complete.
	Wait() error
	// Process returns the underlying process.
	Process() Process
	// String returns the command as a string.
	String() string
}

// Process is an abstraction over [os.Process] type.
type Process interface {
	// Signal sends a signal to the process.
	Signal(os.Signal) error
	// Wait waits for the process to complete.
	Wait() (ProcessState, error)
}

// ProcessState is an abstraction over [os.ProcessState] type.
type ProcessState interface {
	// ExitCode returns the exit code of the process.
	ExitCode() int
	// Exited returns true if the process has exited.
	Exited() bool
}
