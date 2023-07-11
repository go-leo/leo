package explicit

import (
	"github.com/spf13/cobra"
	"path"
	"strings"
)

var (
	module           string
	app              string
	appBaseName      string
	appUpperBaseName string
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
	Long:  "add a service. Example: leo explicit add --path user",
	Run: func(cmd *cobra.Command, args []string) {
		appBaseName = path.Base(app)
		appUpperBaseName = strings.ToUpper(appBaseName[:1]) + appBaseName[1:]

		if !sample && !http && !grpc && !stream && !schedule {
			sample = true
		}

		addService()
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&module, "module", "m", "", "model name (required)")
	Cmd.PersistentFlags().StringVarP(&app, "app", "a", "", "app name")
	Cmd.PersistentFlags().BoolVarP(&sample, "sample", "", false, "sample application")
	Cmd.PersistentFlags().BoolVarP(&http, "http", "", false, "http application")
	Cmd.PersistentFlags().BoolVarP(&grpc, "grpc", "", false, "grpc application")
	Cmd.PersistentFlags().BoolVarP(&stream, "stream", "", false, "stream application")
	Cmd.PersistentFlags().BoolVarP(&schedule, "schedule", "", false, "schedule application")

	_ = initCmd.MarkPersistentFlagRequired("module")
	Cmd.AddCommand(initCmd)

	_ = initCmd.MarkPersistentFlagRequired("module")
	_ = initCmd.MarkPersistentFlagRequired("app")
	Cmd.AddCommand(addCmd)
}
