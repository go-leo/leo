package cobra

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
)

type Command struct {
	rootCmd *cobra.Command
}

func New(rootCmd *cobra.Command) *Command {
	return &Command{rootCmd: rootCmd}
}

func (c *Command) String() string {
	return "cobra command"
}

func (c *Command) Invoke(ctx context.Context) error {
	if c.rootCmd == nil {
		return errors.New("root command is nil")
	}
	return c.rootCmd.ExecuteContext(ctx)
}
