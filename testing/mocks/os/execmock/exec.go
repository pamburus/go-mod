package execmock

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/pamburus/go-mod/testing/mocks/os/exec"
)

func NewExec(t testing.TB) *Exec {
	return withTest(t, &Exec{})
}

func NewCommand(t testing.TB) *Command {
	return withTest(t, &Command{})
}

func NewProcess(t testing.TB) *Process {
	return withTest(t, &Process{})
}

func NewProcessState(t testing.TB) *ProcessState {
	return withTest(t, &ProcessState{})
}

// ---

type Exec struct {
	mock.Mock
}

func (o *Exec) Command(name string, args ...string) exec.Command {
	return o.Called(name, args).Get(0).(exec.Command)
}

// ---

type Command struct {
	mock.Mock
}

func (m *Command) Start() error {
	return m.Called().Error(0)
}

func (m *Command) Wait() error {
	return m.Called().Error(0)
}

func (m *Command) Run() error {
	return m.Called().Error(0)
}

func (m *Command) String() string {
	return m.Called().String(0)
}

func (m *Command) SetEnv(env []string) {
	m.Called(env)
}

func (m *Command) SetExtraEnv(env ...string) {
	m.Called(env)
}

func (m *Command) SetStdout(stdout io.Writer) {
	m.Called(stdout)
}

func (m *Command) SetStderr(stderr io.Writer) {
	m.Called(stderr)
}

func (m *Command) Process() exec.Process {
	return m.Called().Get(0).(exec.Process)
}

// ---

type Process struct {
	mock.Mock
}

func (m *Process) Signal(sig os.Signal) error {
	return m.Called(sig).Error(0)
}

func (m *Process) Wait() (exec.ProcessState, error) {
	ret := m.Called()

	return ret.Get(0).(exec.ProcessState), ret.Error(1)
}

// ---

type ProcessState struct {
	mock.Mock
}

func (m *ProcessState) ExitCode() int {
	return m.Called().Int(0)
}

func (m *ProcessState) Exited() bool {
	return m.Called().Bool(0)
}

func (m *ProcessState) String() string {
	return m.Called().String(0)
}

// ---

func withTest[O interface{ Test(mock.TestingT) }](t testing.TB, obj O) O {
	obj.Test(t)

	return obj
}

// ---

var (
	_ exec.Exec         = (*Exec)(nil)
	_ exec.Command      = (*Command)(nil)
	_ exec.Process      = (*Process)(nil)
	_ exec.ProcessState = (*ProcessState)(nil)
)
