package upgrade

import (
	"fmt"

	"github.com/spf13/cobra"

	"codeup.aliyun.com/qimao/leo/leo/cmd/leo/base"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the leo tools",
	Long:  "Upgrade the leo tools. Example: leo upgrade",
	Run:   Run,
}

// Run upgrade the leo tools.
func Run(_ *cobra.Command, _ []string) {
	err := base.GoInstall(
		"codeup.aliyun.com/qimao/leo/leo/cmd/leo@latest",
		"google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1",
		"google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0",
		"github.com/favadi/protoc-go-inject-tag@latest",
		"github.com/google/gnostic/cmd/protoc-gen-openapi@latest",
		"go.uber.org/mock/mockgen@latest",
		"github.com/envoyproxy/protoc-gen-validate@latest",
		"mvdan.cc/gofumpt@latest",
		"codeup.aliyun.com/qimao/go-contrib/protoc-gen-go-enum@latest",
		"github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}
