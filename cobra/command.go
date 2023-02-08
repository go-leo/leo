package cobra

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
)

type Command struct {
	rootCmd *cobra.Command
}

func New(cmds ...*cobra.Command) *Command {
	var cmd *cobra.Command
	for i := len(cmds) - 1; i >= 0; i-- {
		cur := cmds[i]
		if cmd == nil {
			cmd = cur
			continue
		}
		cur.AddCommand(cmd)
		cmd = cur
	}
	return &Command{rootCmd: cmd}
}

func (c *Command) String() string {
	return "cobra command"
}

func (c *Command) Invoke(ctx context.Context) error {
	if c.rootCmd == nil {
		return errors.New("command is nil")
	}
	return c.rootCmd.ExecuteContext(ctx)
}
