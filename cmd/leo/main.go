package main

import (
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
	Short:   "leo is ",
	Example: "",
	Version: "v3.0.0",
}
