package main

import (
	"github.com/go-leo/leo/v3/cmd/leo/app"
	"github.com/go-leo/leo/v3/cmd/leo/project"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:     "leo",
	Short:   "leo is a tool for generate project layout",
	Version: "v3.0.0",
}

func init() {
	rootCmd.AddCommand(project.New())
	rootCmd.AddCommand(app.New())
}
