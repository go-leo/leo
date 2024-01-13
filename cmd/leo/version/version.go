package version

import (
	"codeup.aliyun.com/qimao/leo/leo"
	"fmt"
	"github.com/spf13/cobra"
)

// Cmd represents the version command.
var Cmd = &cobra.Command{
	Use:   "version",
	Short: "print the version and exit",
	Long:  "print the version and exit. Example: leo version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("leo %v\n", leo.Version)
	},
}
