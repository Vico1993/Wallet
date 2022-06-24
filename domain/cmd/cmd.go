package cmd

import (
	"os"
)

func Execute() {
	rootCmd := InitRootCommand(os.Stdout, os.Stderr)

	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}