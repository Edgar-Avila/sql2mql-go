/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd (Comando base)
var rootCmd = &cobra.Command{
	Use:   "sqlmql",
	Short: "Translate SQL syntax to mongo or execute mongo with SQL syntax",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
