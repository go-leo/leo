package project

import (
	"github.com/go-leo/leo/v3/cmd/leo/internal/gonew"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

func New() *cobra.Command {
	var dstMod string
	cmd := &cobra.Command{
		Use:   "project",
		Short: "create a new project",
		Run: func(cmd *cobra.Command, args []string) {
			srcMod := "github.com/go-leo/project-layout"
			srcModVers := srcMod + "@latest"
			dir := "./" + filepath.Base(dstMod)
			log.Printf("create project %s...\n", dir)
			gonew.GoNew(srcMod, srcModVers, dstMod, dir, []string{})
			log.Println("success")
		},
	}
	cmd.Flags().StringVarP(&dstMod, "mod", "m", "", "project module")
	if err := cmd.MarkFlagRequired("mod"); err != nil {
		panic(err)
	}
	return cmd
}
