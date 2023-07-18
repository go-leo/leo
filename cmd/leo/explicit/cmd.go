package explicit

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

var (
	module           string
	appPath          string
	appPathDot       string
	appBaseName      string
	appUpperBaseName string
	serviceName      string
	sample           bool
	http             bool
	grpc             bool
	stream           bool
	schedule         bool
)

var Cmd = &cobra.Command{
	Use:   "explicit",
	Short: "Create a DDD project as explicit architecture.",
	Long:  "Create a DDD project as explicit architecture.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init a DDD project as explicit architecture",
	Long:  "init a DDD project as explicit architecture. Example: leo explicit init --module mall",
	Run: func(cmd *cobra.Command, args []string) {
		initProject()
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a service",
	Long:  "add a service. Example: leo explicit add --app user",
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(module) == 0 {
			goListCmd := exec.Command("go", "list", "-m")
			output, err := goListCmd.CombinedOutput()
			if err != nil {
				fmt.Println("go list -m: ", err)
				fmt.Printf("\n%s\n", string(output))
				os.Exit(1)
			}
			if len(output) <= 0 {
				fmt.Printf("module is empty")
				os.Exit(1)
			}
			module = strings.TrimSpace(string(output))
		}

		appBaseName = path.Base(appPath)
		appUpperBaseName = strings.ToUpper(appBaseName[:1]) + appBaseName[1:]
		appPathDot = strings.Replace(appPath, "/", ".", -1)
		serviceName = strings.Replace(module, "/", ".", -1) + "." + appPathDot
		if !sample && !http && !grpc && !stream && !schedule {
			sample = true
		}
		addService()
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&module, "module", "m", "", "model name, example: \"mall\",\"github.com/google/cloud\"")

	_ = initCmd.MarkPersistentFlagRequired("module")
	Cmd.AddCommand(initCmd)

	addCmd.PersistentFlags().StringVarP(&appPath, "app", "a", "", "application path, related to cmd dir, example: \"user\",\"order/svc\"")
	addCmd.PersistentFlags().BoolVarP(&sample, "sample", "", false, "sample application")
	addCmd.PersistentFlags().BoolVarP(&http, "http", "", false, "http application")
	addCmd.PersistentFlags().BoolVarP(&grpc, "grpc", "", false, "grpc application")
	addCmd.PersistentFlags().BoolVarP(&stream, "stream", "", false, "stream application")
	addCmd.PersistentFlags().BoolVarP(&schedule, "schedule", "", false, "schedule application")
	_ = addCmd.MarkPersistentFlagRequired("app")
	Cmd.AddCommand(addCmd)
}
