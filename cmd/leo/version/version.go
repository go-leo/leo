package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Version = "v0.0.20"

// Cmd represents the version command.
var Cmd = &cobra.Command{
	Use:   "version",
	Short: "print the version and exit",
	Long:  "print the version and exit. Example: leo version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("leo %v\n", Version)
	},
}
