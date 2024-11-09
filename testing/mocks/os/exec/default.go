package exec

import (
	"io"
	"os"
	"os/exec"
)

// Default returns an implementation of [Exec] that delegates call
// to the standard library's [os] and [os/exec] packages.
func Default() Exec {
	return defaultExec{}
}

// ---

type defaultExec struct{}

func (defaultExec) Command(name string, arg ...string) Command {
	return &defaultCommand{exec.Command(name, arg...)}
}

// ---

type defaultCommand struct {
	cmd *exec.Cmd
}

func (c *defaultCommand) SetEnv(vars []string) {
	c.cmd.Env = vars
}

func (c *defaultCommand) SetExtraEnv(vars ...string) {
	if c.cmd.Env == nil {
		c.cmd.Env = os.Environ()
	}

	c.cmd.Env = append(c.cmd.Env, vars...)
}

func (c *defaultCommand) SetStdout(w io.Writer) {
	c.cmd.Stdout = w
}

func (c *defaultCommand) SetStderr(w io.Writer) {
	c.cmd.Stderr = w
}

func (c *defaultCommand) Run() error {
	return c.cmd.Run()
}

func (c *defaultCommand) Start() error {
	return c.cmd.Start()
}

func (c *defaultCommand) Wait() error {
	return c.cmd.Wait()
}

func (c *defaultCommand) Process() Process {
	return &defaultProcess{c.cmd.Process}
}

func (c *defaultCommand) String() string {
	return c.cmd.String()
}

// ---

type defaultProcess struct {
	p *os.Process
}

func (p *defaultProcess) Signal(sig os.Signal) error {
	return p.p.Signal(sig)
}

func (p *defaultProcess) Wait() (ProcessState, error) {
	state, err := p.p.Wait()
	if err != nil {
		return nil, err
	}

	return &defaultProcessState{state}, nil
}

// ---

type defaultProcessState struct {
	state *os.ProcessState
}

func (s *defaultProcessState) ExitCode() int {
	return s.state.ExitCode()
}

func (s *defaultProcessState) Exited() bool {
	return s.state.Exited()
}

func (s *defaultProcessState) String() string {
	return s.state.String()
}
