package app

import (
	"github.com/go-leo/leo/v3/cmd/leo/internal/gonew"
	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
	"log"
	"os"
	"path"
	"path/filepath"
)

func New() *cobra.Command {
	var appName string
	cmd := &cobra.Command{
		Use:   "app",
		Short: "create a new app",
		Run: func(cmd *cobra.Command, args []string) {
			wd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			goModPath := filepath.Join(wd, "go.mod")
			goModContent, err := os.ReadFile(goModPath)
			if err != nil {
				panic(err)
			}
			goMode, err := modfile.Parse(goModPath, goModContent, nil)
			if err != nil {
				panic(err)
			}
			modPath := goMode.Module.Mod.Path
			dstMod := path.Join(modPath, "app", appName)

			srcMod := "github.com/go-leo/app-layout"
			srcModVers := srcMod + "@latest"

			dir := filepath.Join("./app", filepath.Base(dstMod))
			excludes := []string{"go.mod", "go.sum"}
			log.Printf("create app %s...\n", dir)
			gonew.GoNew(srcMod, srcModVers, dstMod, dir, excludes)
			log.Println("success")
		},
	}
	cmd.Flags().StringVarP(&appName, "name", "n", "", "app name")
	if err := cmd.MarkFlagRequired("name"); err != nil {
		panic(err)
	}
	return cmd
}
