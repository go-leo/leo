package explicit

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
)

var moduleName string

var appPath string

var Cmd = &cobra.Command{
	Use:   "explicit",
	Short: "Create a DDD project as explicit architecture.",
	Long:  "Create a DDD project as explicit architecture.",
	Run: func(cmd *cobra.Command, args []string) {
		gen()
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init a DDD project as explicit architecture",
	Long:  "init a DDD project as explicit architecture. Example: leo explicit init --module mall",
	Run: func(cmd *cobra.Command, args []string) {
		if len(moduleName) == 0 {
			fmt.Println("module name is empty")
			flag.Usage()
			return
		}
		initProject()
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a service",
	Long:  "add a service. Example: leo explicit add --path user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(moduleName) == 0 {
			fmt.Println("module name is empty")
			flag.Usage()
			return
		}
		gen()
	},
}

func init() {
	initCmd.PersistentFlags().StringVarP(&moduleName, "module", "m", "", "model name (required)")
	_ = initCmd.MarkPersistentFlagRequired("module")
	Cmd.AddCommand(initCmd)

	addCmd.PersistentFlags().StringVarP(&appPath, "path", "", "", "app path")
	Cmd.AddCommand(addCmd)

}
