package main

import (
	"log"

	"codeup.aliyun.com/qimao/leo/leo/cmd/leo/project"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "leo",
	Short: "Leo: An elegant toolkit for Go microservices.",
	Long:  `Leo: An elegant toolkit for Go microservices.`,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
