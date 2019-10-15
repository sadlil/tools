package executor

import (
	"os"
	"os/exec"
)

type Command struct {
	executable string

	cmd *exec.Cmd
}

func NewCommand(cmd string) *Command {
	return &Command{
		executable: cmd,
		cmd:        exec.Command(cmd),
	}
}

func (c *Command) PipeSTD() *Command {
	c.cmd.Stdin = os.Stdin
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr

	return c
}

func (c *Command) Run(args ...string) error {
	c.cmd.Args = append([]string{c.executable}, args...)
	c.cmd.Process = nil
	return c.cmd.Run()
}
