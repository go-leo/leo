package project

import (
	"github.com/go-leo/leo/v3/cmd/leo/internal/gonew"
	"github.com/spf13/cobra"
	"path/filepath"
)

func New() *cobra.Command {
	var dstMod string
	cmd := &cobra.Command{
		Use:   "app",
		Short: "Create a new app",
		Run: func(cmd *cobra.Command, args []string) {
			srcMod := "github.com/go-leo/app-layout"
			srcModVers := srcMod + "@latest"
			gonew.GoNew(srcMod, srcModVers, dstMod, "./"+filepath.Base(dstMod))
		},
	}
	cmd.Flags().StringVarP(&dstMod, "dst-mod", "d", "", "destination module")
	if err := cmd.MarkFlagRequired("dst-mod"); err != nil {
		panic(err)
	}
	return cmd
}
